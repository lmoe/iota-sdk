package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"iota_sdk_go"
	"iota_sdk_go/methods"
	"iota_sdk_go/types"

	iotago "github.com/iotaledger/iota.go/v3"
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

func GetLedgerStatus(t *testing.T, sdk *iota_sdk_go.IOTASDK, walletPtr iota_sdk_go.IotaWalletPtr) *types.LedgerNanoStatus {
	ledgerNanoStatus, err := sdk.CallWalletMethod(walletPtr, methods.GetLedgerNanoStatusMethod())
	require.NoError(t, err)

	status, err := iota_sdk_go.ParseResponse[types.LedgerNanoStatus](ledgerNanoStatus, err)
	require.NoError(t, err)
	require.NotNil(t, status)

	return status
}

func CreateAddressLedger(t *testing.T, sdk *iota_sdk_go.IOTASDK, walletPtr iota_sdk_go.IotaWalletPtr) *string {
	ledgerNanoStatus, err := sdk.CallWalletMethod(walletPtr, methods.GenerateEd25519AddressMethod(types.GenerateEd25519AddressMethodData{
		AddressIndex: 0,
		AccountIndex: 0,
		Bech32Hrp:    "SMR",
		Options: &types.GenerateAddressOptions{
			LedgerNanoPrompt: true,
		},
	}))
	require.NoError(t, err)

	status, err := iota_sdk_go.ParseResponse[string](ledgerNanoStatus, err)
	require.NoError(t, err)
	require.NotNil(t, status)

	return status
}

func SignTransactionLedger(t *testing.T, sdk *iota_sdk_go.IOTASDK, walletPtr iota_sdk_go.IotaWalletPtr) *string {
	ledgerNanoStatus, err := sdk.CallWalletMethod(walletPtr, methods.GenerateEd25519AddressMethod(types.GenerateEd25519AddressMethodData{
		AddressIndex: 0,
		AccountIndex: 0,
		Bech32Hrp:    "SMR",
		Options: &types.GenerateAddressOptions{
			LedgerNanoPrompt: true,
		},
	}))
	require.NoError(t, err)

	status, err := iota_sdk_go.ParseResponse[string](ledgerNanoStatus, err)
	require.NoError(t, err)
	require.NotNil(t, status)

	return status
}

func TestIotaGoTranslation(t *testing.T) {
	txMsg, err := os.ReadFile("./iotago_test_transaction.json")
	require.NoError(t, err)

	var tx iotago.Transaction
	json.Unmarshal(txMsg, &tx)

	ret := iota_sdk_go.TranslateIotaGoTransaction(tx)
	require.NotNil(t, ret)
}

func TestWalletLedger(t *testing.T) {
	sdk := InitTest(t)

	wallet, err := sdk.CreateWallet(types.WalletOptions{
		ClientOptions: &types.IClientOptions{
			PrimaryNode: ShimmerNetworkAPI,
		},
		SecretManager: types.LedgerNanoSecretManager{
			LedgerNano: true,
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

	address, err := wallet.GenerateEd25519Address(0, 0, "tst", nil)
	require.NoError(t, err)
	require.NotNil(t, address)

	txMsg, err := os.ReadFile("./iotago_test_transaction.json")
	require.NoError(t, err)

	var tx iotago.Transaction
	json.Unmarshal(txMsg, &tx)

	ret := iota_sdk_go.TranslateIotaGoTransaction(tx)

	signature, err := wallet.SignTransaction(*ret)
	require.NoError(t, err)
	require.NotNil(t, signature)

}
