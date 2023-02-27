package incentives_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	kariapptesting "hub/app/testing"

	"hub/x/incentives/types"
)

// TestCreatesModuleAccountAtGenesis asserts that the incentives module account
// is properly registered with the auth module at genesis.
func TestCreatesModuleAccountAtGenesis(t *testing.T) {
	app := kariapptesting.MakeSimpleMockApp()
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	acc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)
}
