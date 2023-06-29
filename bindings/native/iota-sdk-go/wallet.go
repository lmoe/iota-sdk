package iota_sdk_go

import (
	"iota_sdk_go/methods"
	"iota_sdk_go/types"
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

func BuildBip32Chain(coinType types.CoinType, accountIndex uint32, internalAddress bool, addressIndex uint32) types.IBip32Chain {
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

func (w *Wallet) SignTransactionEssence(txEssence types.HexEncodedString, bip32Chain types.IBip32Chain) (*types.Ed25519Signature, error) {
	signedMessageStr, err := w.sdk.CallSecretManagerMethod(w.secretManagerPtr, methods.SignEd25519Method(types.SignEd25519MethodData{
		Message: txEssence,
		Chain:   bip32Chain,
	}))
	if err != nil {
		return nil, err
	}

	return ParseResponse[types.Ed25519Signature](signedMessageStr, err)
}
