package masterproduct

// DGProductAttributeModule is a module containing freely defined product
// attributes.
type DGProductAttributeModule struct {
	// Product attribute groups used to collect attributes into meaningful
	// sets.
	ProductAttributeGroups *[]ProductAttributeGroup `json:"productAttributeGroup,omitempty"`
}

// ProductAttributeGroup describes product attribute group used to collect
// attributes into meaningful sets.
type ProductAttributeGroup struct {
	// Product-unique data provider assigned identifer.
	ProductAttributeGroupExtID string `json:"productAttributeGroupExtId"`
	// Value indicating the group order.
	ProductAttributeGroupSequence *string `json:"productAttributeGroupSequence,omitempty"`
	// Attribute group name used for presentation.
	ProductAttributeGroupNames *[]MLString `json:"productAttributeGroupName,omitempty"`
	// Product attributes.
	ProductAttributes *[]ProductAttribute `json:"productAttribute,omitempty"`
}

// ProductAttribute describes attribute declaration. Note that while neither
// productAttributeValueString, productAttributeValueNumeric, nor
// productAttributeValueBoolean is required, exactly one of these must be
// provided.
type ProductAttribute struct {
	// Group-unique data provider assigned identifier.
	ProductAttributeExtID string `json:"productAttributeExtId"`
	// Value indicating the attribute order.
	ProductAttributeSequence *string `json:"productAttributeSequence,omitempty"`
	// Code specifying the attribute type. Uses code list
	// productAttributeTypeCode.
	ProductAttributeTypeCode string `json:"productAttributeTypeCode"`
	// An indicator whether or not the attribute is and can be used as a
	// facet attribute.
	IsFacetAttribute             *bool       `json:"isFacetAttribute,omitempty"`
	ProductAttributeNames        *[]MLString `json:"productAttributeName,omitempty"`
	ProductAttributeValueStrings *[]MLString `json:"productAttributeValueString,omitempty"`
	ProductAttributeValueNumeric *float64    `json:"productAttributeValueNumeric,omitempty"`
	ProductAttributeValueBoolean bool        `json:"productAttributeValueBoolean,omitempty"`
}
