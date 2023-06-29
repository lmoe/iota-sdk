package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"iota_sdk_go"
	"iota_sdk_go/types"

	iotago "github.com/iotaledger/iota.go/v3"
)

const ShimmerNetworkAPI = "https://api.shimmer.network"
const UseLedgerSimulator = true

// chosen by fair dice roll.
// guaranteed to be random.
const TestMnemonic = "saddle dune lake festival gain race cancel fragile amused brush donor outer today unique actress rescue abstract curve tail find catch huge cricket crop"

var TestSignMessageFromEssenceHex = "0xcf30a3824d6b2d3b25ec63aa97733e4fc4dd99e6d38c97093a0abd21f5e9016c"
var TestSignMessageFromEssence, _ = iotago.DecodeHex(TestSignMessageFromEssenceHex)

func InitTest(t *testing.T) *iota_sdk_go.IOTASDK {
	sdk := iota_sdk_go.NewIotaSDK()

	success, err := sdk.InitLogger(types.ILoggerConfig{
		LevelFilter: types.ILoggerConfigLevelFilterTrace,
	})
	require.True(t, success)
	require.NoError(t, err)

	return sdk
}
