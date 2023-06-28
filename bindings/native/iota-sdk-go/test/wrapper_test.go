package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"iota_sdk_go/types"
)

func TestFirstCon(t *testing.T) {
	sdk := InitTest(t)

	clientPtr, err := sdk.CreateClient(types.IClientOptions{
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
