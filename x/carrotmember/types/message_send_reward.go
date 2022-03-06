package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendReward = "send_reward"

var _ sdk.Msg = &MsgSendReward{}

func NewMsgSendReward(creator string) *MsgSendReward {
	return &MsgSendReward{
		Creator: creator,
	}
}

func (msg *MsgSendReward) Route() string {
	return RouterKey
}

func (msg *MsgSendReward) Type() string {
	return TypeMsgSendReward
}

func (msg *MsgSendReward) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
