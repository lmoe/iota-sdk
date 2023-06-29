package iota_sdk_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/ebitengine/purego"
	"iota_sdk_go/types"
)

func getIOTASDKLibrary() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	switch runtime.GOOS {
	case "linux":
		return wd + "/../../../../target/debug/libiota_sdk_go.so"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func serializeOrPanic[T any](obj T) string {
	jsonMessage, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return string(jsonMessage)
}

func serialize[T any](obj T) (string, error) {
	jsonMessage, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(jsonMessage), nil
}

type IotaClientPtr uintptr
type IotaWalletPtr uintptr
type IotaSecretManagerPtr uintptr

type IOTASDK struct {
	handle uintptr

	lib_initLogger func(string) bool

	lib_createClient        func(string) IotaClientPtr
	lib_createWallet        func(string) IotaWalletPtr
	lib_createSecretManager func(string) IotaSecretManagerPtr

	lib_getClientFromWallet        func(ptr IotaWalletPtr) IotaClientPtr
	lib_getSecretManagerFromWallet func(ptr IotaWalletPtr) IotaSecretManagerPtr

	lib_destroyClient        func(ptr IotaClientPtr) bool
	lib_destroyWallet        func(ptr IotaWalletPtr) bool
	lib_destroySecretManager func(ptr IotaSecretManagerPtr) bool
	lib_destroyString        func(ptr *string) bool

	lib_callClientMethod        func(IotaClientPtr, string) string
	lib_callWalletMethod        func(IotaWalletPtr, string) string
	lib_callSecretManagerMethod func(IotaSecretManagerPtr, string) string

	lib_getLastError func() string
}

func NewIotaSDK() *IOTASDK {
	libPath := getIOTASDKLibrary()

	iotaSDK, err := purego.Dlopen(libPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	iotaSDKStruct := IOTASDK{}

	purego.RegisterLibFunc(&iotaSDKStruct.lib_initLogger, iotaSDK, "init_logger")

	purego.RegisterLibFunc(&iotaSDKStruct.lib_createClient, iotaSDK, "create_client")
	purego.RegisterLibFunc(&iotaSDKStruct.lib_createWallet, iotaSDK, "create_wallet")
	purego.RegisterLibFunc(&iotaSDKStruct.lib_createSecretManager, iotaSDK, "create_secret_manager")

	purego.RegisterLibFunc(&iotaSDKStruct.lib_destroyClient, iotaSDK, "destroy_client")
	purego.RegisterLibFunc(&iotaSDKStruct.lib_destroyWallet, iotaSDK, "destroy_wallet")
	purego.RegisterLibFunc(&iotaSDKStruct.lib_destroySecretManager, iotaSDK, "destroy_secret_manager")
	purego.RegisterLibFunc(&iotaSDKStruct.lib_destroyString, iotaSDK, "destroy_string")

	purego.RegisterLibFunc(&iotaSDKStruct.lib_callClientMethod, iotaSDK, "call_client_method")
	purego.RegisterLibFunc(&iotaSDKStruct.lib_callWalletMethod, iotaSDK, "call_wallet_method")
	purego.RegisterLibFunc(&iotaSDKStruct.lib_callSecretManagerMethod, iotaSDK, "call_secret_manager_method")

	purego.RegisterLibFunc(&iotaSDKStruct.lib_getClientFromWallet, iotaSDK, "get_client_from_wallet")
	purego.RegisterLibFunc(&iotaSDKStruct.lib_getSecretManagerFromWallet, iotaSDK, "get_secret_manager_from_wallet")

	purego.RegisterLibFunc(&iotaSDKStruct.lib_getLastError, iotaSDK, "binding_get_last_error")

	return &iotaSDKStruct
}

func (i *IOTASDK) GetLastError() error {
	result := i.lib_getLastError()
	if len(result) == 0 {
		return nil
	}

	return errors.New(result)
}

func (i *IOTASDK) InitLogger(loggerConfig types.ILoggerConfig) (bool, error) {
	msg, err := serialize(loggerConfig)
	if err != nil {
		return false, err
	}

	if !i.lib_initLogger(msg) {
		return false, i.GetLastError()
	}

	return true, nil
}

func (i *IOTASDK) CreateClient(clientOptions types.IClientOptions) (clientPtr IotaClientPtr, err error) {
	msg, err := serialize(clientOptions)
	if err != nil {
		return 0, err
	}

	if clientPtr = i.lib_createClient(msg); clientPtr == 0 {
		return 0, i.GetLastError()
	}

	return clientPtr, nil
}

func (i *IOTASDK) CreateWallet(walletOptions types.WalletOptions) (wallet *Wallet, err error) {
	msg, err := serialize(walletOptions)
	if err != nil {
		return nil, err
	}

	walletPtr := IotaWalletPtr(0)
	if walletPtr = i.lib_createWallet(msg); walletPtr == 0 {
		return nil, i.GetLastError()
	}

	clientPtr, err := i.GetClientFromWallet(walletPtr)
	if err != nil {
		return nil, err
	}

	secretManagerPtr, err := i.GetSecretManagerFromWallet(walletPtr)
	if err != nil {
		return nil, err
	}

	return NewWallet(i, walletPtr, clientPtr, secretManagerPtr), nil
}

func (i *IOTASDK) CreateSecretManager(secretManagerOptions types.WalletOptionsSecretManager) (clientPtr IotaSecretManagerPtr, err error) {
	msg, err := serialize(secretManagerOptions)
	if err != nil {
		return 0, err
	}

	if clientPtr = i.lib_createSecretManager(msg); clientPtr == 0 {
		return 0, i.GetLastError()
	}

	return clientPtr, nil
}

func (i *IOTASDK) GetClientFromWallet(iotaWalletPtr IotaWalletPtr) (clientPtr IotaClientPtr, err error) {
	if clientPtr = i.lib_getClientFromWallet(iotaWalletPtr); clientPtr == 0 {
		return 0, i.GetLastError()
	}

	return clientPtr, nil
}

func (i *IOTASDK) GetSecretManagerFromWallet(iotaWalletPtr IotaWalletPtr) (secretManagerPtr IotaSecretManagerPtr, err error) {
	if secretManagerPtr = i.lib_getSecretManagerFromWallet(iotaWalletPtr); secretManagerPtr == 0 {
		return 0, i.GetLastError()
	}

	return secretManagerPtr, nil
}

func (i *IOTASDK) CallClientMethod(iotaClientPtr IotaClientPtr, method any) (response string, err error) {
	msg, err := serialize(method)
	if err != nil {
		return "", err
	}

	if response = i.lib_callClientMethod(iotaClientPtr, msg); len(response) == 0 {
		return "", i.GetLastError()
	}

	return response, nil
}

func (i *IOTASDK) CallWalletMethod(iotaWalletPtr IotaWalletPtr, method any) (response string, err error) {
	msg, err := serialize(method)
	if err != nil {
		return "", err
	}

	if response = i.lib_callWalletMethod(iotaWalletPtr, msg); len(response) == 0 {
		return "", i.GetLastError()
	}

	return response, nil
}

func (i *IOTASDK) CallSecretManagerMethod(iotaSecretManagerPtr IotaSecretManagerPtr, method any) (response string, err error) {
	msg, err := serialize(method)
	if err != nil {
		return "", err
	}

	if response = i.lib_callSecretManagerMethod(iotaSecretManagerPtr, msg); len(response) == 0 {
		return "", i.GetLastError()
	}

	return response, nil
}

func (i *IOTASDK) DestroyClient(client IotaClientPtr) (err error) {
	if success := i.lib_destroyClient(client); !success {
		return i.GetLastError()
	}

	return nil
}

func (i *IOTASDK) DestroyWallet(client IotaWalletPtr) (err error) {
	if success := i.lib_destroyWallet(client); !success {
		return i.GetLastError()
	}

	return nil
}

func (i *IOTASDK) DestroySecretManager(client IotaSecretManagerPtr) (err error) {
	if success := i.lib_destroySecretManager(client); !success {
		return i.GetLastError()
	}

	return nil
}

type ResponseEnvelope struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func ParseResponse[T any](responseString string, err error) (*T, error) {
	if err != nil {
		return nil, err
	}

	responseEnvelope := ResponseEnvelope{}
	if err = json.Unmarshal([]byte(responseString), &responseEnvelope); err != nil {
		return nil, err
	}

	response := new(T)
	if err = json.Unmarshal(responseEnvelope.Payload, response); err != nil {
		return nil, err
	}

	return response, err
}
