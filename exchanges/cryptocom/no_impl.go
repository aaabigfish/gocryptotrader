package sharedtestvalues

import (
	"context"
	"github.com/aaabigfish/gocryptotrader/common"
	"github.com/aaabigfish/gocryptotrader/currency"
	exchange "github.com/aaabigfish/gocryptotrader/exchanges"
	"github.com/aaabigfish/gocryptotrader/exchanges/account"
	"github.com/aaabigfish/gocryptotrader/exchanges/asset"
	"github.com/aaabigfish/gocryptotrader/exchanges/deposit"
	"github.com/aaabigfish/gocryptotrader/exchanges/fundingrate"
	"github.com/aaabigfish/gocryptotrader/exchanges/futures"
	"github.com/aaabigfish/gocryptotrader/exchanges/kline"
	"github.com/aaabigfish/gocryptotrader/exchanges/order"
	"github.com/aaabigfish/gocryptotrader/exchanges/orderbook"
	"github.com/aaabigfish/gocryptotrader/exchanges/stream"
	"github.com/aaabigfish/gocryptotrader/exchanges/subscription"
	"github.com/aaabigfish/gocryptotrader/exchanges/ticker"
	"github.com/aaabigfish/gocryptotrader/exchanges/trade"
	"github.com/aaabigfish/gocryptotrader/portfolio/withdraw"
	"time"
)

// ValidateAPICredentials is a mock method for CryptoCom
func (c *CryptoCom) ValidateAPICredentials(_ context.Context, _ asset.Item) error {
	return nil
}

// FetchTicker is a mock method for CryptoCom
func (c *CryptoCom) FetchTicker(_ context.Context, _ currency.Pair, _ asset.Item) (*ticker.Price, error) {
	return nil, nil
}

// UpdateTickers is a mock method for CryptoCom
func (c *CryptoCom) UpdateTickers(_ context.Context, _ asset.Item) error {
	return nil
}

// FetchOrderbook is a mock method for CryptoCom
func (c *CryptoCom) FetchOrderbook(_ context.Context, _ currency.Pair, _ asset.Item) (*orderbook.Base, error) {
	return nil, nil
}

// FetchTradablePairs is a mock method for CryptoCom
func (c *CryptoCom) FetchTradablePairs(_ context.Context, _ asset.Item) (currency.Pairs, error) {
	return nil, nil
}

// UpdateTradablePairs is a mock method for CryptoCom
func (c *CryptoCom) UpdateTradablePairs(_ context.Context, _ bool) error {
	return nil
}

// GetEnabledPairs is a mock method for CryptoCom
func (c *CryptoCom) GetEnabledPairs(_ asset.Item) (currency.Pairs, error) {
	return nil, nil
}

// GetAvailablePairs is a mock method for CryptoCom
func (c *CryptoCom) GetAvailablePairs(_ asset.Item) (currency.Pairs, error) {
	return nil, nil
}

// FetchAccountInfo is a mock method for CryptoCom
func (c *CryptoCom) FetchAccountInfo(_ context.Context, _ asset.Item) (account.Holdings, error) {
	return account.Holdings{}, nil
}

// SetPairs is a mock method for CryptoCom
func (c *CryptoCom) SetPairs(_ currency.Pairs, _ asset.Item, _ bool) error {
	return nil
}

// GetAssetTypes is a mock method for CryptoCom
func (c *CryptoCom) GetAssetTypes(_ bool) asset.Items {
	return nil
}

// GetRecentTrades is a mock method for CryptoCom
func (c *CryptoCom) GetRecentTrades(_ context.Context, _ currency.Pair, _ asset.Item) ([]trade.Data, error) {
	return nil, nil
}

// GetHistoricTrades is a mock method for CryptoCom
func (c *CryptoCom) GetHistoricTrades(_ context.Context, _ currency.Pair, _ asset.Item, _, _ time.Time) ([]trade.Data, error) {
	return nil, nil
}

// SupportsAutoPairUpdates is a mock method for CryptoCom
func (c *CryptoCom) SupportsAutoPairUpdates() bool {
	return false
}

// SupportsRESTTickerBatchUpdates is a mock method for CryptoCom
func (c *CryptoCom) SupportsRESTTickerBatchUpdates() bool {
	return false
}

// GetServerTime is a mock method for CryptoCom
func (c *CryptoCom) GetServerTime(context.Context, asset.Item) (time.Time, error) {
	return time.Now(), nil
}

// GetFeeByType is a mock method for CryptoCom
func (c *CryptoCom) GetFeeByType(_ context.Context, _ *exchange.FeeBuilder) (float64, error) {
	return 0.0, nil
}

// GetLastPairsUpdateTime is a mock method for CryptoCom
func (c *CryptoCom) GetLastPairsUpdateTime() int64 {
	return 0
}

// GetWithdrawPermissions is a mock method for CryptoCom
func (c *CryptoCom) GetWithdrawPermissions() uint32 {
	return 0
}

// FormatWithdrawPermissions is a mock method for CryptoCom
func (c *CryptoCom) FormatWithdrawPermissions() string {
	return ""
}

// SupportsWithdrawPermissions is a mock method for CryptoCom
func (c *CryptoCom) SupportsWithdrawPermissions(_ uint32) bool {
	return false
}

// GetAccountFundingHistory is a mock method for CryptoCom
func (c *CryptoCom) GetAccountFundingHistory(_ context.Context) ([]exchange.FundingHistory, error) {
	return nil, nil
}

// SubmitOrder is a mock method for CryptoCom
func (c *CryptoCom) SubmitOrders(_ context.Context, _ ...*order.Submit) ([]*order.SubmitResponse, error) {
	return nil, nil
}

// SubmitOrder is a mock method for CryptoCom
func (c *CryptoCom) SubmitOrder(_ context.Context, _ *order.Submit) (*order.SubmitResponse, error) {
	return nil, nil
}

// ModifyOrder is a mock method for CryptoCom
func (c *CryptoCom) ModifyOrder(_ context.Context, _ *order.Modify) (*order.ModifyResponse, error) {
	return nil, nil
}

// CancelOrder is a mock method for CryptoCom
func (c *CryptoCom) CancelOrder(_ context.Context, _ *order.Cancel) error {
	return nil
}

// CancelBatchOrders is a mock method for CryptoCom
func (c *CryptoCom) CancelBatchOrders(_ context.Context, _ []order.Cancel) (*order.CancelBatchResponse, error) {
	return nil, nil
}

// CancelAllOrders is a mock method for CryptoCom
func (c *CryptoCom) CancelAllOrders(_ context.Context, _ *order.Cancel) (order.CancelAllResponse, error) {
	return order.CancelAllResponse{}, nil
}

// GetOrderInfo is a mock method for CryptoCom
func (c *CryptoCom) GetOrderInfo(_ context.Context, _ string, _ currency.Pair, _ asset.Item) (*order.Detail, error) {
	return nil, nil
}

// GetDepositAddress is a mock method for CryptoCom
func (c *CryptoCom) GetDepositAddress(_ context.Context, _ currency.Code, _, _ string) (*deposit.Address, error) {
	return nil, nil
}

// GetOrderHistory is a mock method for CryptoCom
func (c *CryptoCom) GetOrderHistory(_ context.Context, _ *order.MultiOrderRequest) (order.FilteredOrders, error) {
	return nil, nil
}

// GetWithdrawalsHistory is a mock method for CryptoCom
func (c *CryptoCom) GetWithdrawalsHistory(_ context.Context, _ currency.Code, _ asset.Item) ([]exchange.WithdrawalHistory, error) {
	return []exchange.WithdrawalHistory{}, nil
}

// GetActiveOrders is a mock method for CryptoCom
func (c *CryptoCom) GetActiveOrders(_ context.Context, _ *order.MultiOrderRequest) (order.FilteredOrders, error) {
	return []order.Detail{}, nil
}

// WithdrawCryptocurrencyFunds is a mock method for CryptoCom
func (c *CryptoCom) WithdrawCryptocurrencyFunds(_ context.Context, _ *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	return nil, nil
}

// WithdrawFiatFunds is a mock method for CryptoCom
func (c *CryptoCom) WithdrawFiatFunds(_ context.Context, _ *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	return nil, nil
}

// WithdrawFiatFundsToInternationalBank is a mock method for CryptoCom
func (c *CryptoCom) WithdrawFiatFundsToInternationalBank(_ context.Context, _ *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	return nil, nil
}

// SetHTTPClientUserAgent is a mock method for CryptoCom
func (c *CryptoCom) SetHTTPClientUserAgent(_ string) error {
	return nil
}

// GetHTTPClientUserAgent is a mock method for CryptoCom
func (c *CryptoCom) GetHTTPClientUserAgent() (string, error) {
	return "", nil
}

// SetClientProxyAddress is a mock method for CryptoCom
func (c *CryptoCom) SetClientProxyAddress(_ string) error {
	return nil
}

// SupportsREST is a mock method for CryptoCom
func (c *CryptoCom) SupportsREST() bool {
	return true
}

// GetSubscriptions is a mock method for CryptoCom
func (c *CryptoCom) GetSubscriptions() ([]subscription.Subscription, error) {
	return nil, nil
}

// GetBase is a mock method for CryptoCom
func (c *CryptoCom) GetBase() *exchange.Base {
	return nil
}

// SupportsAsset is a mock method for CryptoCom
func (c *CryptoCom) SupportsAsset(_ asset.Item) bool {
	return false
}

// GetHistoricCandles is a mock method for CryptoCom
func (c *CryptoCom) GetHistoricCandles(_ context.Context, _ currency.Pair, _ asset.Item, _ kline.Interval, _, _ time.Time) (*kline.Item, error) {
	return &kline.Item{}, nil
}

// GetHistoricCandlesExtended is a mock method for CryptoCom
func (c *CryptoCom) GetHistoricCandlesExtended(_ context.Context, _ currency.Pair, _ asset.Item, _ kline.Interval, _, _ time.Time) (*kline.Item, error) {
	return &kline.Item{}, nil
}

// DisableRateLimiter is a mock method for CryptoCom
func (c *CryptoCom) DisableRateLimiter() error {
	return nil
}

// EnableRateLimiter is a mock method for CryptoCom
func (c *CryptoCom) EnableRateLimiter() error {
	return nil
}

// GetWebsocket is a mock method for CryptoCom
func (c *CryptoCom) GetWebsocket() (*stream.Websocket, error) {
	return nil, nil
}

// IsWebsocketEnabled is a mock method for CryptoCom
func (c *CryptoCom) IsWebsocketEnabled() bool {
	return false
}

// SupportsWebsocket is a mock method for CryptoCom
func (c *CryptoCom) SupportsWebsocket() bool {
	return false
}

// SubscribeToWebsocketChannels is a mock method for CryptoCom
func (c *CryptoCom) SubscribeToWebsocketChannels(_ []subscription.Subscription) error {
	return nil
}

// UnsubscribeToWebsocketChannels is a mock method for CryptoCom
func (c *CryptoCom) UnsubscribeToWebsocketChannels(_ []subscription.Subscription) error {
	return nil
}

// IsAssetWebsocketSupported is a mock method for CryptoCom
func (c *CryptoCom) IsAssetWebsocketSupported(_ asset.Item) bool {
	return false
}

// FlushWebsocketChannels is a mock method for CryptoCom
func (c *CryptoCom) FlushWebsocketChannels() error {
	return nil
}

// AuthenticateWebsocket is a mock method for CryptoCom
func (c *CryptoCom) AuthenticateWebsocket(_ context.Context) error {
	return nil
}

// GetOrderExecutionLimits is a mock method for CryptoCom
func (c *CryptoCom) GetOrderExecutionLimits(_ asset.Item, _ currency.Pair) (order.MinMaxLevel, error) {
	return order.MinMaxLevel{}, nil
}

// CheckOrderExecutionLimits is a mock method for CryptoCom
func (c *CryptoCom) CheckOrderExecutionLimits(_ asset.Item, _ currency.Pair, _, _ float64, _ order.Type) error {
	return nil
}

// UpdateOrderExecutionLimits is a mock method for CryptoCom
func (c *CryptoCom) UpdateOrderExecutionLimits(_ context.Context, _ asset.Item) error {
	return nil
}

// GetHistoricalFundingRates returns funding rates for a given asset and currency for a time period
func (c *CryptoCom) GetHistoricalFundingRates(_ context.Context, _ *fundingrate.HistoricalRatesRequest) (*fundingrate.HistoricalRates, error) {
	return nil, nil
}

// GetLatestFundingRates returns the latest funding rates data
func (c *CryptoCom) GetLatestFundingRates(_ context.Context, _ *fundingrate.LatestRateRequest) ([]fundingrate.LatestRateResponse, error) {
	return nil, nil
}

// GetFuturesContractDetails returns all contracts from the exchange by asset type
func (c *CryptoCom) GetFuturesContractDetails(context.Context, asset.Item) ([]futures.Contract, error) {
	return nil, common.ErrFunctionNotSupported
}
