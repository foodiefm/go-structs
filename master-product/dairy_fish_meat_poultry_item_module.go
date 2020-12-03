package masterproduct

// DairyFishMeatPoultryItemModule describes content and processing related
// information specific to dairy, fish, meat and poultry products.
type DairyFishMeatPoultryItemModule struct {
	// Content and processing related information specific to dairy, fish,
	// meat and poultry products.
	DairyFishMeatPoultryInformation *DairyFishMeatPoultryInformation `json:":dairyFishMeatPoultryInformation,omitempty"`
}

// DairyFishMeatPoultryInformation describes content and processing related
// information specific to dairy, fish, meat and poultry products.
type DairyFishMeatPoultryInformation struct {
	FishReportingInformation *[]FishReportingInformation `json:":fishReportingInformation,omitempty"`
}

// FishReportingInformation describes fish reporting information for a trade
// item.
type FishReportingInformation struct {
	// The FAO 3 Alpha code of the species of fish for fish and seafood.
	// Uses code list speciesForFisheryStatisticsPurposesCode.
	SpeciesForFisheryStatisticsPurposesCode *string `json:":speciesForFisheryStatisticsPurposesCode"`
	// The scientific name associated with the
	// speciesforFisheryStatisticsPurposesCode.
	SpeciesForFisheryStatisticsPurposesName *string `json:":speciesForFisheryStatisticsPurposesName"`
	// Details related to a fish catch for a trade item.
	FishCatchInformation *[]FishCatchInformation `json:"fishCatchInformation"`
}

// FishCatchInformation describes details related to a fish catch for a trade item.
type FishCatchInformation struct {
	// The catch method for fish and seafood as specified by FAO, Fisheries
	// and Aquaculture Department of the Food and Agriculture Organization
	// of the United Nations. This aattribute will help the global retail
	// industry to fulfil the EU requirements for a common fisheries
	// policy.
	CatchMethodCode *[]string `json:"catchMethodCode"`
}
