package auth_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/nexusblock/heimdall/app"
	authTypes "github.com/nexusblock/heimdall/auth/types"
)

//
// Create test app
//

// returns context and app with params set on account keeper
// nolint: unparam
func createTestApp(isCheckTx bool) (*app.HeimdallApp, sdk.Context) {
	app := app.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})
	app.AccountKeeper.SetParams(ctx, authTypes.DefaultParams())

	return app, ctx
}
