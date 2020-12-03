package masterproduct

// TradeItemDescriptionModule a module carrying general descriptions of the
// trade item including brand, form, variant.
type TradeItemDescriptionModule struct {
	// Description Information for the trade item.
	TradeItemDescriptionInformation *TradeItemDescriptionInformation `json:"tradeItemDescriptionInformation,omitempty"`
}

// TradeItemDescriptionInformation is description information for the trade
// item.
type TradeItemDescriptionInformation struct {
	// Additional variants necessary to communicate to the industry to help
	// define the product.
	AdditionalTradeItemDescriptions *[]MLString `json:"additionalTradeItemDescription,omitempty"`
	// A free form short length description of the trade item that can be
	// used to identify the trade item at point of sale.
	DescriptionShorts *[]MLString `json:"descriptionShort,omitempty"`
	// Describes use of the product or service by the consumer. Should help
	// clarify the product classification associated with the GTIN.
	FunctionalNames *[]MLString `json:"functionalName,omitempty"`
	// An understandable and useable description of a trade item using
	// brand and other descriptors. This attribute is filled with as little
	// abbreviation as possible while keeping to a reasonable length. This
	// should be a meaningful description of the trade item with full
	// spelling to facilitate message processing. Retailers can use this
	// description as the base to fully understand the brand, flavour,
	// scent etc. of the specific GTIN in order to accurately create a
	// product description as needed for their internal systems.
	TradeItemDescriptions *[]MLString `json:"tradeItemDescription,omitempty"`
	// Free text field used to identify the variant of the product.
	// Variants are the distinguishing characteristics that differentiate
	// products with the same brand and size including such things as the
	// particular flavor, fragrance, taste.
	VariantDescriptions *[]MLString `json:"variantDescription,omitempty"`
	// Information on brands and sub-brands for a trade item.
	BrandNameInformation *BrandNameInformation `json:"brandNameInformation,omitempty"`
}

// BrandNameInformation contains information on brands and sub-brands for a trade item.
type BrandNameInformation struct {
	// The recognisable name used by a brand owner to uniquely identify a
	// line of trade item or services. This is recognizable by the
	// consumer.
	BrandName *string `json:"brandName,omitempty"`
	// The recognisable name used by a brand owner to uniquely identify a
	// line of trade item or services expressed in a different language
	// than the primary brand name (brandName).
	LanguageSpecificBrandNames *[]MLString `json:"languageSpecificBrandName,omitempty"`
	// A second level of brand expressed in a different language than the
	// primary sub-brand name (subBrand).
	LanguageSpecificSubbrandNames *[]MLString `json:"languageSpecificSubbrandName,omitempty"`
	// Second level of brand. Can be a trademark. It is the primary
	// differentiating factor that a brand owner wants to communicate to
	// the consumer or buyer.
	SubBrand *string `json:"subBrand,omitempty"`
}
