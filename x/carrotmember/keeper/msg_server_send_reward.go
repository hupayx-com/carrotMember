package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
)

func (k msgServer) SendReward(goCtx context.Context, msg *types.MsgSendReward) (*types.MsgSendRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)

	// 다음 리워드 일정이 지났으면 보상을 지급하자
	// timeStore := prefix.NewStore(store, []byte(types.RewardTimeKey))

	// bz := timeStore.Get([]byte(types.RewardTimeKey))

	// types.MustUnMarshalTime()

	// MemberKey = []byte{0x31} // key for a delegation
	iterator := sdk.KVStorePrefixIterator(store, types.MemberKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var member types.Member
		k.cdc.MustUnmarshal(iterator.Value(), &member)

		fmt.Print(member)

		coins := sdk.NewCoins(member.GetAmount())

		to, err := sdk.AccAddressFromBech32(member.MemberAddress)
		if err != nil {
			panic(err)
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins); err != nil {
			return nil, err
		}

		// 일단 멤버에서 삭제하고
		store.Delete(types.GetRewardMemberKey(to))

		// 남은횟수가 1회보다 크면 재등록한다.
		if member.RemainCnt > 1 {
			// 새로 저장하자...
			member.RemainCnt = member.RemainCnt - 1
			b := k.cdc.MustMarshal(&member)
			store.Set(types.GetRewardMemberKey(to), b)
		}
	}

	// 다음 리워드 일시를 저장해주자
	// timeStore.Set()

	return &types.MsgSendRewardResponse{}, nil
}
