package sharedtestvalues

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/aaabigfish/gocryptotrader/common"
	"github.com/aaabigfish/gocryptotrader/config"
	exchange "github.com/aaabigfish/gocryptotrader/exchanges"
	"github.com/aaabigfish/gocryptotrader/exchanges/request"
	"github.com/aaabigfish/gocryptotrader/log"
	"io/ioutil"
	"net/http"
	"time"
)

// CryptoCom creates a mock custom exchange
type CryptoCom struct {
	exchange.Base
}

const (
	apiBaseURL = "https://api.crypto.com/v2/"
)

// Setup is a mock method for CryptoCom
func (c *CryptoCom) Setup(_ *config.Exchange) error {
	return nil
}

// SetDefaults is a mock method for CryptoCom
func (c *CryptoCom) SetDefaults() {
	c.Name = "cryptocom"
	c.Enabled = true
	c.Verbose = false

	c.API.CredentialsValidator.RequiresKey = true
	c.API.CredentialsValidator.RequiresSecret = true
	c.API.CredentialsValidator.RequiresClientID = true

	var err error
	c.Requester, err = request.New(c.Name,
		common.NewHTTPClientWithTimeout(exchange.DefaultHTTPTimeout))
	if err != nil {
		log.Errorln(log.ExchangeSys, err)
	}
}

// IsEnabled is a mock method for CryptoCom
func (c *CryptoCom) IsEnabled() bool {
	return true
}

// SetEnabled is a mock method for CryptoCom
func (c *CryptoCom) SetEnabled(bool) {
}

// SendHTTPRequest sends an unauthenticated request
func (c *CryptoCom) SendHTTPRequest(ctx context.Context, ePath exchange.URL, path string, f request.EndpointLimit, result interface{}) error {
	endpointPath, err := c.API.Endpoints.GetURL(ePath)
	if err != nil {
		return err
	}
	item := &request.Item{
		Method:        http.MethodGet,
		Path:          endpointPath + path,
		Result:        result,
		Verbose:       c.Verbose,
		HTTPDebugging: c.HTTPDebugging,
		HTTPRecording: c.HTTPRecording}

	return c.SendPayload(ctx, f, func() (*request.Item, error) {
		return item, nil
	}, request.UnauthenticatedRequest)
}

func (c *CryptoCom) signRequest(method string, params map[string]interface{}) string {
	var apiKey, apiSecret string
	credentials, err := c.GetCredentials(context.Background())
	if err == nil && credentials != nil {
		apiKey = credentials.Key
		apiSecret = credentials.Secret
	}

	nonce := time.Now().UnixNano() / int64(time.Millisecond)
	paramsString := ""
	for k, v := range params {
		paramsString += k + fmt.Sprintf("%v", v)
	}
	payload := method + apiKey + paramsString + fmt.Sprintf("%d", nonce)
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *CryptoCom) makeRequest(method, url string, params map[string]interface{}) ([]byte, error) {
	var apiKey string
	credentials, err := c.GetCredentials(context.Background())
	if err == nil && credentials != nil {
		apiKey = credentials.Key
	}
	signature := c.signRequest(method, params)
	params["sig"] = signature
	params["api_key"] = apiKey
	params["nonce"] = time.Now().UnixNano() / int64(time.Millisecond)
	body, _ := json.Marshal(params)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
