package keeper_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	evm "github.com/elysiumstation/blackfury/v13/x/evm/types"

	"github.com/elysiumstation/blackfury/v13/app"
	"github.com/elysiumstation/blackfury/v13/x/epochs/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx            sdk.Context
	app            *app.Blackfury
	queryClientEvm evm.QueryClient
	queryClient    types.QueryClient
	consAddress    sdk.ConsAddress
}

var s *KeeperTestSuite

func TestKeeperTestSuite(t *testing.T) {
	s = new(KeeperTestSuite)
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.DoSetupTest()
}
