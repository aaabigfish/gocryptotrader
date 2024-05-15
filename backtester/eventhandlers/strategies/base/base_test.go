package base

import (
	"errors"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/aaabigfish/gocryptotrader/backtester/common"
	"github.com/aaabigfish/gocryptotrader/backtester/data"
	datakline "github.com/aaabigfish/gocryptotrader/backtester/data/kline"
	"github.com/aaabigfish/gocryptotrader/backtester/eventtypes/event"
	"github.com/aaabigfish/gocryptotrader/backtester/eventtypes/kline"
	gctcommon "github.com/aaabigfish/gocryptotrader/common"
	"github.com/aaabigfish/gocryptotrader/currency"
	"github.com/aaabigfish/gocryptotrader/exchanges/asset"
	gctkline "github.com/aaabigfish/gocryptotrader/exchanges/kline"
)

func TestGetBase(t *testing.T) {
	t.Parallel()
	s := Strategy{}
	_, err := s.GetBaseData(nil)
	if !errors.Is(err, gctcommon.ErrNilPointer) {
		t.Errorf("received: %v, expected: %v", err, gctcommon.ErrNilPointer)
	}

	_, err = s.GetBaseData(datakline.NewDataFromKline())
	if !errors.Is(err, common.ErrNilEvent) {
		t.Errorf("received: %v, expected: %v", err, common.ErrNilEvent)
	}
	tt := time.Now()
	exch := "binance"
	a := asset.Spot
	p := currency.NewPair(currency.BTC, currency.USDT)
	d := &data.Base{}
	err = d.SetStream([]data.Event{&kline.Kline{
		Base: &event.Base{
			Exchange:     exch,
			Time:         tt,
			Interval:     gctkline.OneDay,
			CurrencyPair: p,
			AssetType:    a,
		},
		Open:   decimal.NewFromInt(1337),
		Close:  decimal.NewFromInt(1337),
		Low:    decimal.NewFromInt(1337),
		High:   decimal.NewFromInt(1337),
		Volume: decimal.NewFromInt(1337),
	}})
	if !errors.Is(err, nil) {
		t.Errorf("received: %v, expected: %v", err, nil)
	}

	_, err = d.Next()
	if !errors.Is(err, nil) {
		t.Errorf("received: %v, expected: %v", err, nil)
	}
	_, err = s.GetBaseData(&datakline.DataFromKline{
		Item:        &gctkline.Item{},
		Base:        d,
		RangeHolder: &gctkline.IntervalRangeHolder{},
	})
	if !errors.Is(err, nil) {
		t.Errorf("received: %v, expected: %v", err, nil)
	}
}

func TestSetSimultaneousProcessing(t *testing.T) {
	t.Parallel()
	s := Strategy{}
	is := s.UsingSimultaneousProcessing()
	if is {
		t.Error("expected false")
	}
	s.SetSimultaneousProcessing(true)
	is = s.UsingSimultaneousProcessing()
	if !is {
		t.Error("expected true")
	}
}

func TestCloseAllPositions(t *testing.T) {
	t.Parallel()
	s := &Strategy{}
	_, err := s.CloseAllPositions(nil, nil)
	if !errors.Is(err, gctcommon.ErrFunctionNotSupported) {
		t.Errorf("received '%v' expected '%v'", err, gctcommon.ErrFunctionNotSupported)
	}
}
