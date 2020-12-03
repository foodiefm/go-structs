package masterproduct

// TradeItemTemperatureInformationModule is information on temperature
// considerations for trade item.
type TradeItemTemperatureInformationModule struct {
	// The condition of the product sold to the end consumer. Uses code
	// list tradeItemTemperatureConditionTypeCode.
	TradeItemTemperatureConditionTypeCode *string `json:"tradeItemTemperatureConditionTypeCode,omitempty"`
	// Details on permissible temperatures of a trade item during various
	// points of the supply chain.
	TradeItemTemperatureInformations *[]TradeItemTemperatureInformation `json:"tradeItemTemperatureInformation,omitempty"`
}

// TradeItemTemperatureInformation describes details on permissible
// temperatures of a trade item during various points of the supply chain.
type TradeItemTemperatureInformation struct {
	MaximumTemperature          *TemperatureMeasurement `json:"maximumTemperature,omitempty"`
	MaximumToleranceTemperature *TemperatureMeasurement `json:"maximumToleranceTemperature,omitempty"`
	MinimumTemperature          *TemperatureMeasurement `json:"minimumTemperature,omitempty"`
	MinumumToleranceTemperature *TemperatureMeasurement `json:"minumumToleranceTemperature,omitempty"`
	// Code qualifying the type of a temperature requirement for example
	// Storage. Uses code list temperatureQualifierCode.
	TemperatureQualifierCode *string `json:"temperatureQualifierCode,omitempty"`
}
