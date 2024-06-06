package bitget

type TickersResp struct {
	Code        string        `json:"code"`
	Msg         string        `json:"msg"` //success
	RequestTime int64         `json:"requestTime"`
	Data        []TickersData `json:"data"`
}

type TickersData struct {
	Symbol       string `json:"symbol"`
	High24H      string `json:"high24h"`
	Open         string `json:"open"`
	Low24H       string `json:"low24h"`
	LastPr       string `json:"lastPr"`
	QuoteVolume  string `json:"quoteVolume"`
	BaseVolume   string `json:"baseVolume"`
	UsdtVolume   string `json:"usdtVolume"`
	BidPr        string `json:"bidPr"`
	AskPr        string `json:"askPr"`
	BidSz        string `json:"bidSz"`
	AskSz        string `json:"askSz"`
	OpenUtc      string `json:"openUtc"`
	Ts           string `json:"ts"`
	ChangeUtc24H string `json:"changeUtc24h"`
	Change24H    string `json:"change24h"`
}

type OrderbookResp struct {
	Code        string        `json:"code"`
	Msg         string        `json:"msg"`
	RequestTime int64         `json:"requestTime"`
	Data        OrderbookData `json:"data"`
}

type OrderbookData struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
	Ts   string     `json:"ts"`
}

type AssetsResp struct {
	Code        string       `json:"code"`
	Message     string       `json:"message"`
	RequestTime int64        `json:"requestTime"`
	Data        []AssetsData `json:"data"`
}

type AssetsData struct {
	Coin           string `json:"coin"`
	Available      string `json:"available"`
	Frozen         string `json:"frozen"`
	Locked         string `json:"locked"`
	LimitAvailable string `json:"limitAvailable"`
	UTime          string `json:"uTime"`
}

type BatchOrdersResp struct {
	Code        string          `json:"code"`
	Msg         string          `json:"msg"`
	RequestTime int64           `json:"requestTime"`
	Data        BatchOrdersData `json:"data"`
}

type BatchOrdersData struct {
	SuccessList []struct {
		OrderId   string `json:"orderId"`
		ClientOid string `json:"clientOid"`
	} `json:"successList"`
	FailureList []struct {
		OrderId   string `json:"orderId"`
		ClientOid string `json:"clientOid"`
		ErrorMsg  string `json:"errorMsg"`
	} `json:"failureList"`
}

type FillsResp struct {
	Code        string      `json:"code"`
	Msg         string      `json:"msg"`
	RequestTime int64       `json:"requestTime"`
	Data        []FillsData `json:"data"`
}

type FillsData struct {
	UserId    string `json:"userId"`
	Symbol    string `json:"symbol"`
	OrderId   string `json:"orderId"`
	TradeId   string `json:"tradeId"`
	OrderType string `json:"orderType"`
	Side      string `json:"side"`
	PriceAvg  string `json:"priceAvg"`
	Size      string `json:"size"`
	Amount    string `json:"amount"`
	FeeDetail struct {
		Deduction         string `json:"deduction"`
		FeeCoin           string `json:"feeCoin"`
		TotalDeductionFee string `json:"totalDeductionFee"`
		TotalFee          string `json:"totalFee"`
	} `json:"feeDetail"`
	TradeScope string `json:"tradeScope"`
	CTime      string `json:"cTime"`
	UTime      string `json:"uTime"`
}

type UnfilledOrdersResp struct {
	Code        string               `json:"code"`
	Message     string               `json:"message"`
	RequestTime int64                `json:"requestTime"`
	Data        []UnfilledOrdersData `json:"data"`
}

type UnfilledOrdersData struct {
	UserId           string      `json:"userId"`
	Symbol           string      `json:"symbol"`
	OrderId          string      `json:"orderId"`
	ClientOid        string      `json:"clientOid"`
	PriceAvg         string      `json:"priceAvg"`
	Size             string      `json:"size"`
	OrderType        string      `json:"orderType"`
	Side             string      `json:"side"`
	Status           string      `json:"status"`
	BasePrice        string      `json:"basePrice"`
	BaseVolume       string      `json:"baseVolume"`
	QuoteVolume      string      `json:"quoteVolume"`
	EnterPointSource string      `json:"enterPointSource"`
	TriggerPrice     interface{} `json:"triggerPrice"`
	TpslType         string      `json:"tpslType"`
	CTime            string      `json:"cTime"`
}
