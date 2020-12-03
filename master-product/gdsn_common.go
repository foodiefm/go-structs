package masterproduct

// MLString is a string with language.
type MLString struct {
	Value        string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// MLStringEmphasised is an emphasised string with language.
type MLStringEmphasised struct {
	MLString
	Emphases *[]XEmphasis `json:"x_emphasis,omitempty"`
}

// XEmphasis is a substring emphasis. Emphases may overlap.
type XEmphasis struct {
	// Emphasis starting index in characters from the beginning
	// of the string. Index starts at zero.
	StartAt int `json:"startAt"`
	// Emphasis length in characters.
	Length int `json:"length"`
}

// ExternalCodeValue represents an external code value.-
type ExternalCodeValue struct {
	// The name of the agency that manages a code list.
	ExternalAgencyName *string `json:"externalAgencyName.omitempty"`
	// The name of the code list maintained by an external agency.
	ExternalCodeListName         *string            `json:"externalCodeListName,omitempty"`
	EnumerationValueInformations []EnumerationValue `json:"enumerationValueInformation"`
}

// EnumerationValue is an enumeration value.
type EnumerationValue struct {
	// Code List Value maintained by an external code list agency.
	EnumerationValue string `json:"enumerationValue"`
}

// AVPList is attribute value pair information.
type AVPList struct {
	// Attribute values
	StringAVPs *[]StringAVP `json:"stringAVP,omitempty"`
}

// StringAVP presents Attribute values.
type StringAVP []struct {
	Value string `json:"$"`
	//Normalised attribute name
	AttributeName string `json:"@attributeName"`
}

// Measurement represents measurement.
type Measurement struct {
	Measurement         float64 `json:"$"`
	MeasurementUnitCode string  `json:"@measurementUnitCode"`
}

// TemperatureMeasurement represents temperature measurement.
type TemperatureMeasurement struct {
	Temperature                    int    `json:"$"`
	TemperatureMeasurementUnitCode string `json:"@temperatureMeasurementUnitCode"`
}

// CountryOfOrigin is the country the item may have originated from or has been
// processed.
type CountryOfOrigin struct {
	// Code specifying a country. Use code list countryCode.
	CountryCode string `json:"countryCode"`
}

// ProductActivityDetail contains details on the activity (e.g. bottling) taken
// place for a trade item as well as the associated geographic area.
type ProductActivityDetail struct {
	// A code depicting the type of activity being performed on a trade
	// item. Uses code list productActivityTypeCode.
	ProductActivityTypeCode string `json:"productActivityTypeCode"`
	// Country where activity happens.
	CountryOfActivities *[]CountryOfActivity `json:"countryOfActivity,omitempty"`
	// An external code value that depicts a specific zone or region for
	// example a FAO Catch Zone.
	ProductActivityRegionZoneCodeReferences *[]ProductActivityRegionZoneCodeReference `json:"productActivityRegionZoneCodeReference,omitempty"`
	// Free text field used to describe the activity region.
	XStatements *[]MLString `json:"x_statement,omitempty"`
}

// CountryOfActivity contains country where activity happens.
type CountryOfActivity struct {
	// Code specifying a country. Use code list countryCode
	CountryCode string `json:"countryCode"`
}

// ProductActivityRegionZoneCodeReference is an external code value that
// depicts a specific zone or region for example a FAO Catch Zone.
type ProductActivityRegionZoneCodeReference struct {
	// The name of the agency that manages a code list.
	ExternalAgencyName *string `json:"externalAgencyName,omitempty"`
	// The name of the code list maintained by an external agency.
	ExternalCodeListName *string `json:"externalCodeListName,omitempty"`
	// The version of the code list maintained by an external agency.
	ExternalCodeListVersion *string `json:"externalCodeListVersion,omitempty"`
	// Code list values.
	EnumerationValueInformation []ExternalCodeValue `json:"enumerationValueInformation"`
}
