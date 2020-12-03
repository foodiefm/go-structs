package masterproduct

// PackagingMarkingModule is a module containing details on markings on the
// packaging of the trade item for example dates, environment.
type PackagingMarkingModule struct {
	// Details on markings on the packaging of the trade item.
	PackagingMarking *PackagingMarking `json:"packagingMarking,omitempty"`
}

// PackagingMarking is details on markings on the packaging of the trade item.
type PackagingMarking struct {
	// A marking that the trade item received recognition, endorsement,
	// certification by following guidelines by the label issuing agency.
	// Uses code list packagingMarkedLabelAccreditationCode.
	PackagingMarkedLabelAccreditationCode *[]string `json:"packagingMarkedLabelAccreditationCode,omitempty"`
}
