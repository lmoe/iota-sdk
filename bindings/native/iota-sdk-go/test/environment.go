package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"iota_sdk_go"
)

const ShimmerNetworkAPI = "https://api.shimmer.network"

// chosen by fair dice roll.
// guaranteed to be random.
const TestMnemonic = "saddle dune lake festival gain race cancel fragile amused brush donor outer today unique actress rescue abstract curve tail find catch huge cricket crop"

func InitTest(t *testing.T) *iota_sdk_go.IOTASDK {
	sdk := iota_sdk_go.NewIotaSDK()

	success, err := sdk.InitLogger(iota_sdk_go.ILoggerConfig{
		LevelFilter: iota_sdk_go.ILoggerConfigLevelFilterTrace,
	})
	require.True(t, success)
	require.NoError(t, err)

	return sdk
}
