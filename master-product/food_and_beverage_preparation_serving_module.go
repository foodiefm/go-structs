package masterproduct

// FoodAndBeveragePreparationServingModule is information on way the product
// can be prepared or served.
type FoodAndBeveragePreparationServingModule struct {
	// Preparation and serving information for a food and beverage item.
	PreparationServings *[]PreparationServing `json:"preparationServing,omitempty"`
}

// PreparationServing contains preparation and serving information for a food
// and beverage item.
type PreparationServing struct {
	// An indication of the ease of preparation for semi-prepared products.
	// The convenience level indicates the level of preparation in
	// percentage required to prepare and helps the consumer to assess how
	// long it will take to prepare the meal.
	ConvenienceLevelPercent float64 `json:"convenienceLevelPercent,omitempty"`
	// Textual instruction on how to prepare the product before serving.
	PreparationInstructions *[]MLString `json:"preparationInstructions,omitempty"`
	// A code specifying the technique used to make the product ready for
	// consumption. Uses code list preparationTypeCode.
	PreparationTypeCode *string `json:"preparationTypeCode,omitempty"`
	// Free text field for serving suggestion.
	ServingSuggestions *[]MLString `json:"servingSuggestion,omitempty"`
	// Information on the yield of a product.
	ProductYieldInformations *[]ProductYieldInformation `json:"productYieldInformation,omitempty"`
}

// ProductYieldInformation is a information on the yield of a product.
type ProductYieldInformation struct {
	// Measurement.
	ProductYield *Measurement `json:"productYield.omitempty"`
	// Code indicating the type of yield measurement. Uses code list
	// productYieldTypeCode.
	ProductYieldTypeCode *string `json:"productYieldTypeCode,omitempty"`
}
