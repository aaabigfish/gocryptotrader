package bitget

import (
	"github.com/aaabigfish/gocryptotrader/exchanges/account"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// doc https://www.bitget.com/zh-CN/api-doc/spot/account/Get-Account-Assets
const (
	BaseUrl       = "https://api.bitget.com"
	WsUrl         = "wss://ws.bitget.com/mix/v1/stream"
	TimeoutSecond = 30
	SignType      = SHA256

	ContentType        = "Content-Type"
	BgAccessKey        = "ACCESS-KEY"
	BgAccessSign       = "ACCESS-SIGN"
	BgAccessTimestamp  = "ACCESS-TIMESTAMP"
	BgAccessPassphrase = "ACCESS-PASSPHRASE"
	ApplicationJson    = "application/json"

	EN_US  = "en_US"
	ZH_CN  = "zh_CN"
	LOCALE = "locale="

	/*
	 * http methods
	 */
	GET  = "GET"
	POST = "POST"

	/*
	 * websocket
	 */
	WsAuthMethod        = "GET"
	WsAuthPath          = "/user/verify"
	WsOpLogin           = "login"
	WsOpUnsubscribe     = "unsubscribe"
	WsOpSubscribe       = "subscribe"
	TimerIntervalSecond = 5
	ReconnectWaitSecond = 60

	/*
	 * SignType
	 */
	RSA    = "RSA"
	SHA256 = "SHA256"
)

type BitgetRestClient struct {
	ApiKey       string
	ApiSecretKey string
	Passphrase   string
	BaseUrl      string
	HttpClient   http.Client
	Signer       *Signer
}

func (p *BitgetRestClient) Init() *BitgetRestClient {
	p.BaseUrl = BaseUrl
	p.HttpClient = http.Client{
		Timeout: time.Duration(TimeoutSecond) * time.Second,
	}
	return p
}
func (p *BitgetRestClient) Credentials(cre *account.Credentials) {
	p.ApiKey = cre.Key
	p.ApiSecretKey = cre.Secret
	p.Passphrase = cre.ClientID
	p.Signer = new(Signer).Init(cre.Secret)
}

func (p *BitgetRestClient) DoPost(uri string, params string) ([]byte, error) {
	timesStamp := TimesStamp()
	//body, _ := internal.BuildJsonParams(params)

	sign := p.Signer.Sign(POST, uri, params, timesStamp)
	if RSA == SignType {
		sign = p.Signer.SignByRSA(POST, uri, params, timesStamp)
	}
	requestUrl := BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(POST, requestUrl, buffer)

	Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)
	if err != nil {
		return nil, err
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyStr, err
}

func (p *BitgetRestClient) DoGet(uri string, params map[string]string) ([]byte, error) {
	timesStamp := TimesStamp()
	body := BuildGetParams(params)

	requestUrl := p.BaseUrl + uri + body

	request, err := http.NewRequest(GET, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	Headers(request, "", timesStamp, "", "")

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyStr, err
}

func (p *BitgetRestClient) DoGetWithCredentials(uri string, params map[string]string) ([]byte, error) {
	timesStamp := TimesStamp()
	body := BuildGetParams(params)

	sign := p.Signer.Sign(GET, uri, body, timesStamp)

	requestUrl := p.BaseUrl + uri + body

	request, err := http.NewRequest(GET, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyStr, err
}
