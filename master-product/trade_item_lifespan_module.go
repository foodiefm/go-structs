package masterproduct

// TradeItemLifespanModule is a module containing information on the amount of
// time the item can or should be used, sold, etc.
type TradeItemLifespanModule struct {
	// Information on the amount of time the item can or should be used,
	// sold, etc.
	TradeItemLifespan *TradeItemLifespan `json:"tradeItemLifespan,omitempty"`
}

// TradeItemLifespan contains information on the amount of time the item can or
// should be used, sold, etc.
type TradeItemLifespan struct {
	// The period of day, guaranteed by the manufacturer, before the
	// expiration date of the product, based on the production.
	MinimumTradeItemLifespanFromTimeOfProduction *int `json:"minimumTradeItemLifespanFromTimeOfProduction,omitempty"`
	// The number of days the trade item that had been opened can remain on
	// the shelf and must then be removed.
	OpenedTradeItemLifespan *int `json:"openedTradeItemLifespan,omitempty"`
}
