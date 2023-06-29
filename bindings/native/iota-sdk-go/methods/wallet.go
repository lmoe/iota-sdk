package methods

import "iota_sdk_go/types"

func BackupMethod[T types.BackupMethodData](data T) BaseRequest[T] {
	var method = "backup"

	return NewBaseRequest[T](method, data)
}

func ChangeStrongholdPasswordMethod[T types.ChangeStrongholdPasswordMethodData](data T) BaseRequest[T] {
	var method = "changeStrongholdPassword"

	return NewBaseRequest[T](method, data)
}

func ClearStrongholdPasswordMethod() BaseRequest[NoType] {
	var method = "clearStrongholdPassword"

	return NewBaseRequestNoData(method)
}

func ClearListenersMethod[T types.ClearListenersMethodData](data T) BaseRequest[T] {
	var method = "clearListeners"

	return NewBaseRequest(method, data)
}

func CreateAccountMethod[T types.CreateAccountPayload](data T) BaseRequest[T] {
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

func GetAccountMethod[T types.GetAccountMethodData](data T) BaseRequest[T] {
	var method = "getAccount"

	return NewBaseRequest(method, data)
}

func IsStrongholdPasswordAvailableMethod() BaseRequest[NoType] {
	var method = "isStrongholdPasswordAvailable"

	return NewBaseRequestNoData(method)
}

func RecoverAccountsMethod[T types.RecoverAccountsMethodData](data T) BaseRequest[T] {
	var method = "recoverAccounts"

	return NewBaseRequest[T](method, data)
}

func RemoveLatestAccountMethod() BaseRequest[NoType] {
	var method = "removeLatestAccount"

	return NewBaseRequestNoData(method)
}

func RestoreBackupMethod[T types.RestoreBackupMethodData](data T) BaseRequest[T] {
	var method = "restoreBackup"

	return NewBaseRequest[T](method, data)
}

func SetClientOptionsMethod[T types.SetClientOptionsMethodData](data T) BaseRequest[T] {
	var method = "setClientOptions"

	return NewBaseRequest[T](method, data)
}

func SetStrongholdPasswordMethod[T types.SetStrongholdPasswordMethodData](data T) BaseRequest[T] {
	var method = "setStrongholdPassword"

	return NewBaseRequest[T](method, data)
}

func SetStrongholdPasswordClearIntervalMethod[T types.SetStrongholdPasswordClearIntervalMethodData](data T) BaseRequest[T] {
	var method = "setStrongholdPasswordClearInterval"

	return NewBaseRequest[T](method, data)
}

func StartBackgroundSyncMethod[T types.StartBackgroundSyncMethodData](data T) BaseRequest[T] {
	var method = "startBackgroundSync"

	return NewBaseRequest[T](method, data)
}

func StopBackgroundSyncMethod() BaseRequest[NoType] {
	var method = "stopBackgroundSync"

	return NewBaseRequestNoData(method)
}

func UpdateNodeAuthMethod[T types.UpdateNodeAuthMethodData](data T) BaseRequest[T] {
	var method = "updateNodeAuth"

	return NewBaseRequest[T](method, data)
}
