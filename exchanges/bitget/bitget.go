package bitget

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"strconv"
	"strings"
	"time"

	"github.com/aaabigfish/gocryptotrader/common"
	"github.com/aaabigfish/gocryptotrader/config"
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
)

// Bitget creates a mock custom exchange
type Bitget struct {
	exchange.Base
	nativeClient BitgetRestClient
}

// Setup is a mock method for Bitget
func (b *Bitget) Setup(_ *config.Exchange) error {
	return nil
}

// SetDefaults is a mock method for Bitget
func (b *Bitget) SetDefaults() {
	b.Name = "bitget"
	b.Enabled = true
	b.nativeClient.Init()
}

// GetName is a mock method for Bitget
func (b *Bitget) GetName() string {
	return b.Name
}

// IsEnabled is a mock method for Bitget
func (b *Bitget) IsEnabled() bool {
	return true
}

// SetEnabled is a mock method for Bitget
func (b *Bitget) SetEnabled(bool) {
}

// ValidateAPICredentials is a mock method for Bitget
func (b *Bitget) ValidateAPICredentials(_ context.Context, _ asset.Item) error {
	return nil
}

// FetchTicker is a mock method for Bitget
func (b *Bitget) FetchTicker(_ context.Context, pair currency.Pair, assetType asset.Item) (*ticker.Price, error) {
	param := map[string]string{"symbol": strings.ReplaceAll(pair.String(), pair.Delimiter, "")}
	tickers, err := b.nativeClient.Tickers(param)
	if err != nil {
		return nil, err
	}
	if len(tickers) == 0 {
		return nil, errors.New("ticker not found")
	}
	_ticker := tickers[0]
	_price := &ticker.Price{
		ExchangeName: b.Name,
		Pair:         pair,
		AssetType:    assetType,
		Bid:          stf(_ticker.BidPr),
		Ask:          stf(_ticker.AskPr),
		Open:         stf(_ticker.Open),
		LastUpdated:  time.UnixMilli(int64(stf(_ticker.Ts))),
	}
	return _price, nil
}

// UpdateTickers is a mock method for Bitget
func (b *Bitget) UpdateTickers(ctx context.Context, pair asset.Item) error {
	return nil
}

// UpdateTicker is a mock method for Bitget
func (b *Bitget) UpdateTicker(ctx context.Context, pair currency.Pair, a asset.Item) (*ticker.Price, error) {
	return b.FetchTicker(ctx, pair, a)
}

// FetchOrderbook is a mock method for Bitget
func (b *Bitget) FetchOrderbook(_ context.Context, _ currency.Pair, _ asset.Item) (*orderbook.Base, error) {
	return nil, nil
}
func stf(num string) float64 {
	_n, _ := strconv.ParseFloat(num, 64)
	return _n
}

// UpdateOrderbook is a mock method for Bitget
func (b *Bitget) UpdateOrderbook(_ context.Context, pair currency.Pair, _ asset.Item) (*orderbook.Base, error) {
	param := map[string]string{"symbol": strings.ReplaceAll(pair.String(), pair.Delimiter, ""), "type": "step0"}
	_orderbook, err := b.nativeClient.Orderbook(param)
	if err != nil {
		return nil, err
	}

	ob := &orderbook.Base{
		Pair:        pair,
		Exchange:    b.Name,
		LastUpdated: time.UnixMilli(int64(stf(_orderbook.Ts))),
	}
	for i := range _orderbook.Asks {
		ob.Asks = append(ob.Asks, orderbook.Tranche{
			Price:  stf(_orderbook.Asks[i][0]),
			Amount: stf(_orderbook.Asks[i][1]),
		})
	}
	for i := range _orderbook.Bids {
		ob.Bids = append(ob.Bids, orderbook.Tranche{
			Price:  stf(_orderbook.Bids[i][0]),
			Amount: stf(_orderbook.Bids[i][1]),
		})
	}
	return ob, nil
}

// FetchTradablePairs is a mock method for Bitget
func (b *Bitget) FetchTradablePairs(_ context.Context, _ asset.Item) (currency.Pairs, error) {
	return nil, nil
}

// UpdateTradablePairs is a mock method for Bitget
func (b *Bitget) UpdateTradablePairs(_ context.Context, _ bool) error {
	return nil
}

// GetEnabledPairs is a mock method for Bitget
func (b *Bitget) GetEnabledPairs(_ asset.Item) (currency.Pairs, error) {
	return nil, nil
}

// GetAvailablePairs is a mock method for Bitget
func (b *Bitget) GetAvailablePairs(_ asset.Item) (currency.Pairs, error) {
	return nil, nil
}

// FetchAccountInfo is a mock method for Bitget
func (b *Bitget) FetchAccountInfo(_ context.Context, _ asset.Item) (account.Holdings, error) {
	return account.Holdings{}, nil
}

// UpdateAccountInfo is a mock method for Bitget
func (b *Bitget) UpdateAccountInfo(_ context.Context, _ asset.Item) (account.Holdings, error) {
	creds, err := b.GetCredentials(context.Background())
	if err != nil {
		return account.Holdings{}, err
	}
	b.nativeClient.Credentials(creds)
	_assets, err := b.nativeClient.Assets(nil)
	if err != nil {
		return account.Holdings{}, err
	}
	holdings := account.Holdings{
		Exchange: b.Name,
	}
	subA := account.SubAccount{}
	for _, _a := range _assets {
		subA.Currencies = append(subA.Currencies, account.Balance{
			Currency: currency.NewCode(_a.Coin),
			Free:     stf(_a.Available),
			Hold:     stf(_a.Frozen) + stf(_a.Locked),
		})
	}
	holdings.Accounts = append(holdings.Accounts, subA)
	return holdings, nil
}

// SetPairs is a mock method for Bitget
func (b *Bitget) SetPairs(_ currency.Pairs, _ asset.Item, _ bool) error {
	return nil
}

// GetAssetTypes is a mock method for Bitget
func (b *Bitget) GetAssetTypes(_ bool) asset.Items {
	return nil
}

// GetRecentTrades is a mock method for Bitget
func (b *Bitget) GetRecentTrades(_ context.Context, _ currency.Pair, _ asset.Item) ([]trade.Data, error) {
	return nil, nil
}

// GetHistoricTrades is a mock method for Bitget
func (b *Bitget) GetHistoricTrades(_ context.Context, _ currency.Pair, _ asset.Item, _, _ time.Time) ([]trade.Data, error) {
	return nil, nil
}

// SupportsAutoPairUpdates is a mock method for Bitget
func (b *Bitget) SupportsAutoPairUpdates() bool {
	return false
}

// SupportsRESTTickerBatchUpdates is a mock method for Bitget
func (b *Bitget) SupportsRESTTickerBatchUpdates() bool {
	return false
}

// GetServerTime is a mock method for Bitget
func (b *Bitget) GetServerTime(context.Context, asset.Item) (time.Time, error) {
	return time.Now(), nil
}

// GetFeeByType is a mock method for Bitget
func (b *Bitget) GetFeeByType(_ context.Context, _ *exchange.FeeBuilder) (float64, error) {
	return 0.0, nil
}

// GetLastPairsUpdateTime is a mock method for Bitget
func (b *Bitget) GetLastPairsUpdateTime() int64 {
	return 0
}

// GetWithdrawPermissions is a mock method for Bitget
func (b *Bitget) GetWithdrawPermissions() uint32 {
	return 0
}

// FormatWithdrawPermissions is a mock method for Bitget
func (b *Bitget) FormatWithdrawPermissions() string {
	return ""
}

// SupportsWithdrawPermissions is a mock method for Bitget
func (b *Bitget) SupportsWithdrawPermissions(_ uint32) bool {
	return false
}

// GetAccountFundingHistory is a mock method for Bitget
func (b *Bitget) GetAccountFundingHistory(_ context.Context) ([]exchange.FundingHistory, error) {
	return nil, nil
}

// SubmitOrder is a mock method for Bitget
func (b *Bitget) SubmitOrders(_ context.Context, ss ...*order.Submit) ([]*order.SubmitResponse, error) {
	creds, err := b.GetCredentials(context.Background())
	if err != nil {
		return nil, err
	}
	b.nativeClient.Credentials(creds)
	var clientOIDs []string
	param := make(map[string]interface{})
	param["symbol"] = strings.ReplaceAll(ss[0].Pair.String(), ss[0].Pair.Delimiter, "")
	orderList := make([]map[string]string, 0)
	for _, s := range ss {
		v6, _ := uuid.NewGen().NewV6()
		clientOID := strings.ReplaceAll(v6.String(), "-", "")
		clientOIDs = append(clientOIDs, clientOID)
		_order := make(map[string]string)
		_order["side"] = s.Side.Lower()
		_order["orderType"] = s.Type.Lower()
		_order["clientOid"] = clientOID
		if s.Ioc {
			_order["force"] = "ioc"
		} else {
			_order["force"] = "gtc"
		}
		_order["price"] = fmt.Sprintf("%v", s.Price)
		_order["size"] = fmt.Sprintf("%v", s.Amount)
		orderList = append(orderList, _order)
	}
	param["orderList"] = orderList
	placeOrders, err := b.nativeClient.BatchPlaceOrder(param)
	if err != nil {
		return nil, err
	}
	var res []*order.SubmitResponse
	var errs string
	for i, clientOID := range clientOIDs {
		for _, success := range placeOrders.SuccessList {
			if success.ClientOid == clientOID {
				response, err := ss[i].DeriveSubmitResponse(success.OrderId)
				if err != nil {
					errs += err.Error() + "\t"
					continue
				}
				res = append(res, response)
			}
		}
		for _, fail := range placeOrders.FailureList {
			if fail.ClientOid == clientOID {
				response, err := ss[i].DeriveSubmitResponse(fail.OrderId)
				if err != nil {
					errs += err.Error() + "\t"
					continue
				}
				res = append(res, response)
			}
		}
	}
	return res, nil
}

// SubmitOrder is a mock method for Bitget
func (b *Bitget) SubmitOrder(_ context.Context, _ *order.Submit) (*order.SubmitResponse, error) {
	return nil, nil
}

// ModifyOrder is a mock method for Bitget
func (b *Bitget) ModifyOrder(_ context.Context, _ *order.Modify) (*order.ModifyResponse, error) {
	return nil, nil
}

// CancelOrder is a mock method for Bitget
func (b *Bitget) CancelOrder(_ context.Context, _ *order.Cancel) error {
	return nil
}

// CancelBatchOrders is a mock method for Bitget
func (b *Bitget) CancelBatchOrders(_ context.Context, ss []order.Cancel) (*order.CancelBatchResponse, error) {
	creds, err := b.GetCredentials(context.Background())
	if err != nil {
		return nil, err
	}
	b.nativeClient.Credentials(creds)
	param := make(map[string]interface{})
	param["symbol"] = strings.ReplaceAll(ss[0].Pair.String(), ss[0].Pair.Delimiter, "")
	orderList := make([]map[string]string, 0)
	for _, s := range ss {
		_order := make(map[string]string)
		_order["orderId"] = s.OrderID
		orderList = append(orderList, _order)
	}
	param["orderList"] = orderList
	orders, err := b.nativeClient.BatchCancelOrders(param)
	if err != nil {
		return nil, err
	}
	status := map[string]string{}
	for _, suc := range orders.SuccessList {
		status[suc.OrderId] = "success"
	}
	for _, fail := range orders.FailureList {
		status[fail.OrderId] = "fail"
	}
	res := &order.CancelBatchResponse{Status: status}
	return res, nil
}

// CancelAllOrders is a mock method for Bitget
func (b *Bitget) CancelAllOrders(_ context.Context, _ *order.Cancel) (order.CancelAllResponse, error) {
	return order.CancelAllResponse{}, nil
}

// GetOrderInfo is a mock method for Bitget
func (b *Bitget) GetOrderInfo(_ context.Context, _ string, _ currency.Pair, _ asset.Item) (*order.Detail, error) {
	return nil, nil
}

// GetDepositAddress is a mock method for Bitget
func (b *Bitget) GetDepositAddress(_ context.Context, _ currency.Code, _, _ string) (*deposit.Address, error) {
	return nil, nil
}

// GetOrderHistory is a mock method for Bitget
func (b *Bitget) GetOrderHistory(_ context.Context, req *order.MultiOrderRequest) (order.FilteredOrders, error) {
	creds, err := b.GetCredentials(context.Background())
	if err != nil {
		return nil, err
	}
	b.nativeClient.Credentials(creds)
	param := make(map[string]string)
	if len(req.Pairs) > 0 {
		param["symbol"] = strings.ReplaceAll(req.Pairs[0].String(), req.Pairs[0].Delimiter, "")
	}
	fills, err := b.nativeClient.Fills(param)
	if err != nil {
		return nil, err
	}
	var res order.FilteredOrders
	for _, fill := range fills {
		od := order.Detail{
			Exchange:             b.Name,
			Pair:                 req.Pairs[0],
			AccountID:            fill.UserId,
			Price:                stf(fill.PriceAvg),
			Amount:               stf(fill.Size),
			OrderID:              fill.OrderId,
			Fee:                  -stf(fill.FeeDetail.TotalFee),
			AverageExecutedPrice: stf(fill.PriceAvg),
		}
		od.Side, _ = order.StringToOrderSide(fill.Side)
		od.Type, _ = order.StringToOrderType(fill.OrderType)
		od.LastUpdated = time.UnixMilli(int64(stf(fill.UTime)))
		res = append(res, od)
	}
	return res, nil
}

// GetWithdrawalsHistory is a mock method for Bitget
func (b *Bitget) GetWithdrawalsHistory(_ context.Context, _ currency.Code, _ asset.Item) ([]exchange.WithdrawalHistory, error) {
	return []exchange.WithdrawalHistory{}, nil
}

// GetActiveOrders is a mock method for Bitget
func (b *Bitget) GetActiveOrders(_ context.Context, req *order.MultiOrderRequest) (order.FilteredOrders, error) {
	creds, err := b.GetCredentials(context.Background())
	if err != nil {
		return nil, err
	}
	b.nativeClient.Credentials(creds)
	param := make(map[string]string)
	if len(req.Pairs) > 0 {
		param["symbol"] = strings.ReplaceAll(req.Pairs[0].String(), req.Pairs[0].Delimiter, "")
	}
	fills, err := b.nativeClient.OrdersPending(param)
	if err != nil {
		return nil, err
	}
	var res order.FilteredOrders
	for _, fill := range fills {
		od := order.Detail{
			Pair:          req.Pairs[0],
			AccountID:     fill.UserId,
			Price:         stf(fill.PriceAvg),
			Amount:        stf(fill.Size),
			OrderID:       fill.OrderId,
			ClientOrderID: fill.ClientOid,
		}
		od.Side, _ = order.StringToOrderSide(fill.Side)
		od.Type, _ = order.StringToOrderType(fill.OrderType)
		od.LastUpdated = time.UnixMilli(int64(stf(fill.CTime)))
		od.Status, _ = order.StringToOrderStatus(fill.Status)

		res = append(res, od)
	}
	return res, nil
}

// WithdrawCryptocurrencyFunds is a mock method for Bitget
func (b *Bitget) WithdrawCryptocurrencyFunds(_ context.Context, _ *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	return nil, nil
}

// WithdrawFiatFunds is a mock method for Bitget
func (b *Bitget) WithdrawFiatFunds(_ context.Context, _ *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	return nil, nil
}

// WithdrawFiatFundsToInternationalBank is a mock method for Bitget
func (b *Bitget) WithdrawFiatFundsToInternationalBank(_ context.Context, _ *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	return nil, nil
}

// SetHTTPClientUserAgent is a mock method for Bitget
func (b *Bitget) SetHTTPClientUserAgent(_ string) error {
	return nil
}

// GetHTTPClientUserAgent is a mock method for Bitget
func (b *Bitget) GetHTTPClientUserAgent() (string, error) {
	return "", nil
}

// SetClientProxyAddress is a mock method for Bitget
func (b *Bitget) SetClientProxyAddress(_ string) error {
	return nil
}

// SupportsREST is a mock method for Bitget
func (b *Bitget) SupportsREST() bool {
	return true
}

// GetSubscriptions is a mock method for Bitget
func (b *Bitget) GetSubscriptions() ([]subscription.Subscription, error) {
	return nil, nil
}

// SupportsAsset is a mock method for Bitget
func (b *Bitget) SupportsAsset(_ asset.Item) bool {
	return false
}

// GetHistoricCandles is a mock method for Bitget
func (b *Bitget) GetHistoricCandles(_ context.Context, _ currency.Pair, _ asset.Item, _ kline.Interval, _, _ time.Time) (*kline.Item, error) {
	return &kline.Item{}, nil
}

// GetHistoricCandlesExtended is a mock method for Bitget
func (b *Bitget) GetHistoricCandlesExtended(_ context.Context, _ currency.Pair, _ asset.Item, _ kline.Interval, _, _ time.Time) (*kline.Item, error) {
	return &kline.Item{}, nil
}

// DisableRateLimiter is a mock method for Bitget
func (b *Bitget) DisableRateLimiter() error {
	return nil
}

// EnableRateLimiter is a mock method for Bitget
func (b *Bitget) EnableRateLimiter() error {
	return nil
}

// GetWebsocket is a mock method for Bitget
func (b *Bitget) GetWebsocket() (*stream.Websocket, error) {
	return nil, nil
}

// IsWebsocketEnabled is a mock method for Bitget
func (b *Bitget) IsWebsocketEnabled() bool {
	return false
}

// SupportsWebsocket is a mock method for Bitget
func (b *Bitget) SupportsWebsocket() bool {
	return false
}

// SubscribeToWebsocketChannels is a mock method for Bitget
func (b *Bitget) SubscribeToWebsocketChannels(_ []subscription.Subscription) error {
	return nil
}

// UnsubscribeToWebsocketChannels is a mock method for Bitget
func (b *Bitget) UnsubscribeToWebsocketChannels(_ []subscription.Subscription) error {
	return nil
}

// IsAssetWebsocketSupported is a mock method for Bitget
func (b *Bitget) IsAssetWebsocketSupported(_ asset.Item) bool {
	return false
}

// FlushWebsocketChannels is a mock method for Bitget
func (b *Bitget) FlushWebsocketChannels() error {
	return nil
}

// AuthenticateWebsocket is a mock method for Bitget
func (b *Bitget) AuthenticateWebsocket(_ context.Context) error {
	return nil
}

// GetOrderExecutionLimits is a mock method for Bitget
func (b *Bitget) GetOrderExecutionLimits(_ asset.Item, _ currency.Pair) (order.MinMaxLevel, error) {
	return order.MinMaxLevel{}, nil
}

// CheckOrderExecutionLimits is a mock method for Bitget
func (b *Bitget) CheckOrderExecutionLimits(_ asset.Item, _ currency.Pair, _, _ float64, _ order.Type) error {
	return nil
}

// UpdateOrderExecutionLimits is a mock method for Bitget
func (b *Bitget) UpdateOrderExecutionLimits(_ context.Context, _ asset.Item) error {
	return nil
}

// GetHistoricalFundingRates returns funding rates for a given asset and currency for a time period
func (b *Bitget) GetHistoricalFundingRates(_ context.Context, _ *fundingrate.HistoricalRatesRequest) (*fundingrate.HistoricalRates, error) {
	return nil, nil
}

// GetLatestFundingRates returns the latest funding rates data
func (b *Bitget) GetLatestFundingRates(_ context.Context, _ *fundingrate.LatestRateRequest) ([]fundingrate.LatestRateResponse, error) {
	return nil, nil
}

// GetFuturesContractDetails returns all contracts from the exchange by asset type
func (b *Bitget) GetFuturesContractDetails(context.Context, asset.Item) ([]futures.Contract, error) {
	return nil, common.ErrFunctionNotSupported
}
