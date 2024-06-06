package chainlib

import (
	"context"
	"net/http"
	"sync"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/lavanet/lava/protocol/chainlib/extensionslib"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionManager_HappyFlow(t *testing.T) {
	playbook := []struct {
		name                    string
		specId                  string
		apiInterface            string
		connectionType          string
		subscriptionRequestData []byte
		subscriptionFirstReply  []byte
	}{
		{
			name:                    "TendermintRPC",
			specId:                  "LAV1",
			apiInterface:            spectypes.APIInterfaceTendermintRPC,
			connectionType:          "",
			subscriptionRequestData: []byte(`{"jsonrpc":"2.0","id":3,"method":"subscribe","params":{"query":"tm.event='NewBlock'"}}`),
			subscriptionFirstReply:  []byte(`{"jsonrpc":"2.0","id":3,"result":{}}`),
		},
		{
			name:                    "JsonRPC",
			specId:                  "ETH1",
			apiInterface:            spectypes.APIInterfaceJsonRPC,
			connectionType:          "POST",
			subscriptionRequestData: []byte(`{"jsonrpc":"2.0","id":5,"method":"eth_subscribe","params":["newHeads"]}`),
			subscriptionFirstReply:  []byte(`{"jsonrpc":"2.0","id":5,"result":"0x1234567890"}`),
		},
	}

	for _, play := range playbook {
		t.Run(play.name, func(t *testing.T) {
			ts := SetupForTests(t, 1, play.specId, "../../")

			wg := sync.WaitGroup{}
			wg.Add(1)
			// msgCount := 0
			upgrader := websocket.Upgrader{}

			// Create a simple websocket server that mocks the node
			handleWebSocket := func(w http.ResponseWriter, r *http.Request) {
				conn, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					require.NoError(t, err)
					return
				}
				defer conn.Close()

				for {
					// Read the request
					messageType, message, err := conn.ReadMessage()
					if err != nil {
						require.NoError(t, err)
						return
					}

					wg.Done()

					require.Equal(t, string(play.subscriptionRequestData)+"\n", string(message))

					// Write the first reply
					err = conn.WriteMessage(messageType, play.subscriptionFirstReply)
					if err != nil {
						require.NoError(t, err)
						return
					}
				}
			}

			chainParser, chainRouter, _, closeServer, _, err := CreateChainLibMocks(context.Background(), play.specId, play.apiInterface, nil, handleWebSocket, "../../", nil)
			require.NoError(t, err)
			if closeServer != nil {
				defer closeServer()
			}

			// Create the relay request and chain message
			relayRequest := &pairingtypes.RelayRequest{
				RelayData: &pairingtypes.RelayPrivateData{
					Data: play.subscriptionRequestData,
				},
				RelaySession: &pairingtypes.RelaySession{},
			}

			chainMessage, err := chainParser.ParseMsg("", []byte(play.subscriptionRequestData), play.connectionType, nil, extensionslib.ExtensionInfo{LatestBlock: 0})
			require.NoError(t, err)

			// Create the provider node subscription manager
			pnsm := NewProviderNodeSubscriptionManager(chainRouter, chainParser, ts.Providers[0].SK)

			consumerChannel := make(chan *pairingtypes.RelayReply)

			// Read the consumer channel that simulates consumer
			go func() {
				reply := <-consumerChannel
				require.NotNil(t, reply)
				require.Equal(t, string(play.subscriptionFirstReply), string(reply.Data))
			}()

			// Subscribe to the chain
			subscriptionId, err := pnsm.AddConsumer(ts.Ctx, relayRequest, chainMessage, ts.Consumer.Addr, consumerChannel)
			require.NoError(t, err)
			require.NotEmpty(t, subscriptionId)

			wg.Wait() // Make sure the subscription manager sent a message to the node

			// Subscribe to the same subscription again, should return the same subscription id
			subscriptionIdNew, err := pnsm.AddConsumer(ts.Ctx, relayRequest, chainMessage, ts.Consumer.Addr, consumerChannel)
			require.NoError(t, err)
			require.NotEmpty(t, subscriptionId)
			require.Equal(t, subscriptionId, subscriptionIdNew)

			// Cut the subscription, and re-subscribe, should send another message to node
			err = pnsm.RemoveConsumer(ts.Ctx, chainMessage, ts.Consumer.Addr, true)
			require.NoError(t, err)

			// Make sure the consumer channel is closed
			_, ok := <-consumerChannel
			require.False(t, ok)

			consumerChannel = make(chan *pairingtypes.RelayReply)

			// Read the consumer channel that simulates consumer
			go func() {
				reply := <-consumerChannel
				require.NotNil(t, reply)
				require.Equal(t, string(play.subscriptionFirstReply), string(reply.Data))
			}()

			wg.Add(1) // Should send another message to the node

			subscriptionId, err = pnsm.AddConsumer(ts.Ctx, relayRequest, chainMessage, ts.Consumer.Addr, consumerChannel)
			require.NoError(t, err)
			require.NotEmpty(t, subscriptionId)

			wg.Wait() // Make sure the subscription manager sent another message to the node
		})
	}
}