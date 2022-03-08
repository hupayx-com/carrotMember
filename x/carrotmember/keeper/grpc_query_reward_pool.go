package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RewardPool(goCtx context.Context, req *types.QueryRewardPoolRequest) (*types.QueryRewardPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	rewardPool := k.authKeeper.GetModuleAccount(ctx, types.ModuleName)
	coins := k.bankKeeper.GetAllBalances(ctx, rewardPool.GetAddress())

	return &types.QueryRewardPoolResponse{Amount: coins}, nil
}
