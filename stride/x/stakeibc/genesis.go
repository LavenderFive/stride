package stakeibc

import (
	"github.com/Stride-labs/stride/x/stakeibc/keeper"
	"github.com/Stride-labs/stride/x/stakeibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// TODO: TEST-10
	// Set if defined
	// 	if genState.HostZone != nil {
	// 		k.SetHostZone(ctx, *genState.HostZone)
	// 	}
	// 	// Set if defined
	// if genState.ICAAccount != nil {
	// 	k.SetICAAccount(ctx, *genState.ICAAccount)
	// }

	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	// Get all hostZone

	// TODO: TEST-10
	// 	hostZone, found := k.GetHostZone(ctx)
	// 	if found {
	// 		genesis.HostZone = &hostZone
	// 	}
	// 	// Get all iCAAccount
	// iCAAccount, found := k.GetICAAccount(ctx)
	// if found {
	// 	genesis.ICAAccount = &iCAAccount
	// }
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
