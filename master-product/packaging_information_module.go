package masterproduct

// PackagingInformationModule  contains packaging information for a trade item.
type PackagingInformationModule struct {
	// Details on packaging for a trade item.
	Packagings *[]Packaging `json:"packaging.omitempty"`
}

// Packaging details for a trade item.
type Packaging struct {
	// The process the packaging could undertake for recyclable &
	// sustainability programs.
	PackagingRecyclingProcessTypeCode *[]string `json:"packagingRecyclingProcessTypeCode,omitempty"`
	// The dominant means used to transport, store, handle or display the
	// trade item as defined by the data source. This packaging is not used
	// to describe any manufacturing process.Uses code list
	// packagingTypeCode.
	PackagingTypeCode *string `json:"packagingTypeCode,omitempty"`
	// Details on packaging material for a trade item's packaging.
	PackagingMaterials *[]PackagingMaterial `json:"packagingMaterial,omitempty"`
}

// PackagingMaterial is details on packaging material for a trade item's
// packaging.
type PackagingMaterial struct {
	// The materials used for the packaging of the trade item. Uses code
	// list packagingMaterialTypeCode.
	PackagingMaterialTypeCode string `json:"packagingMaterialTypeCode"`
	// Determines whether packaging material is recoverable. Recoverable
	// materials are those which are capable of beingreused or returned to
	// use in the form of raw materials.
	IsPackagingMaterialRecoverable *bool `json:"isPackagingMaterialRecoverable,omitempty"`
}
