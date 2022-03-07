package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NextRewardTime(goCtx context.Context, req *types.QueryNextRewardTimeRequest) (*types.QueryNextRewardTimeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	nextTimeStr := ""
	store := ctx.KVStore(k.storeKey)

	ntime := store.Get(types.NextRewardTimeKey)

	if ntime != nil {
		nextTime, _ := sdk.ParseTimeBytes(ntime)
		nextTimeStr = nextTime.Format(time.RFC3339)
	}

	// TODO: Process the query
	_ = ctx

	return &types.QueryNextRewardTimeResponse{NextTime: nextTimeStr}, nil
}
