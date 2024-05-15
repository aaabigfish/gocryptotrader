package strategies

import (
	"errors"

	"github.com/aaabigfish/gocryptotrader/backtester/data"
	"github.com/aaabigfish/gocryptotrader/backtester/eventhandlers/portfolio"
	"github.com/aaabigfish/gocryptotrader/backtester/eventhandlers/portfolio/holdings"
	"github.com/aaabigfish/gocryptotrader/backtester/eventtypes/signal"
	"github.com/aaabigfish/gocryptotrader/backtester/funding"
)

// ErrStrategyAlreadyExists returned when a strategy matches the same name
var ErrStrategyAlreadyExists = errors.New("strategy already exists")

// StrategyHolder holds strategies
type StrategyHolder []Handler

// Handler defines all functions required to run strategies against data events
type Handler interface {
	Name() string
	Description() string
	OnSignal(data.Handler, funding.IFundingTransferer, portfolio.Handler) (signal.Event, error)
	OnSimultaneousSignals([]data.Handler, funding.IFundingTransferer, portfolio.Handler) ([]signal.Event, error)
	UsingSimultaneousProcessing() bool
	SupportsSimultaneousProcessing() bool
	SetSimultaneousProcessing(bool)
	SetCustomSettings(map[string]interface{}) error
	SetDefaults()
	CloseAllPositions([]holdings.Holding, []data.Event) ([]signal.Event, error)
}
