package iota_sdk_go

import (
	"encoding/json"

	"iota_sdk_go/methods"
	"iota_sdk_go/types"

	"github.com/iotaledger/hive.go/serializer/v2"
	iotago "github.com/iotaledger/iota.go/v3"
)

type Wallet struct {
	sdk              *IOTASDK
	walletPtr        IotaWalletPtr
	clientPtr        IotaClientPtr
	secretManagerPtr IotaSecretManagerPtr
}

func NewWallet(sdk *IOTASDK, walletPtr IotaWalletPtr, clientPtr IotaClientPtr, secretManagerPtr IotaSecretManagerPtr) *Wallet {
	return &Wallet{
		sdk:              sdk,
		walletPtr:        walletPtr,
		clientPtr:        clientPtr,
		secretManagerPtr: secretManagerPtr,
	}
}

func (w *Wallet) Destroy() error {
	return w.sdk.DestroyWallet(w.walletPtr)
}

func (w *Wallet) GetLedgerStatus() (*types.LedgerNanoStatus, error) {
	ledgerNanoStatus, err := w.sdk.CallWalletMethod(w.walletPtr, methods.GetLedgerNanoStatusMethod())
	if err != nil {
		return nil, err
	}

	status, err := ParseResponse[types.LedgerNanoStatus](ledgerNanoStatus, err)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (w *Wallet) GenerateEd25519Address(addressIndex float64, accountIndex float64, bech32Hrp string, options *types.GenerateAddressOptions) (string, error) {
	ledgerNanoStatus, err := w.sdk.CallWalletMethod(w.walletPtr, methods.GenerateEd25519AddressMethod(types.GenerateEd25519AddressMethodData{
		AddressIndex: addressIndex,
		AccountIndex: accountIndex,
		Bech32Hrp:    bech32Hrp,
		Options:      options,
	}))
	if err != nil {
		return "", err
	}

	address, err := ParseResponse[string](ledgerNanoStatus, err)
	if err != nil || address == nil {
		return "", err
	}

	return *address, nil
}

func TranslateIotaGoTransaction(transaction iotago.Transaction) *types.PreparedTransactionData {
	inputs := make([]types.InputSigningData, 0)

	for _, input := range transaction.Essence.Outputs {
		var signData types.InputSigningData = types.InputSigningData{}

		b, _ := input.MarshalJSON()
		json.Unmarshal(b, &signData.Output)
		inputs = append(inputs, signData)
	}

	signingBytes, err := transaction.Essence.Serialize(serializer.DeSeriModeNoValidation, nil)
	if err != nil {
		panic(err)
	}
	//signingMsg, err := transaction.Essence.SigningMessage()

	signedHex := iotago.EncodeHex(signingBytes)

	transactionData := types.PreparedTransactionData{
		Essence:    types.NewRegularTransactionEssence("", signedHex),
		InputsData: inputs,
	}

	return &transactionData
}

func buildBip32Chain(coinType types.CoinType, accountIndex uint32, internalAddress bool, addressIndex uint32) types.IBip32Chain {
	var internalAddressInt uint32 = 0

	if internalAddress {
		internalAddressInt = 1
	}

	return types.IBip32Chain{
		uint32(types.HDWalletType),
		uint32(coinType),
		accountIndex,
		internalAddressInt,
		addressIndex,
	}
}

func (w *Wallet) SignTransactionEssence(txEssence types.HexEncodedString, accountIndex uint32, addressIndex uint32) (*types.Ed25519Signature, error) {
	bip32Chain := buildBip32Chain(types.CoinTypeSMR, accountIndex, false, addressIndex)

	signedMessageStr, err := w.sdk.CallSecretManagerMethod(w.secretManagerPtr, methods.SignEd25519Method(types.SignEd25519MethodData{
		Message: txEssence,
		Chain:   bip32Chain,
	}))
	if err != nil {
		return nil, err
	}

	var signature types.Ed25519Signature
	if err = json.Unmarshal([]byte(signedMessageStr), &signature); err != nil {
		return nil, err
	}

	return &signature, nil
}

func (w *Wallet) SignTransaction(transaction types.PreparedTransactionData) (any, error) {
	_, err := w.sdk.CallSecretManagerMethod(w.secretManagerPtr, methods.SignTransactionMethod(types.SignTransactionMethodData{
		SecretManager: types.LedgerNanoSecretManager{
			LedgerNano: false,
		},
		PreparedTransactionData: transaction,
	}))
	if err != nil {
		return "", err
	}

	return nil, nil
}
