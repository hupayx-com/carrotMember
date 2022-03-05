package carrotmember_test

import (
	"testing"

	keepertest "github.com/hupayx-com/carrotMember/testutil/keeper"
	"github.com/hupayx-com/carrotMember/testutil/nullify"
	"github.com/hupayx-com/carrotMember/x/carrotmember"
	"github.com/hupayx-com/carrotMember/x/carrotmember/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CarrotmemberKeeper(t)
	carrotmember.InitGenesis(ctx, *k, genesisState)
	got := carrotmember.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
