package bank_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nexusblock/heimdall/app"
	authTypes "github.com/nexusblock/heimdall/auth/types"
	bankTypes "github.com/nexusblock/heimdall/bank/types"
	abci "github.com/tendermint/tendermint/abci/types"
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
	app.BankKeeper.SetSendEnabled(ctx, bankTypes.DefaultSendEnabled)

	return app, ctx
}
