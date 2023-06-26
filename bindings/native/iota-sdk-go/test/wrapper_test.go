package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"iota_sdk_go"
)

const ShimmerNetworkAPI = "https://api.shimmer.network"

func TestFirstCon(t *testing.T) {
	sdk := iota_sdk_go.NewIotaSDK()

	success, err := sdk.InitLogger(iota_sdk_go.ILoggerConfig{
		LevelFilter: iota_sdk_go.ILoggerConfigLevelFilterTrace,
	})
	require.True(t, success)
	require.NoError(t, err)

	clientPtr, err := sdk.CreateClient(iota_sdk_go.IClientOptions{
		PrimaryNode: ShimmerNetworkAPI,
		Nodes:       []interface{}{ShimmerNetworkAPI},
	})
	require.NoError(t, err)
	require.NotNil(t, clientPtr)

	response, err := sdk.CallClientMethod(clientPtr, "RUBBISH")
	require.Empty(t, response)
	require.Error(t, err)

	t.Log(clientPtr)
}
