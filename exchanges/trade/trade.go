package trade

import (
	"time"
)

// Setup configures necessary fields to the `Trade` structure that govern trade data
// processing.
func (t *Trade) Setup(exchangeName string, tradeFeedEnabled bool, c chan interface{}) {
	t.exchangeName = exchangeName
	t.dataHandler = c
	t.tradeFeedEnabled = tradeFeedEnabled
}

// Update processes trade data, either by saving it or routing it through
// the data channel.
func (t *Trade) Update(save bool, data ...Data) error {
	if len(data) == 0 {
		// nothing to do
		return nil
	}

	if t.tradeFeedEnabled {
		t.dataHandler <- data
	}

	if save {
		if err := AddTradesToBuffer(t.exchangeName, data...); err != nil {
			return err
		}
	}

	return nil
}

// AddTradesToBuffer will push trade data onto the buffer
func AddTradesToBuffer(exchangeName string, data ...Data) error {
	return nil
}

// FilterTradesByTime removes any trades that are not between the start
// and end times
func FilterTradesByTime(trades []Data, startTime, endTime time.Time) []Data {
	if startTime.IsZero() || endTime.IsZero() {
		// can't filter without boundaries
		return trades
	}
	var filteredTrades []Data
	for i := range trades {
		if trades[i].Timestamp.After(startTime) && trades[i].Timestamp.Before(endTime) {
			filteredTrades = append(filteredTrades, trades[i])
		}
	}

	return filteredTrades
}
