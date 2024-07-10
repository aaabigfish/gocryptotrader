package ticker

import (
	"errors"
	"fmt"
	"github.com/aaabigfish/gocryptotrader/common/key"
	"github.com/aaabigfish/gocryptotrader/currency"
	"github.com/aaabigfish/gocryptotrader/exchanges/asset"
	"github.com/gofrs/uuid"
	"strings"
)

var (
	// ErrNoTickerFound is when a ticker is not found
	ErrNoTickerFound = errors.New("no ticker found")
	// ErrBidEqualsAsk error for locked markets
	ErrBidEqualsAsk = errors.New("bid equals ask this is a crossed or locked market")
	// ErrExchangeNameIsEmpty is an error for when an exchange name is empty
	ErrExchangeNameIsEmpty = errors.New("exchange name is empty")

	errExchangeNotFound = errors.New("exchange not found")
)

func init() {
	service = new(Service)
	service.Tickers = make(map[key.ExchangePairAsset]*Ticker)
	service.Exchange = make(map[string]uuid.UUID)
}

// GetTicker checks and returns a requested ticker if it exists
func GetTicker(exchange string, p currency.Pair, a asset.Item) (*Price, error) {
	if exchange == "" {
		return nil, ErrExchangeNameIsEmpty
	}
	if p.IsEmpty() {
		return nil, currency.ErrCurrencyPairEmpty
	}
	if !a.IsValid() {
		return nil, fmt.Errorf("%w %v", asset.ErrNotSupported, a)
	}
	exchange = strings.ToLower(exchange)
	service.mu.Lock()
	defer service.mu.Unlock()
	tick, ok := service.Tickers[key.ExchangePairAsset{
		Exchange: exchange,
		Base:     p.Base.Item,
		Quote:    p.Quote.Item,
		Asset:    a,
	}]
	if !ok {
		return nil, fmt.Errorf("%w %s %s %s",
			ErrNoTickerFound, exchange, p, a)
	}

	cpy := tick.Price // Don't let external functions have access to underlying
	return &cpy, nil
}

func GetExchangeTickers(exchange string) ([]*Price, error) {
	return service.getExchangeTickers(exchange)
}

func (s *Service) getExchangeTickers(exchange string) ([]*Price, error) {
	if exchange == "" {
		return nil, ErrExchangeNameIsEmpty
	}
	exchange = strings.ToLower(exchange)
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.Exchange[exchange]
	if !ok {
		return nil, fmt.Errorf("%w %v", errExchangeNotFound, exchange)
	}
	tickers := make([]*Price, 0, len(s.Tickers))
	for k, v := range s.Tickers {
		if k.Exchange != exchange {
			continue
		}
		cpy := v.Price // Don't let external functions have access to underlying
		tickers = append(tickers, &cpy)
	}
	return tickers, nil
}

// ProcessTicker processes incoming tickers, creating or updating the Tickers
// list
func ProcessTicker(p *Price) error {
	return nil
}
