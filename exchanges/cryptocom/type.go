package sharedtestvalues

type GetTickers struct {
	Id     int    `json:"id"`
	Method string `json:"method"`
	Code   int    `json:"code"`
	Result struct {
		Data []struct {
			I  string `json:"i"`
			H  string `json:"h"`
			L  string `json:"l"`
			A  string `json:"a"`
			V  string `json:"v"`
			Vv string `json:"vv"`
			C  string `json:"c"`
			B  string `json:"b"`
			K  string `json:"k"`
			Oi string `json:"oi"`
			T  int64  `json:"t"`
		} `json:"data"`
	} `json:"result"`
}
