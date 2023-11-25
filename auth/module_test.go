package auth_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/nexusblock/heimdall/app"
	authTypes "github.com/nexusblock/heimdall/auth/types"
	supplyTypes "github.com/nexusblock/heimdall/supply/types"
)

//
// Test suite
//

// ModuleTestSuite integrate test suite context object
type ModuleTestSuite struct {
	suite.Suite

	app *app.HeimdallApp
	ctx sdk.Context
}

func (suite *ModuleTestSuite) SetupTest() {
	suite.app, suite.ctx = createTestApp(false)
}

func TestModuleTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ModuleTestSuite))
}

//
// Tests
//

func (suite *ModuleTestSuite) TestModuleAccount() {
	t, happ, ctx := suite.T(), suite.app, suite.ctx
	acc := happ.AccountKeeper.GetAccount(ctx, supplyTypes.NewModuleAddress(authTypes.FeeCollectorName))
	require.NotNil(t, acc)
}
