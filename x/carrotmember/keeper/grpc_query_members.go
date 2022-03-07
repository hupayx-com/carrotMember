package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Members(goCtx context.Context, req *types.QueryMembersRequest) (*types.QueryMembersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var members []*types.Member
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)

	memberStore := prefix.NewStore(store, types.MemberKey)

	pageRes, err := query.Paginate(memberStore, req.Pagination, func(key []byte, value []byte) error {
		var member types.Member
		if err := k.cdc.Unmarshal(value, &member); err != nil {
			return err
		}
		members = append(members, &member)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryMembersResponse{Member: members, Pagination: pageRes}, nil
}
