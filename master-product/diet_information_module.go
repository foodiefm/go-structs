package masterproduct

// DietInformationModule is a module contain a product dietary suitability.
type DietInformationModule struct {
	// The diet the product is suitable for.
	DietInformation *DietInformation `json:"dietInformation,omitempty"`
}

// DietInformation is the diet the product is suitable for.
type DietInformation struct {
	// Expresses in text the dietary description of a product which are
	// normally held on the label or accompanying the product. This
	// information may or may not be labeled on the pack. Instructions may
	// refer to a suggested lifestyle or dietary preference.
	DietTypeDescriptions *[]MLString            `json:"dietTypeDescription"`
	DietTypeInformations *[]DietTypeInformation `json:"dietTypeInformation"`
}

// DietTypeInformation Expresses in text the suggested dietary suitability of a
// product which are normally held on the label or accompanying the product.
// This information may or may not be labeled on the pack.
type DietTypeInformation struct {
	DietTypeCode    string  `json:"dietTypeCode"`
	DietTypeSubcode *string `json:"dietTypeSubcode.omitempty"`
}
