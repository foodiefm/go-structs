package masterproduct

// MarketingInformationModule contains information of a trade item meant to
// convey features and benefits and targeted customer.
type MarketingInformationModule struct {
	// Information on a trade item meant to convey features and benefits.
	MarketingInformation *MarketingInformation `json:"marketingInformation,omitempty"`
}

// MarketingInformation contains information of a trade item meant to convey
// features and benefits.
type MarketingInformation struct {
	// Marketing message associated to the Trade item.
	TradeItemMarketingMessages *[]MLString `json:"tradeItemMarketingMessage,omitempty"`
	// Words or phrases that enables web search engines to find trade items
	// on the internet for example Shampoo, Lather, Baby.
	TradeItemKeyWords *[]MLString `json:"tradeItemKeyWords,omitempty"`
	// An indicator whether or not the Trade Item is excluded and hidden
	// from promotions. When not defined, assumed to be false.
	XHideTradeItemFromPromotions *bool `json:"x_hideTradeItemFromPromotions,omitempty"`
}
