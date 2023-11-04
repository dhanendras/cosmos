package keeper

import (
	"github.com/dhanendras/cosmos/x/cosmos/types"
)

var _ types.QueryServer = Keeper{}
