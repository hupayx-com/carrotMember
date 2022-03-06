package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddMember = "add_member"

var _ sdk.Msg = &MsgAddMember{}

func NewMsgAddMember(creator string, memberAddress string, totalCnt uint64, amount sdk.Coin) *MsgAddMember {
	return &MsgAddMember{
		Creator:       creator,
		MemberAddress: memberAddress,
		TotalCnt:      totalCnt,
		Amount:        amount,
	}
}

func (msg *MsgAddMember) Route() string {
	return RouterKey
}

func (msg *MsgAddMember) Type() string {
	return TypeMsgAddMember
}

func (msg *MsgAddMember) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddMember) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddMember) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
