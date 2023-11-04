package keeper_test

import (
	"testing"

	testkeeper "github.com/dhanendras/cosmos/testutil/keeper"
	"github.com/dhanendras/cosmos/x/cosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmosKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
