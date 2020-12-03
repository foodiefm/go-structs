package masterproduct

// AllergenInformationModule is a module containing information on allergens
// for a trade item.
type AllergenInformationModule struct {
	// Information on substances that might cause allergic reactions and
	// substances subject to intolerance when consumed. The allergy
	// information refers to specified regulations that apply to the target
	// market to which the item information is published.
	AllergenRelatedInformations *[]AllergenRelatedInformation `json:"allergenRelatedInformation,omitempty"`
}

// AllergenRelatedInformation contains information on substances that might
// cause allergic reactions and substances subject to intolerance when
// consumed. The allergy information refers to specified regulations that apply
// to the target market to which the item information is published.
type AllergenRelatedInformation struct {
	// Agency that controls the allergen definition.
	AllergenSpecificationAgency *string `json:"allergenSpecificationAgency.omitempty"`
	// Free text field containing the name and version of the regulation or
	// standard that contains the definition of the allergen.
	AllergenSpecificationName *string `json:"allergenSpecificationName.omitempty"`
	// Textual description of the presence or absence of allergens as
	// governed by local rules and regulations, specified as one string.
	AllergenStatements *[]MLString `json:"allergenStatement.omitempty"`
	// Description of the presence or absence of allergens as governed by
	// local rules and regulations, specified per allergen.
	Allergens *[]Allergen `json:"allergen,omitempty"`
}

// Allergen is a description of the presence or absence of allergens as
// governed by local rules and regulations, specified per allergen.
type Allergen struct {
	// Code indicating the type of allergen. Uses code list
	// allergenTypeCode.
	AllergenTypeCode string `json:"allergenTypeCode"`
	// Code indicating the level of presence of the allergen.
	LevelOfContainmentCode string `json:"levelOfContainmentCode"`
}
