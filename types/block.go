// Copyright Tharsis Labs Ltd.(Blackfury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/elysiumstation/blackfury/blob/main/LICENSE)
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BlockGasLimit returns the max gas (limit) defined in the block gas meter. If the meter is not
// set, it returns the max gas from the application consensus params.
// NOTE: see https://github.com/cosmos/cosmos-sdk/issues/9514 for full reference
func BlockGasLimit(ctx sdk.Context) uint64 {
	blockGasMeter := ctx.BlockGasMeter()

	// Get the limit from the gas meter only if its not null and not an InfiniteGasMeter
	if blockGasMeter != nil && blockGasMeter.Limit() != 0 {
		return blockGasMeter.Limit()
	}

	// Otherwise get from the consensus parameters
	cp := ctx.ConsensusParams()
	if cp == nil || cp.Block == nil {
		return 0
	}

	maxGas := cp.Block.MaxGas
	if maxGas > 0 {
		return uint64(maxGas) // #nosec G701 -- maxGas is int64 type. It can never be greater than math.MaxUint64
	}

	return 0
}
