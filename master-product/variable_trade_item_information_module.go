package masterproduct

// VariableTradeItemInformationModule is a module with information specific to
// variable weight or dimension trade items.
type VariableTradeItemInformationModule struct {
	// Information specific to variable weight or dimension trade items.
	VariableTradeItemInformation *VariableTradeItemInformation `json:"variableTradeItemInformation,omitempty"`
}

// VariableTradeItemInformation is information specific to variable weight or
// dimension trade items.
type VariableTradeItemInformation struct {
	// Indicates that an article is not a fixed quantity, but that the
	// quantity is variable. Can be weight, length, volume. trade item is
	// used or traded in continuous rather than discrete quantities.
	IsTradeItemAVariableUnit *bool `json:"isTradeItemAVariableUnit,omitempty"`
	// Indicator to show whether product is loose or pre-packed. Uses code
	// list variableTradeItemTypeCode.
	VariableTradeItemTypeCode *string `json:"variableTradeItemTypeCode,omitempty"`
	// Indication of the percentage value that the actual weight of the
	// trade item may differ from the average or estimated weight given.
	// For example, Roast beef off the bone 3.5 kg, Gross weight 3500
	// Grams, Range = 14 %. This means that this item may be produced with
	// weight values ranging from 3.010 kg to 3.990 kg.
	VariableWeightAllowableDeviationPercentage *float64 `json:"variableWeightAllowableDeviationPercentage,omitempty"`
}
