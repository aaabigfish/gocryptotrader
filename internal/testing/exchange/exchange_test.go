package exchange

import (
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/aaabigfish/gocryptotrader/config"
	"github.com/aaabigfish/gocryptotrader/exchanges/binance"
	"github.com/aaabigfish/gocryptotrader/exchanges/sharedtestvalues"
)

// TestSetup exercises Setup
func TestSetup(t *testing.T) {
	b := new(binance.Binance)
	require.NoError(t, Setup(b), "Setup must not error")
	assert.NotNil(t, b.Websocket, "Websocket should not be nil after Setup")

	e := new(sharedtestvalues.CustomEx)
	assert.ErrorIs(t, Setup(e), config.ErrExchangeNotFound, "Setup should error correctly on a missing exchange")
}

// TestMockHTTPInstance exercises MockHTTPInstance
func TestMockHTTPInstance(t *testing.T) {
	b := new(binance.Binance)
	require.NoError(t, Setup(b), "Test exchange Setup must not error")
	require.NoError(t, MockHTTPInstance(b), "MockHTTPInstance must not error")
}

// TestMockWsInstance exercises MockWsInstance
func TestMockWsInstance(t *testing.T) {
	b := MockWsInstance[binance.Binance](t, CurryWsMockUpgrader(t, func(_ []byte, _ *websocket.Conn) error { return nil }))
	require.NotNil(t, b, "MockWsInstance must not be nil")
}
