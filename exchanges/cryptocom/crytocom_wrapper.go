package sharedtestvalues

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aaabigfish/gocryptotrader/currency"
	"github.com/aaabigfish/gocryptotrader/exchanges/asset"
	"github.com/aaabigfish/gocryptotrader/exchanges/ticker"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// GetName is a mock method for CryptoCom
func (c *CryptoCom) GetName() string {
	return c.Name
}

//func (c *CryptoCom) sendAuthenticatedRequest(method, endpoint string, params url.Values) ([]byte, error) {
//	params.Set("sig", c.signRequest(params))
//	headers := make(map[string]string)
//	var err error
//	credentials, err := c.GetCredentials(context.Background())
//	if err != nil {
//		return nil, err
//	}
//	headers["X-MBX-APIKEY"] = credentials.Key
//	url := apiBaseURL + endpoint + "?" + params.Encode()
//
//	var resp []byte
//	switch method {
//	case http.MethodGet:
//		resp, err = request.NewHTTPRequest().Get(url, headers)
//	case http.MethodPost:
//		resp, err = request.NewHTTPRequest().Post(url, headers)
//	}
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}

//func (c *CryptoCom) UpdateAccountInfo() (*AccountInfo, error) {
//	params := url.Values{}
//	resp, err := c.sendAuthenticatedRequest(http.MethodPost, "private/get-account-summary", params)
//	if err != nil {
//		return nil, err
//	}
//	var accountResponse AccountInfo
//	if err := json.Unmarshal(resp, &accountResponse); err != nil {
//		return nil, err
//	}
//	return &accountResponse, nil
//}

func (c *CryptoCom) UpdateTicker(_ context.Context, p currency.Pair, a asset.Item) (*ticker.Price, error) {

	baseUrl := "https://api.crypto.com/exchange/v1/public/get-tickers?"
	params := url.Values{}
	quoteName := p.Quote.Upper().String()
	if quoteName == "USDT" {
		quoteName = "USD"
	}
	params.Add("instrument_name", fmt.Sprintf("%v_%v", p.Base.Upper().String(), quoteName))
	queryUrl := baseUrl + params.Encode()
	stf := func(n string) float64 {
		_float, _ := strconv.ParseFloat(n, 64)
		return _float
	}
	switch a {
	case asset.Spot, asset.Margin:
		client := &http.Client{}
		req, err := http.NewRequest("GET", queryUrl, nil)
		if err != nil {
			return nil, err
		}
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var gt GetTickers
		if err = json.Unmarshal(body, &gt); err != nil {
			return nil, err
		}
		if len(gt.Result.Data) == 0 {
			return nil, errors.New("no data found")
		}
		return &ticker.Price{
			High:         stf(gt.Result.Data[0].H),
			Low:          stf(gt.Result.Data[0].L),
			Bid:          stf(gt.Result.Data[0].B),
			Ask:          stf(gt.Result.Data[0].A),
			Volume:       stf(gt.Result.Data[0].V),
			QuoteVolume:  stf(gt.Result.Data[0].Vv),
			Open:         0,
			Close:        0,
			Pair:         p,
			ExchangeName: c.Name,
			LastUpdated:  time.UnixMilli(gt.Result.Data[0].T),
			AssetType:    a,
		}, nil
	default:
		return nil, errors.New("unknown asset type")
	}
}

//// UpdateOrderbook is a mock method for CryptoCom
//func (c *CryptoCom) UpdateOrderbook(_ context.Context, _ currency.Pair, _ asset.Item) (*orderbook.Base, error) {
//	url := "https://api.crypto.com/v2/public/get-ticker"
//	params := map[string]interface{}{"instrument_name": "BTC_USDT"}
//	resp, err := c.makeRequest("public/get-ticker", url, params)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//	fmt.Println("Response:", string(resp))
//	return nil, nil
//}

//
//func (c *CryptoCom) PlaceOrders(orders []Order) error {
//	params := url.Values{}
//	orderData, err := json.Marshal(orders)
//	if err != nil {
//		return err
//	}
//	resp, err := c.sendAuthenticatedRequest(http.MethodPost, "private/batch-order", params)
//	if err != nil {
//		return err
//	}
//	// Handle response
//	return nil
//}
