package sharedtestvalues

import (
	"context"
	"github.com/aaabigfish/gocryptotrader/currency"
	"github.com/aaabigfish/gocryptotrader/exchanges/asset"
	"testing"
)

func TestUpdateTicker(t *testing.T) {
	cc := CryptoCom{}

	mxcPair, _ := currency.NewPairFromString("MXC_USDT")

	t.Log(cc.UpdateTicker(context.Background(), mxcPair, asset.Spot))
}
