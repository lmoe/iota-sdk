package types

func BackupMethod[T BackupMethodData](data T) BaseRequest[T] {
	var method = "backup"

	return NewBaseRequest[T](method, data)
}

func ChangeStrongholdPasswordMethod[T ChangeStrongholdPasswordMethodData](data T) BaseRequest[T] {
	var method = "changeStrongholdPassword"

	return NewBaseRequest[T](method, data)
}

func ClearStrongholdPasswordMethod() BaseRequest[NoType] {
	var method = "clearStrongholdPassword"

	return NewBaseRequestNoData(method)
}

func ClearListenersMethod[T ClearListenersMethodData](data T) BaseRequest[T] {
	var method = "clearListeners"

	return NewBaseRequest(method, data)
}

func CreateAccountMethod[T CreateAccountPayload](data T) BaseRequest[T] {
	var method = "createAccount"

	return NewBaseRequest(method, data)
}

func GenerateMnemonicMethod() BaseRequest[NoType] {
	var method = "generateMnemonic"

	return NewBaseRequestNoData(method)
}

func GetAccountIndexesMethod() BaseRequest[NoType] {
	var method = "getAccountIndexes"

	return NewBaseRequestNoData(method)
}

func GetAccountsMethod() BaseRequest[NoType] {
	var method = "getAccounts"

	return NewBaseRequestNoData(method)
}

func GetAccountMethod[T GetAccountMethodData](data T) BaseRequest[T] {
	var method = "getAccount"

	return NewBaseRequest(method, data)
}

func GetLedgerNanoStatusMethod() BaseRequest[NoType] {
	var method = "getLedgerNanoStatus"

	return NewBaseRequestNoData(method)
}

func GenerateEd25519AddressMethod[T GenerateEd25519AddressMethodData](data T) BaseRequest[T] {
	var method = "generateEd25519Address"

	return NewBaseRequest[T](method, data)
}

func IsStrongholdPasswordAvailableMethod() BaseRequest[NoType] {
	var method = "isStrongholdPasswordAvailable"

	return NewBaseRequestNoData(method)
}

func RecoverAccountsMethod[T RecoverAccountsMethodData](data T) BaseRequest[T] {
	var method = "recoverAccounts"

	return NewBaseRequest[T](method, data)
}

func RemoveLatestAccountMethod() BaseRequest[NoType] {
	var method = "removeLatestAccount"

	return NewBaseRequestNoData(method)
}

func RestoreBackupMethod[T RestoreBackupMethodData](data T) BaseRequest[T] {
	var method = "restoreBackup"

	return NewBaseRequest[T](method, data)
}

func SetClientOptionsMethod[T SetClientOptionsMethodData](data T) BaseRequest[T] {
	var method = "setClientOptions"

	return NewBaseRequest[T](method, data)
}

func SetStrongholdPasswordMethod[T SetStrongholdPasswordMethodData](data T) BaseRequest[T] {
	var method = "setStrongholdPassword"

	return NewBaseRequest[T](method, data)
}

func SetStrongholdPasswordClearIntervalMethod[T SetStrongholdPasswordClearIntervalMethodData](data T) BaseRequest[T] {
	var method = "setStrongholdPasswordClearInterval"

	return NewBaseRequest[T](method, data)
}

func StartBackgroundSyncMethod[T StartBackgroundSyncMethodData](data T) BaseRequest[T] {
	var method = "startBackgroundSync"

	return NewBaseRequest[T](method, data)
}

func StopBackgroundSyncMethod() BaseRequest[NoType] {
	var method = "stopBackgroundSync"

	return NewBaseRequestNoData(method)
}

func StoreMnemonicMethod[T StoreMnemonicMethodData](data T) BaseRequest[T] {
	var method = "storeMnemonic"

	return NewBaseRequest[T](method, data)
}

func UpdateNodeAuthMethod[T UpdateNodeAuthMethodData](data T) BaseRequest[T] {
	var method = "updateNodeAuth"

	return NewBaseRequest[T](method, data)
}

func SignTransactionMethod[T SignTransactionMethodData](data T) BaseRequest[T] {
	var method = "signTransaction"

	return NewBaseRequest[T](method, data)
}

func SignatureUnlockMethod[T SignatureUnlockMethodData](data T) BaseRequest[T] {
	var method = "signatureUnlock"

	return NewBaseRequest[T](method, data)
}

func SignEd25519Method[T SignEd25519MethodData](data T) BaseRequest[T] {
	var method = "signEd25519"

	return NewBaseRequest[T](method, data)
}

func SignSecp256K1EcdsaMethod[T SignSecp256K1EcdsaMethodData](data T) BaseRequest[T] {
	var method = "signSecp256k1Ecdsa"

	return NewBaseRequest[T](method, data)
}
