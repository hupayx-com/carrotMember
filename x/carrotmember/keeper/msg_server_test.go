package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/hupayx-com/carrotMember/testutil/keeper"
	"github.com/hupayx-com/carrotMember/x/carrotmember/keeper"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CarrotmemberKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
