package masterproduct

// TradeItemMeasurementsModule is a module containing measurement information
// for the trade item.
type TradeItemMeasurementsModule struct {
	// Measurement information for the trade item.
	TradeItemMeasurements *TradeItemMeasurements `json:"tradeItemMeasurements,omitempty"`
}

// TradeItemMeasurements is measurement information for the trade item.
type TradeItemMeasurements struct {
	Depth  *Measurement `json:"depth,omitempty"`
	Height *Measurement `json:"height,omitempty"`
	Width  *Measurement `json:"width,omitempty"`
	// The amount of the trade item contained by a package, usually as
	// claimed on the label. For example, Water 750ml - net content = "750
	// MLT" ; 20 count pack of diapers, net content = "20 ea.". In case of
	// multi-pack, indicates the net content of the total trade item. For
	// fixed value trade items use the value claimed on the package, to
	// avoid variable fill rate issue that arises with some trade item
	// which are sold by volume or weight, and whose actual content may
	// vary slightly from batch to batch. In case of variable quantity
	// trade items, indicates the average quantity.
	NetContent *[]Measurement `json:"netContent,omitempty"`
	// Information on the weight of a trade item.
	TradeItemWeight *TradeItemWeight `json:"tradeItemWeight,omitempty"`
}

// TradeItemWeight is information on the weight of a trade item.
type TradeItemWeight struct {
	DrainedWeight *Measurement `json:"drainedWeight,omitempty"`
	GrossWeight   *Measurement `json:"grossWeight,omitempty"`
	NetWeight     *Measurement `json:"netWeight,omitempty"`
}
