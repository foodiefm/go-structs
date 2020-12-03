package masterproduct

// NutritionalInformationModule contains information about content of
// nutrients. Multiple sets of nutrient information can be specified with
// varying state, serving size and daily value intake base.
type NutritionalInformationModule struct {
	// Free text field for any additional nutritional claims.
	NutritionalClaims *[]MLString `json:"nutritionalClaim.omitempty"`
	// Details on a nutritional claim for a trade item permitted by known
	// regulations for a target market.
	NutritionalClaimDetails *[]NutritionalClaimDetail `json:"nutritionalClaimDetail,omitempty"`
	// Nutrient information for a trade item.
	NutrientHeaders *[]NutrientHeader `json:"nutrientHeader,omitempty"`
}

// NutritionalClaimDetail contains Details on a nutritional claim for a trade
// item permitted by known regulations for a target market.
type NutritionalClaimDetail struct {
	// A code depicting the degree to which a trade item contains a
	// specific nutrient or ingredient in relation to a health claim. Uses
	// code list nutritionalClaimTypeCode.
	NutritionalClaimTypeCode string `json:"nutritionalClaimTypeCode"`
	// The type of nutrient, ingredient, vitamins and minerals that the
	// nutritional claim is in reference to for example fat, copper, milk.
	// Uses code list nutritionalClaimNutrientElementCode.
	NutritionalClaimNutrientElementCode string `json:"nutritionalClaimNutrientElementCode"`
}

// NutrientHeader contains nutrient  information for a trade item.
type NutrientHeader struct {
	// Code specifying the preparation state or type the nutrient
	// information applies to, for example, unprepared, boiled, fried. Uses
	// code list preparationStateCode.
	PreparationStateCode string `json:"preparationStateCode"`
	// Free text field specifying the daily value intake base for on which
	// the daily value intake per nutrient has been based.
	DailyValueIntakeReferences *[]MLString `json:"dailyValueIntakeReference,omitempty"`
	// Unit of measure code. Uses code list measurementUnitCode.
	NutrientBasisQuantity *Measurement `json:"nutrientBasisQuantity,omitempty"`
	// Measurement value specifying the serving size in which the
	// information per nutrient has been stated.
	ServingSizes *[]Measurement `json:"servingSize,omitempty"`
	// A free text field specifying the serving size for which the nutrient
	// information has been stated.
	ServingSizeDescriptions *[]MLString `json:"servingSizeDescription,omitempty"`
	// Nutrient detail for a trade item.
	NutrientDetails *[]NutrientDetail `json:"nutrientDetail,omitempty"`
}

// NutrientDetail describes nutrient detail for a trade item.
type NutrientDetail struct {
	// Nutrient type code. Uses code list nutrientTypeCode.
	NutrientTypeCode string `json:"nutrientTypeCode"`
	// The percentage of the recommended daily intake of a nutrient as
	// recommended by authorities of the target market. Is expressed
	// relative to the serving size and base daily value intake.
	DailyValueIntakePercent *float64 `json:"dailyValueIntakePercent,omitempty"`
	// Code indicating whether the specified nutrient content is exact or
	// approximate. One should follow local regulatory guidelines when
	// selecting a precision. Uses code list measurementPrecisionCode.
	MeasurementPrecisionCode *string `json:"measurementPrecisionCode,omitempty"`
	// Measurement value indicating the amount of nutrient contained in the
	// product. Is expressed relative to the serving size.
	QuantitiesContained *[]Measurement `json:"quantityContained.omitempty"`
}
