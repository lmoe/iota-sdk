package methods

import "iota_sdk_go/types"

func SignEd25519Method[T types.SignEd25519MethodData](data T) BaseRequest[T] {
	var method = "signEd25519"

	return NewBaseRequest[T](method, data)
}

func SignSecp256K1EcdsaMethod[T types.SignSecp256K1EcdsaMethodData](data T) BaseRequest[T] {
	var method = "signSecp256k1Ecdsa"

	return NewBaseRequest[T](method, data)
}

func StoreMnemonicMethod[T types.StoreMnemonicMethodData](data T) BaseRequest[T] {
	var method = "storeMnemonic"

	return NewBaseRequest[T](method, data)
}

func SignatureUnlockMethod[T types.SignatureUnlockMethodData](data T) BaseRequest[T] {
	var method = "signatureUnlock"

	return NewBaseRequest[T](method, data)
}

func SignTransactionMethod[T types.SignTransactionMethodData](data T) BaseRequest[T] {
	var method = "signTransaction"

	return NewBaseRequest[T](method, data)
}

func GetLedgerNanoStatusMethod() BaseRequest[NoType] {
	var method = "getLedgerNanoStatus"

	return NewBaseRequestNoData(method)
}

func GenerateEd25519AddressMethod[T types.GenerateEd25519AddressMethodData](data T) BaseRequest[T] {
	var method = "generateEd25519Address"

	return NewBaseRequest[T](method, data)
}
