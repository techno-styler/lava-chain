package common

import (
	"context"
	"testing"

	btcSecp256k1 "github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/lavanet/lava/testutil/keeper"
	"github.com/lavanet/lava/utils/sigs"
	conflicttypes "github.com/lavanet/lava/x/conflict/types"
	conflictconstruct "github.com/lavanet/lava/x/conflict/types/construct"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	"github.com/lavanet/lava/x/pairing/types"
	plantypes "github.com/lavanet/lava/x/plans/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	subscriptiontypes "github.com/lavanet/lava/x/subscription/types"
	"github.com/stretchr/testify/require"
)

type Account struct {
	SK   *btcSecp256k1.PrivateKey
	Addr sdk.AccAddress
}

func CreateMockSpec() spectypes.Spec {
	specName := "mockSpec"
	spec := spectypes.Spec{}
	spec.Name = specName
	spec.Index = specName
	spec.Enabled = true
	spec.ReliabilityThreshold = 4294967295
	spec.BlockDistanceForFinalizedData = 0
	spec.DataReliabilityEnabled = true
	spec.MinStakeClient = sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(100))
	spec.MinStakeProvider = sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(1000))
	spec.ApiCollections = []*spectypes.ApiCollection{{Enabled: true, CollectionData: spectypes.CollectionData{ApiInterface: "stub", Type: "GET"}, Apis: []*spectypes.Api{{Name: specName + "API", ComputeUnits: 100, Enabled: true}}}}
	spec.BlockDistanceForFinalizedData = 0
	return spec
}

func CreateMockPlan() plantypes.Plan {
	policy := plantypes.Policy{
		TotalCuLimit:       100000,
		EpochCuLimit:       10000,
		MaxProvidersToPair: 3,
		GeolocationProfile: 1,
	}
	plan := plantypes.Plan{
		Index:                    "mockPlan",
		Description:              "plan for testing",
		Type:                     "rpc",
		Block:                    100,
		Price:                    sdk.NewCoin("ulava", sdk.NewInt(100)),
		AllowOveruse:             true,
		OveruseRate:              10,
		AnnualDiscountPercentage: 20,
		PlanPolicy:               policy,
	}

	return plan
}

func CreateNewAccount(ctx context.Context, keepers testkeeper.Keepers, balance int64) (acc Account) {
	acc.SK, acc.Addr = sigs.GenerateFloatingKey()
	keepers.BankKeeper.SetBalance(sdk.UnwrapSDKContext(ctx), acc.Addr, sdk.NewCoins(sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(balance))))
	return
}

func StakeAccount(t *testing.T, ctx context.Context, keepers testkeeper.Keepers, servers testkeeper.Servers, acc Account, spec spectypes.Spec, stake int64) {
	endpoints := []epochstoragetypes.Endpoint{}
	for _, collection := range spec.ApiCollections {
		endpoints = append(endpoints, epochstoragetypes.Endpoint{IPPORT: "123", ApiInterfaces: []string{collection.CollectionData.ApiInterface}, Geolocation: 1})
	}
	_, err := servers.PairingServer.StakeProvider(ctx, &types.MsgStakeProvider{Creator: acc.Addr.String(), ChainID: spec.Index, Amount: sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.NewInt(stake)), Geolocation: 1, Endpoints: endpoints})
	require.Nil(t, err)
}

func BuySubscription(t *testing.T, ctx context.Context, keepers testkeeper.Keepers, servers testkeeper.Servers, acc Account, plan string) {
	servers.SubscriptionServer.Buy(ctx, &subscriptiontypes.MsgBuy{Creator: acc.Addr.String(), Consumer: acc.Addr.String(), Index: plan, Duration: 1})
}

func BuildRelayRequest(ctx context.Context, provider string, contentHash []byte, cuSum uint64, spec string, qos *types.QualityOfServiceReport) *types.RelaySession {
	return BuildRelayRequestWithBadge(ctx, provider, contentHash, uint64(1), cuSum, spec, qos, nil)
}

func BuildRelayRequestWithSession(ctx context.Context, provider string, contentHash []byte, sessionId uint64, cuSum uint64, spec string, qos *types.QualityOfServiceReport) *types.RelaySession {
	return BuildRelayRequestWithBadge(ctx, provider, contentHash, sessionId, cuSum, spec, qos, nil)
}

func BuildRelayRequestWithBadge(ctx context.Context, provider string, contentHash []byte, sessionId uint64, cuSum uint64, spec string, qos *types.QualityOfServiceReport, badge *types.Badge) *types.RelaySession {
	relaySession := &types.RelaySession{
		Provider:    provider,
		ContentHash: contentHash,
		SessionId:   sessionId,
		SpecId:      spec,
		CuSum:       cuSum,
		Epoch:       sdk.UnwrapSDKContext(ctx).BlockHeight(),
		RelayNum:    0,
		QosReport:   qos,
		LavaChainId: sdk.UnwrapSDKContext(ctx).BlockHeader().ChainID,
		Badge:       badge,
	}
	if qos != nil {
		qos.ComputeQoS()
	}
	return relaySession
}

func CreateMsgDetectionTest(ctx context.Context, consumer Account, provider0 Account, provider1 Account, spec spectypes.Spec) (detectionMsg *conflicttypes.MsgDetection, reply1 *types.RelayReply, reply2 *types.RelayReply, errRet error) {
	msg := &conflicttypes.MsgDetection{}
	msg.Creator = consumer.Addr.String()
	// request 0
	msg.ResponseConflict = &conflicttypes.ResponseConflict{ConflictRelayData0: &conflicttypes.ConflictRelayData{Request: &types.RelayRequest{}, Reply: &conflicttypes.ReplyMetadata{}}, ConflictRelayData1: &conflicttypes.ConflictRelayData{Request: &types.RelayRequest{}, Reply: &conflicttypes.ReplyMetadata{}}}
	msg.ResponseConflict.ConflictRelayData0.Request.RelayData = &types.RelayPrivateData{
		ConnectionType: "",
		ApiUrl:         "",
		Data:           []byte("DUMMYREQUEST"),
		RequestBlock:   100,
		ApiInterface:   "",
		Salt:           []byte{1},
	}
	msg.ResponseConflict.ConflictRelayData0.Request.RelaySession = &types.RelaySession{
		Provider:    provider0.Addr.String(),
		ContentHash: sigs.CalculateContentHashForRelayData(msg.ResponseConflict.ConflictRelayData0.Request.RelayData),
		SessionId:   uint64(1),
		SpecId:      spec.Index,
		CuSum:       0,
		Epoch:       sdk.UnwrapSDKContext(ctx).BlockHeight(),
		RelayNum:    0,
		QosReport:   &types.QualityOfServiceReport{Latency: sdk.OneDec(), Availability: sdk.OneDec(), Sync: sdk.OneDec()},
	}

	sig, err := sigs.SignRelay(consumer.SK, *msg.ResponseConflict.ConflictRelayData0.Request.RelaySession)
	if err != nil {
		return msg, nil, nil, err
	}

	msg.ResponseConflict.ConflictRelayData0.Request.RelaySession.Sig = sig

	// request 1
	temp, _ := msg.ResponseConflict.ConflictRelayData0.Request.Marshal()
	msg.ResponseConflict.ConflictRelayData1.Request.Unmarshal(temp)
	msg.ResponseConflict.ConflictRelayData1.Request.RelaySession.Provider = provider1.Addr.String()
	msg.ResponseConflict.ConflictRelayData1.Request.RelaySession.Sig = []byte{}
	sig, err = sigs.SignRelay(consumer.SK, *msg.ResponseConflict.ConflictRelayData1.Request.RelaySession)
	if err != nil {
		return msg, nil, nil, err
	}
	msg.ResponseConflict.ConflictRelayData1.Request.RelaySession.Sig = sig

	// reply 0
	reply := &types.RelayReply{
		Data:                  []byte("DUMMYREPLY"),
		Sig:                   sig,
		LatestBlock:           msg.ResponseConflict.ConflictRelayData0.Request.RelayData.RequestBlock + int64(spec.BlockDistanceForFinalizedData),
		FinalizedBlocksHashes: []byte{},
		SigBlocks:             sig,
		Metadata:              []types.Metadata{},
	}
	sig, err = sigs.SignRelayResponse(provider0.SK, reply, msg.ResponseConflict.ConflictRelayData0.Request)
	if err != nil {
		return msg, nil, nil, err
	}
	reply.Sig = sig
	sigBlocks, err := sigs.SignResponseFinalizationData(provider0.SK, reply, msg.ResponseConflict.ConflictRelayData0.Request, consumer.Addr)
	if err != nil {
		return msg, nil, nil, err
	}
	reply.SigBlocks = sigBlocks
	msg.ResponseConflict.ConflictRelayData0 = conflictconstruct.ConstructConflictRelayData(reply, msg.ResponseConflict.ConflictRelayData0.Request)
	// reply 1
	temp, _ = reply.Marshal()
	reply2 = &types.RelayReply{}
	reply2.Unmarshal(temp)
	reply2.Data = append(reply2.Data, []byte("DIFF")...)
	sig, err = sigs.SignRelayResponse(provider1.SK, reply2, msg.ResponseConflict.ConflictRelayData1.Request)
	if err != nil {
		return msg, nil, nil, err
	}
	reply2.Sig = sig
	sigBlocks, err = sigs.SignResponseFinalizationData(provider1.SK, reply2, msg.ResponseConflict.ConflictRelayData1.Request, consumer.Addr)
	if err != nil {
		return msg, nil, nil, err
	}
	reply2.SigBlocks = sigBlocks
	msg.ResponseConflict.ConflictRelayData1 = conflictconstruct.ConstructConflictRelayData(reply2, msg.ResponseConflict.ConflictRelayData1.Request)
	return msg, reply, reply2, err
}
