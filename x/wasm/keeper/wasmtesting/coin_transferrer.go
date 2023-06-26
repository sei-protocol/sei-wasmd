package wasmtesting

import sdk "github.com/cosmos/cosmos-sdk/types"

type MockCoinTransferrer struct {
	TransferCoinsFn         func(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	TransferCoinsToModuleFn func(ctx sdk.Context, fromAddr sdk.AccAddress, m string, amt sdk.Coins) error
}

func (m *MockCoinTransferrer) TransferCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	if m.TransferCoinsFn == nil {
		panic("not expected to be called")
	}
	return m.TransferCoinsFn(ctx, fromAddr, toAddr, amt)
}

func (m *MockCoinTransferrer) TransferCoinsToModule(ctx sdk.Context, fromAddr sdk.AccAddress, module string, amt sdk.Coins) error {
	if m.TransferCoinsToModuleFn == nil {
		panic("not expected to be called")
	}
	return m.TransferCoinsToModule(ctx, fromAddr, module, amt)
}
