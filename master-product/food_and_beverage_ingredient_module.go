package masterproduct

// FoodAndBeverageIngredientModule contains information on the constituent
// ingredient make up of the product.
type FoodAndBeverageIngredientModule struct {
	// Information on the constituent ingredient make up of the product
	// specified as one string.
	IngredientStatements *[]MLString `json:"ingredientStatement.omitempty"`
	// The fruit juice content of the trade item expressed as a percentage.
	JuiceContentPercent *float64 `json:"juiceContentPercent.omitempty"`
	// Information on presence or absence of additives or genetic
	// modifications contained in the trade item.
	AdditiveInformations *[]AdditiveInformation `json:"additiveInformation.omitempty"`
	// Information on the constituent ingredient make up of the product
	// split out per ingredient.
	FoodAndBeverageIngredients *[]FoodAndBeverageIngredient `json:"foodAndBeverageIngredient,omitempty"`
	// Free text field for any additional ingredient information.
	XAdditionalIngredientStatements *[]MLString `json:"x_additionalIngredientStatement,omitempty"`
	// Denotes that the product in question is either a food item or a
	// beverage.
	XIsFoodOrBeverage bool `json:"x_isFoodOrBeverage"`
}

// AdditiveInformation contains information on presence or absence of additives
// or genetic modifications contained in the trade item.
type AdditiveInformation struct {
	// The name of any additive or genetic modification contained or not
	// contained in the trade item.
	AdditiveName string `json:"additiveName"`
	// Code indicating the level of presence of the additive. Uses code
	// list levelOfContainmentCode.
	LevelOfContainmentCode string `json:"levelOfContainmentCode"`
}

// FoodAndBeverageIngredient contains information on the constituent ingredient
// make up of the product split out per ingredient.
type FoodAndBeverageIngredient struct {
	// Value indicating the ingredient order.
	IngredientSequence *string `json:"ingredientSequence.omitempty"`
	// Indication of the percentage of the ingredient contained in the
	// product.
	IngredientContentPercentage *float64 `json:"ingredientContentPercentage,omitempty"`
	// Text field indicating one ingredient or ingredient group (according
	// to regulations of the target market). Ingredients include any
	// additives (colorings, preservatives, e-numbers, etc) that are
	// encompassed.
	IngredientNames []MLStringEmphasised `json:"ingredientName"`
	// Denotes that the ingredient should have it's text emphasised.
	IsIngredientEmphasised *bool `json:"isIngredientEmphasised.omitempty"`
	// Details on any methods and techniques used by a manufacturer or
	// supplier to the trade item, ingredients or raw materials.
	IngredientFarmingProcessing *IngredientFarmingProcessing `json:"ingredientFarmingProcessing.omitempty"`
	// Information on the organic nature of ingredient.
	IngredientOrganicInformation *IngredientOrganicInformation `json:"ingredientOrganicInformation,omitempty"`
	// Information on the activity (e.g. bottling) taken place for an
	// ingredient as well as the associated geographic area.
	IngredientPlaceOfActivities *[]IngredientPlaceOfActivity `json:"ingredientPlaceOfActivity,omitempty"`
}

// IngredientFarmingProcessing details on any methods and techniques used by a
// manufacturer or supplier to the trade item, ingredients or raw materials.
type IngredientFarmingProcessing struct {
	// A statement of the presence or absence of genetically modified
	// protein or DNA. Uses code list geneticallyModifiedDeclarationCode.
	GeneticallyModifiedDeclarationCode *string `json:"geneticallyModifiedDeclarationCode.omitempty"`
	// Code value indicating the preservation technique used to preserve
	// the product from deterioration. Uses code list
	// preservationTechniqueCode.
	PreservationTechniqueCode *[]string `json:"preservationTechniqueCode,omitempty"`
}

// IngredientOrganicInformation contains information on the organic nature of
// ingredient.
type IngredientOrganicInformation struct {
	// Indication of the place where the agricultural raw materials of
	// which the product is composed have been farmed. It applies only to
	// the trade item, not ingredient by ingredient. Uses code list
	// organicProductPlaceOfFarmingCode.
	OrganicProductPlaceOfFarmingCode *string `json:"organicProductPlaceOfFarmingCode.omitempty"`
	// Any claim to indicate the organic status of a trade item or of one
	// or more of its components.
	OrganicClaim *[]IngredientOrganicClaim `json:"organicClaim,omitempty"`
}

// IngredientOrganicClaim Any claim to indicate the organic status of a trade
// item or of one or more of its components.
type IngredientOrganicClaim struct {
	// A Governing body that creates and maintains standards related to
	// organic products. Uses code list organicClaimAgencyCode.
	OrganicClaimAgencyCode *[]string `json:"organicClaimAgencyCode,omitempty"`
	// The percent of actual organic materials per weight of the trade
	// item. This is usually claimed on the product.
	OrganicPercentClaim float64 `json:"organicPercentClaim,omitempty"`
}

// IngredientPlaceOfActivity contains information on the activity (e.g.
// bottling) taken place for an ingredient as well as the associated geographic
// area.
type IngredientPlaceOfActivity struct {
	// A description of the country the item may have originated from or
	// has been processed.
	CountryOfOriginStatements *[]MLString `json:"countryOfOriginStatement,omitempty"`
	// The place a trade item originates from. This is to be specifically
	// used to enable things such as cities, mountain ranges, regions that
	// do not comply with ISO standards.
	ProvenanceStatements []MLString `json:"provenanceStatement.omitempty"`
	// The country the item may have originated from or has been processed.
	CountryOfOrigins *[]CountryOfOrigin `json:"countryOfOrigin,omitempty"`
	// Details on the activity (e.g. bottling) taken place for a trade item
	// as well as the associated geographic area.
	ProductActivityDetails *[]ProductActivityDetail `json:"productActivityDetails.omitempty"`
}
