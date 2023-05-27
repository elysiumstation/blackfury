// Copyright Tharsis Labs Ltd.(Blackfury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/elysiumstation/blackfury/blob/main/LICENSE)

package epochs

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elysiumstation/blackfury/v13/x/epochs/keeper"
	"github.com/elysiumstation/blackfury/v13/x/epochs/types"
)

// InitGenesis initializes the epochs module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// set epoch info from genesis
	for _, epoch := range genState.Epochs {
		// Initialize empty epoch values via Cosmos SDK
		if epoch.StartTime.Equal(time.Time{}) || epoch.StartTime.IsZero() {
			epoch.StartTime = ctx.BlockTime()
		}

		epoch.CurrentEpochStartHeight = ctx.BlockHeight()

		k.SetEpochInfo(ctx, epoch)
	}
}

// ExportGenesis returns the epochs module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Epochs: k.AllEpochInfos(ctx),
	}
}
