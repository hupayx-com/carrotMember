package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		CreateTime:    ctx.BlockTime(),
		TotalCnt:      msg.TotalCnt,
		RemainCnt:     msg.TotalCnt,
	}

	fmt.Print(member)

	k.AddRewardMember(ctx, member)

	// 보상 실행날짜를 지정한다.

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

	test := store.Get(types.GetRewardMemberKey(memAddress))
	if test != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "duplicate member address ")
	}

	amount := member.Amount.Amount.Mul(sdk.NewIntFromUint64(member.TotalCnt))
	coins := sdk.NewCoins(sdk.NewCoin(member.Amount.Denom, amount))

	fmt.Print(amount)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, toAddress, types.ModuleName, coins); err != nil {
		return err
	}

	currentTime := time.Now()

	// nextTime := time.Date(currentTime.Year(), currentTime.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	nextTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour()+1, 0, 0, 0, time.UTC)

	fmt.Println(nextTime)

	ntime := store.Get(types.NextRewardTimeKey)

	// 다음 시간이 없으면 가장 가까운 시간으로 지정한다.
	if ntime == nil {
		store.Set(types.NextRewardTimeKey, sdk.FormatTimeBytes(nextTime))
	}

	// 저장하자
	store.Set(types.GetRewardMemberKey(memAddress), b)

	return nil
}
