package bitget

import "encoding/json"

func (p *BitgetRestClient) Tickers(params map[string]string) ([]TickersData, error) {
	resp, err := p.DoGet("/api/v2/spot/market/tickers", params)
	var obj TickersResp
	err = json.Unmarshal(resp, &obj)
	if err != nil {
		return nil, err
	}
	return obj.Data, err
}

func (p *BitgetRestClient) Orderbook(params map[string]string) (*OrderbookData, error) {
	resp, err := p.DoGet("/api/v2/spot/market/orderbook", params)
	var obj OrderbookResp
	err = json.Unmarshal(resp, &obj)
	if err != nil {
		return nil, err
	}
	return &obj.Data, err
}

func (p *BitgetRestClient) Symbols(params map[string]string) (string, error) {
	resp, err := p.DoGet("/api/v2/spot/public/symbols", params)
	return string(resp), err
}

// ########### auth

func (p *BitgetRestClient) Assets(params map[string]string) ([]AssetsData, error) {
	resp, err := p.DoGetWithCredentials("/api/v2/spot/account/assets", params)
	var obj AssetsResp
	err = json.Unmarshal(resp, &obj)
	if err != nil {
		return nil, err
	}
	return obj.Data, err
}

func (p *BitgetRestClient) BatchPlaceOrder(params map[string]interface{}) (*BatchOrdersData, error) {
	postBody, jsonErr := ToJson(params)
	if jsonErr != nil {
		return nil, jsonErr
	}
	resp, err := p.DoPost("/api/v2/spot/trade/batch-orders", postBody)
	var obj BatchOrdersResp
	err = json.Unmarshal(resp, &obj)
	if err != nil {
		return nil, err
	}
	return &obj.Data, err
}

func (p *BitgetRestClient) CancelOrder(params map[string]string) (string, error) {
	postBody, jsonErr := ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.DoPost("/api/v2/spot/trade/cancel-order", postBody)
	return string(resp), err
}

func (p *BitgetRestClient) BatchCancelOrders(params map[string]interface{}) (*BatchOrdersData, error) {
	postBody, jsonErr := ToJson(params)
	if jsonErr != nil {
		return nil, jsonErr
	}
	resp, err := p.DoPost("/api/v2/spot/trade/batch-cancel-order", postBody)
	var obj BatchOrdersResp
	err = json.Unmarshal(resp, &obj)
	if err != nil {
		return nil, err
	}
	return &obj.Data, err
}

// 查询历史成交
func (p *BitgetRestClient) Fills(params map[string]string) ([]FillsData, error) {
	resp, err := p.DoGetWithCredentials("/api/v2/spot/trade/fills", params)
	var obj FillsResp
	err = json.Unmarshal(resp, &obj)
	if err != nil {
		return nil, err
	}
	return obj.Data, err
}

func (p *BitgetRestClient) OrdersPending(params map[string]string) ([]UnfilledOrdersData, error) {
	resp, err := p.DoGetWithCredentials("/api/v2/spot/trade/unfilled-orders", params)
	var obj UnfilledOrdersResp
	err = json.Unmarshal(resp, &obj)
	if err != nil {
		return nil, err
	}
	return obj.Data, err
}
