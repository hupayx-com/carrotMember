package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
)

func (k msgServer) AddMember(goCtx context.Context, msg *types.MsgAddMember) (*types.MsgAddMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	// _ = ctx

	var member = types.Member{
		Creator:       msg.Creator,
		MemberAddress: msg.MemberAddress,
		Amount:        msg.Amount,
		// CreateTime: ctx.BlockTime().time,
		TotalCnt:  msg.TotalCnt,
		RemainCnt: msg.TotalCnt,
	}

	fmt.Print(member)

	k.AddRewardMember(ctx, member)

	return &types.MsgAddMemberResponse{}, nil
}

// 멤버를 등록한다.
func (k Keeper) AddRewardMember(ctx sdk.Context, member types.Member) error {
	toAddress, err := sdk.AccAddressFromBech32(member.Creator)

	if err != nil {
		panic(err)
	}

	fmt.Print(toAddress)

	// 주소 체크하고
	memAddress, err := sdk.AccAddressFromBech32(member.MemberAddress)
	if err != nil {
		panic(err)
	}

	fmt.Print(memAddress)

	// // ModuleName
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&member)

	// TODO
	amount := member.Amount.Amount
	coins := sdk.NewCoins(sdk.NewCoin(member.Amount.Denom, amount))

	fmt.Print(amount)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, toAddress, types.ModuleName, coins); err != nil {
		return err
	}

	// 저장하자
	store.Set(types.GetRewardMemberKey(memAddress), b)

	return nil
}
