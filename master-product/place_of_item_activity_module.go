package masterproduct

// PlaceOfItemActivityModule contains information on the activity (e.g.
// bottling) taken place for a trade item as well as the associated geographic
// area.
type PlaceOfItemActivityModule struct {
	// Information on the activity (e.g. bottling) taken place for a trade
	// item as well as the associated geographic area.
	PlaceOfProductActivity *PlaceOfProductActivity `json:"placeOfProductActivity,omitempty"`
}

// PlaceOfProductActivity contains information on the activity (e.g. bottling)
// taken place for a trade item as well as the associated geographic area.
type PlaceOfProductActivity struct {
	// A description of the country the item may have originated from or
	// has been processed.
	CountryOfOriginStatements *[]MLString `json:"countryOfOriginStatement,omitempty"`
	// The place a trade item originates from. This is to be specifically
	// used to enable things such as cities, mountain ranges, regions that
	// do not comply with ISO standards.
	ProvenanceStatements *[]MLString `json:"provenanceStatement,omitempty"`
	// The country the item may have originated from or has been processed.
	CountriesOfOrigin *[]CountryOfOrigin `json:"countryOfOrigin,omitempty"`
	// Details on the activity (e.g. bottling) taken place for a trade item
	// as well as the associated geographic area.
	ProductActivityDetails *[]ProductActivityDetail `json:"productActivityDetails,omitempty"`
}
