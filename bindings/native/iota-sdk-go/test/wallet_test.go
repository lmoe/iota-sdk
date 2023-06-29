package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"iota_sdk_go"
	"iota_sdk_go/types"
)

func TestWalletMnemonic(t *testing.T) {
	sdk := InitTest(t)

	clientPtr, err := sdk.CreateWallet(types.WalletOptions{
		ClientOptions: &types.IClientOptions{
			PrimaryNode: ShimmerNetworkAPI,
		},
		SecretManager: types.MnemonicSecretManager{
			Mnemonic: TestMnemonic,
		},
		StoragePath: "./testdb/mnemonic",
		CoinType:    types.CoinTypeSMR,
	})

	require.NoError(t, err)
	require.NotNil(t, clientPtr)

	t.Log(clientPtr)
}

func TestWalletLedger(t *testing.T) {
	sdk := InitTest(t)

	wallet, err := sdk.CreateWallet(types.WalletOptions{
		ClientOptions: &types.IClientOptions{
			PrimaryNode: ShimmerNetworkAPI,
		},
		SecretManager: types.LedgerNanoSecretManager{
			LedgerNano: UseLedgerSimulator,
		},
		StoragePath: "./testdb/ledger",
		CoinType:    types.CoinTypeSMR,
	})
	defer wallet.Destroy()

	require.NoError(t, err)
	require.NotNil(t, wallet)

	status, err := wallet.GetLedgerStatus()
	require.NoError(t, err)
	require.NotNil(t, status)

	address, err := wallet.GenerateEd25519Address(0, 0, "smr", nil)
	require.NoError(t, err)
	require.NotNil(t, address)

	bip39Chain := iota_sdk_go.BuildBip32Chain(types.CoinTypeSMR, 0, false, 0)
	signedEssence, err := wallet.SignTransactionEssence(types.HexEncodedString("asds5d4f56sd4f56sd4f56sd4f6sd56f456sd4f56sdf456asd"), bip39Chain)

	fmt.Println(err.Error())
	require.NoError(t, err)
	require.NotNil(t, signedEssence)
}
