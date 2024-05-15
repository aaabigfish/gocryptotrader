package modules

import (
	"context"
	"time"

	"github.com/aaabigfish/gocryptotrader/currency"
	"github.com/aaabigfish/gocryptotrader/exchanges/account"
	"github.com/aaabigfish/gocryptotrader/exchanges/asset"
	"github.com/aaabigfish/gocryptotrader/exchanges/deposit"
	"github.com/aaabigfish/gocryptotrader/exchanges/kline"
	"github.com/aaabigfish/gocryptotrader/exchanges/order"
	"github.com/aaabigfish/gocryptotrader/exchanges/orderbook"
	"github.com/aaabigfish/gocryptotrader/exchanges/ticker"
	"github.com/aaabigfish/gocryptotrader/portfolio/withdraw"
)

const (
	// ErrParameterConvertFailed error to return when type conversion fails
	ErrParameterConvertFailed = "%v failed conversion"
	// ErrParameterWithPositionConvertFailed error to return when a positional conversion fails
	ErrParameterWithPositionConvertFailed = "%v at position %v failed conversion"
)

// Wrapper instance of GCT to use for modules
var Wrapper GCTExchange

// GCTExchange interface requirements
type GCTExchange interface {
	Exchanges(enabledOnly bool) []string
	IsEnabled(exch string) bool
	Orderbook(ctx context.Context, exch string, pair currency.Pair, item asset.Item) (*orderbook.Base, error)
	Ticker(ctx context.Context, exch string, pair currency.Pair, item asset.Item) (*ticker.Price, error)
	Pairs(exch string, enabledOnly bool, item asset.Item) (*currency.Pairs, error)
	QueryOrder(ctx context.Context, exch, orderid string, pair currency.Pair, assetType asset.Item) (*order.Detail, error)
	SubmitOrder(ctx context.Context, submit *order.Submit) (*order.SubmitResponse, error)
	CancelOrder(ctx context.Context, exch, orderid string, pair currency.Pair, item asset.Item) (bool, error)
	AccountInformation(ctx context.Context, exch string, assetType asset.Item) (account.Holdings, error)
	DepositAddress(exch, chain string, currencyCode currency.Code) (*deposit.Address, error)
	WithdrawalFiatFunds(ctx context.Context, bankAccountID string, request *withdraw.Request) (out string, err error)
	WithdrawalCryptoFunds(ctx context.Context, request *withdraw.Request) (out string, err error)
	OHLCV(ctx context.Context, exch string, pair currency.Pair, item asset.Item, start, end time.Time, interval kline.Interval) (*kline.Item, error)
}

// SetModuleWrapper link the wrapper and interface to use for modules
func SetModuleWrapper(wrapper GCTExchange) {
	Wrapper = wrapper
}
