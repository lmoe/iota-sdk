package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"iota_sdk_go"
)

func TestWallet(t *testing.T) {
	sdk := InitTest(t)

	clientPtr, err := sdk.CreateWallet(iota_sdk_go.WalletOptions{
		ClientOptions: &iota_sdk_go.IClientOptions{
			PrimaryNode: ShimmerNetworkAPI,
		},
		SecretManager: iota_sdk_go.MnemonicSecretManager{
			Mnemonic: TestMnemonic,
		},
		StoragePath: "./testdb",
		CoinType:    iota_sdk_go.CoinTypeSMR,
	})
	require.NoError(t, err)
	require.NotNil(t, clientPtr)

	t.Log(clientPtr)
}
