//go:build gofuzz || go1.18

package types_test

import (
	"testing"

	utiltx "github.com/elysiumstation/blackfury/v13/testutil/tx"
	"github.com/elysiumstation/blackfury/v13/x/incentives/types"
)

func FuzzSplitGasMeterKey(f *testing.F) {
	contract := utiltx.GenerateAddress()
	user := utiltx.GenerateAddress()

	key := types.KeyPrefixGasMeter
	key = append(key, contract.Bytes()...)
	key = append(key, user.Bytes()...)
	f.Add(key)
	f.Fuzz(func(t *testing.T, key []byte) {
		types.SplitGasMeterKey(key)
	})
}
