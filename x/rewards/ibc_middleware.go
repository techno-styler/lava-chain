package rewards

import (
	"errors"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/x/rewards/keeper"
	"github.com/lavanet/lava/x/rewards/types"
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 callbacks for the transfer middleware given
// the rewards keeper and the underlying application.
type IBCMiddleware struct {
	app    porttypes.IBCModule // transfer stack
	keeper keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(app porttypes.IBCModule, k keeper.Keeper) IBCMiddleware {
	return IBCMiddleware{
		app:    app,
		keeper: k,
	}
}

// IBCModule interface implementation. Only OnRecvPacket() calls a callback from the keeper. The rest have default implementations

func (im IBCMiddleware) OnChanOpenInit(ctx sdk.Context, order channeltypes.Order, connectionHops []string, portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	return im.app.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version)
}

func (im IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID, channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (version string, err error) {
	return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, counterpartyVersion)
}

func (im IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID, channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

func (im IBCMiddleware) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

func (im IBCMiddleware) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

func (im IBCMiddleware) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

// OnRecvPacket checks the packet's memo and funds the IPRPC pool accordingly. If the memo is not the expected JSON,
// the packet is transferred normally to the next IBC module in the transfer stack
func (im IBCMiddleware) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) exported.Acknowledgement {
	// unmarshal the packet's data with the transfer module codec (expect an ibc-transfer packet)
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// extract the packet's memo
	memo, err := im.keeper.ExtractIprpcMemoFromPacket(ctx, data)
	if errors.Is(err, types.ErrMemoNotIprpcOverIbc) {
		// not a packet that should be handled as IPRPC over IBC (not considered as error)
		utils.LavaFormatDebug("rewards module IBC middleware processing skipped, memo is invalid for IPRPC over IBC funding",
			utils.LogAttr("memo", memo),
		)
		return im.app.OnRecvPacket(ctx, packet, relayer)
	} else if errors.Is(err, types.ErrIprpcMemoInvalid) {
		// memo is in the right format of IPRPC over IBC but the data is invalid
		utils.LavaFormatWarning("rewards module IBC middleware processing failed, memo data is invalid", err,
			utils.LogAttr("memo", memo))
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// change the ibc-transfer packet receiver address to be a temp address and empty the memo
	data.Receiver = types.IbcIprpcReceiverAddress().String()
	data.Memo = ""
	marshelledData, err := transfertypes.ModuleCdc.MarshalJSON(&data)
	if err != nil {
		utils.LavaFormatError("rewards module IBC middleware processing failed, cannot marshal packet data", err,
			utils.LogAttr("data", data))
		return channeltypes.NewErrorAcknowledgement(err)
	}
	packet.Data = marshelledData

	// get the IBC IPRPC receiver balance before getting the new IBC tokens
	balanceBefore := im.keeper.GetIbcIprpcReceiverBalance(ctx)

	// call the next OnRecvPacket() of the transfer stack to make the IbcIprpcReceiver address get the IBC tokens
	ack := im.app.OnRecvPacket(ctx, packet, relayer)
	if ack == nil || !ack.Success() {
		return ack
	}

	// get the IBC IPRPC receiver balance after getting the new IBC tokens and get the difference
	// this is done to get the new tokens with the IBC denom
	balanceAfter := im.keeper.GetIbcIprpcReceiverBalance(ctx)
	ibcFundCoins := balanceAfter.Sub(balanceBefore...)
	if ibcFundCoins.IsZero() {
		utils.LavaFormatError("rewards module IBC middleware processing failed", fmt.Errorf("ibc iprpc receiver did not get new tokens"),
			utils.LogAttr("data", data),
		)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// ibcFundCoins should have only one token since ibc-transfer allows sending only one type of token
	if len(ibcFundCoins) != 1 {
		utils.LavaFormatError("rewards module IBC middleware processing failed", fmt.Errorf("ibcFundCoins has more than one type of coins"),
			utils.LogAttr("data", data),
		)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// set pending IPRPC over IBC requests on-chain
	piif, err := im.keeper.NewPendingIbcIprpcFund(ctx, memo.Creator, memo.Spec, memo.Duration, ibcFundCoins[0])
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// make event
	utils.LogLavaEvent(ctx, im.keeper.Logger(ctx), types.NewPendingIbcIprpcFundEventName, map[string]string{
		"index":        strconv.FormatUint(piif.Index, 10),
		"creator":      piif.Creator,
		"spec":         piif.Spec,
		"duration":     strconv.FormatUint(piif.Duration, 10),
		"monthly_fund": piif.Fund.String(),
		"expiry":       strconv.FormatUint(piif.Expiry, 10),
	}, "New pending IBC IPRPC fund was created successfully")

	return nil
}

func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}

// ICS4Wrapper interface (default implementations)
func (im IBCMiddleware) SendPacket(ctx sdk.Context, chanCap *capabilitytypes.Capability, sourcePort string, sourceChannel string,
	timeoutHeight clienttypes.Height, timeoutTimestamp uint64, data []byte,
) (sequence uint64, err error) {
	return im.keeper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

func (im IBCMiddleware) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet exported.PacketI,
	ack exported.Acknowledgement,
) error {
	return im.keeper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

func (im IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return im.keeper.GetAppVersion(ctx, portID, channelID)
}