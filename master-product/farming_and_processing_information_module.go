package masterproduct

// FarmingAndProcessingInformationModule contains information on any farming or
// processing performed on and agricultural trade item.
type FarmingAndProcessingInformationModule struct {
	// Details on the trade item regarding the extent of organic
	// production.
	TradeItemOrganicInformation *TradeItemOrganicInformation `json:"tradeItemOrganicInformation,omitempty"`
	// Information on farming and processing for a trade item.
	TradeItemFarmingAndProcessing *TradeItemFarmingAndProcessing `json:"tradeItemFarmingAndProcessing,omitempty"`
	// Attribute value pair information.
	AVPList *AVPList `json:"avpList"`
}

// TradeItemOrganicInformation details on the trade item regarding the extent
// of organic production.
type TradeItemOrganicInformation struct {
	// Indication of the place where the agricultural raw materials of
	// which the product is composed have been farmed. It applies only to
	// the trade item, not ingredient by ingredient. Uses code list
	// organicProductPlaceOfFarmingCode
	OrganicProductPlaceOfFarmingCode *string `json:"organicProductPlaceOfFarmingCode,omitempty"`
	// Any claim to indicate the organic status of a trade item or of one
	// or more of its components.
	OrganicClaims *[]OrganicClaim `json:"organicClaim,omitempty"`
}

// OrganicClaim contains any claim to indicate the organic status of a trade
// item or of one or more of its components.
type OrganicClaim struct {
	// A Governing body that creates and maintains standards related to
	// organic products. Uses code list organicClaimAgencyCode
	OrganicClaimAgencyCode *[]string `json:"organicClaimAgencyCode,omitempty"`
	// The percent of actual organic materials per weight of the trade
	// item. This is usually claimed on the product
	OrganicPercentClaim *float64 `json:"organicPercentClaim,omitempty"`
}

// TradeItemFarmingAndProcessing contains information on farming and processing
// for a trade item.
type TradeItemFarmingAndProcessing struct {
	// A statement of the presence or absence of genetically modified
	// protein or DNA. Uses code list geneticallyModifiedDeclarationCode.
	GeneticallyModifiedDeclarationCode *string `json:"geneticallyModifiedDeclarationCode.omitempty"`
	// Code value indicating the preservation technique used to preserve
	// the product from deterioration. Uses code list
	// preservationTechniqueCode.
	PreservationTechniqueCode *[]string `json:"preservationTechniqueCode.omitempty"`
}
