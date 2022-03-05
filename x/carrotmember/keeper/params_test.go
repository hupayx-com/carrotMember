package keeper_test

import (
	"testing"

	testkeeper "github.com/hupayx-com/carrotMember/testutil/keeper"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CarrotmemberKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
