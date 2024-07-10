package account

import (
	"errors"
	"fmt"
	"github.com/aaabigfish/gocryptotrader/currency"
	"github.com/aaabigfish/gocryptotrader/exchanges/asset"
	"strings"
)

func init() {
	service.exchangeAccounts = make(map[string]*Accounts)
}

var (
	errHoldingsIsNil                = errors.New("holdings cannot be nil")
	errExchangeNameUnset            = errors.New("exchange name unset")
	errExchangeHoldingsNotFound     = errors.New("exchange holdings not found")
	errAssetHoldingsNotFound        = errors.New("asset holdings not found")
	errExchangeAccountsNotFound     = errors.New("exchange accounts not found")
	errNoExchangeSubAccountBalances = errors.New("no exchange sub account balances")
	errBalanceIsNil                 = errors.New("balance is nil")
	errNoCredentialBalances         = errors.New("no balances associated with credentials")
	errCredentialsAreNil            = errors.New("credentials are nil")
)

// CollectBalances converts a map of sub-account balances into a slice
func CollectBalances(accountBalances map[string][]Balance, assetType asset.Item) (accounts []SubAccount, err error) {
	if accountBalances == nil {
		return nil, errAccountBalancesIsNil
	}

	if !assetType.IsValid() {
		return nil, fmt.Errorf("%s, %w", assetType, asset.ErrNotSupported)
	}

	accounts = make([]SubAccount, 0, len(accountBalances))
	for accountID, balances := range accountBalances {
		accounts = append(accounts, SubAccount{
			ID:         accountID,
			AssetType:  assetType,
			Currencies: balances,
		})
	}
	return
}

// Process processes new account holdings updates
func Process(h *Holdings, c *Credentials) error {
	return nil
}

// GetHoldings returns full holdings for an exchange.
// NOTE: Due to credentials these amounts could be N*APIKEY actual holdings.
// TODO: Add jurisdiction and differentiation between APIKEY holdings.
func GetHoldings(exch string, creds *Credentials, assetType asset.Item) (Holdings, error) {
	if exch == "" {
		return Holdings{}, errExchangeNameUnset
	}

	if creds.IsEmpty() {
		return Holdings{}, fmt.Errorf("%s %s %w", exch, assetType, errCredentialsAreNil)
	}

	if !assetType.IsValid() {
		return Holdings{}, fmt.Errorf("%s %s %w", exch, assetType, asset.ErrNotSupported)
	}

	exch = strings.ToLower(exch)

	service.mu.Lock()
	defer service.mu.Unlock()
	accounts, ok := service.exchangeAccounts[exch]
	if !ok {
		return Holdings{}, fmt.Errorf("%s %s %w", exch, assetType, errExchangeHoldingsNotFound)
	}

	subAccountHoldings, ok := accounts.SubAccounts[*creds]
	if !ok {
		return Holdings{}, fmt.Errorf("%s %s %s %w",
			exch,
			creds,
			assetType,
			errNoCredentialBalances)
	}

	var currencyBalances = make([]Balance, 0, len(subAccountHoldings))
	cpy := *creds
	for mapKey, assetHoldings := range subAccountHoldings {
		if mapKey.Asset != assetType {
			continue
		}
		assetHoldings.m.Lock()
		currencyBalances = append(currencyBalances, Balance{
			Currency:               currency.Code{Item: mapKey.Currency, UpperCase: true},
			Total:                  assetHoldings.total,
			Hold:                   assetHoldings.hold,
			Free:                   assetHoldings.free,
			AvailableWithoutBorrow: assetHoldings.availableWithoutBorrow,
			Borrowed:               assetHoldings.borrowed,
		})
		assetHoldings.m.Unlock()
		if cpy.SubAccount == "" && mapKey.SubAccount != "" {
			// TODO: fix this backwards population
			// the subAccount here may not be associated with the balance across all subAccountHoldings
			cpy.SubAccount = mapKey.SubAccount
		}
	}
	if len(currencyBalances) == 0 {
		return Holdings{}, fmt.Errorf("%s %s %w",
			exch,
			assetType,
			errAssetHoldingsNotFound)
	}
	return Holdings{Exchange: exch, Accounts: []SubAccount{{
		Credentials: Protected{creds: cpy},
		ID:          cpy.SubAccount,
		AssetType:   assetType,
		Currencies:  currencyBalances,
	}}}, nil
}
