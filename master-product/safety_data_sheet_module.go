package masterproduct

// SafetyDataSheetModule is a module containing information usually contained
// on a safety data sheet or on a material safety data sheet as it is referred
// to in some target markets.
type SafetyDataSheetModule struct {
	// Trade item information usually contained on a safety data sheet or
	// on a material safety data sheet as it is referred to in some target
	// markets.
	SafetyDataSheetInformation *[]SafetyDataSheetInformation `json:"safetyDataSheetInformation,omitempty"`
}

// SafetyDataSheetInformation contains trade item information usually contained
// on a safety data sheet or on a material safety data sheet as it is referred
// to in some target markets.
type SafetyDataSheetInformation struct {
	// An indicator whether the Trade Item is regulated for shipment by any
	// agency.
	IsRegulatedForTransportation *bool `json:"isRegulatedForTransportation,omitempty"`
	// Details related to the Globally Harmonized System of Classification
	// and Labelling of Chemicals.
	GHSDetail *GHSDetail `json:"gHSDetail,omitempty"`
	// Information on Physical or Chemical Properties for a trade item for
	// example water solubility.
	PhysicalChemicalPropertyInformation *PhysicalChemicalPropertyInformation `json:"physicalChemicalPropertyInformation,omitempty"`
}

// GHSDetail Details related to the Globally Harmonized System of
// Classification and Labelling of Chemicals.
type GHSDetail struct {
	// Words such as "Danger" or "Warning" used to emphasize hazards and
	// indicate the relative level of severity of the hazard. For GHS these
	// are assigned to a GHS hazard class and category. Some lower level
	// hazard categories do not use signal words. Uses code list
	// gHSSignalWordsCode.
	GHSSignalWordsCode *string `json:"gHSSignalWordsCode,omitempty"`
	// A code depicting the symbols which convey health, physical and
	// environmental hazard information, assigned to a hazard class and
	// category for example GHS. Pictograms include the harmonized hazard
	// symbols plus other graphic elements, such as borders, background
	// patterns or colours that are intended to convey specific
	// information. Examples of all the pictograms and downloadable files
	// for GHS can be accessed on the UN website for the GHS. Uses code
	// list gHSSymbolDescriptionCode.
	GHSSymbolDescriptionCode *[]string `json:"gHSSymbolDescriptionCode,omitempty"`
	// Standard phrases describing the nature of a hazard per GHS.
	HazardStatements *[]HazardStatement `json:"hazardStatement,omitempty"`
	// Measures listed on a hazardous label to minimize or prevent adverse
	// effects related to GHS.
	PrecautionaryStatements *[]PrecautionaryStatement `json:"precautionaryStatement,omitempty"`
}

// HazardStatement contains standard phrases describing the nature of a hazard
// per GHS.
type HazardStatement struct {
	// Standard phrases assigned to a hazard class and category that
	// describe the nature of the hazard.
	HazardStatementsCode *string `json:"hazardStatementsCode,omitempty"`
	// A description of standard phrases assigned to a hazard class and
	// category that describe the nature of the hazard.
	HazardStatementsDescriptions *[]MLString `json:"hazardStatementsDescription,omitempty"`
}

// PrecautionaryStatement contains measures listed on a hazardous label to
// minimize or prevent adverse effects related to GHS.
type PrecautionaryStatement struct {
	// Measures listed on a hazardous label to minimize or prevent adverse
	// effects. For GHS, the precautionary statements have been linked to
	// each GHS hazard statement and type of hazard. Precautionary
	// statements for GHS cover prevention, response in cases of accidental
	// spillage or exposure, storage, and disposal.
	PrecautionaryStatementsCode *string `json:"precautionaryStatementsCode,omitempty"`
	// A description of the measures listed on a hazardous label to
	// minimize or prevent adverse effects.
	PrecautionaryStatementsDescriptions *[]MLString `json:"precautionaryStatementsDescription,omitempty"`
}

// PhysicalChemicalPropertyInformation contains information on Physical or
// Chemical Properties for a trade item for example water solubility.
type PhysicalChemicalPropertyInformation struct {
	// Details on a flash point for a trade item.
	FlashPoints *[]FlashPoint `json:"flashPoint,omitempty"`
	// PH is defined as the acidity or alkalinity of an aqueous solution.
	// It is defined as the logarithm of the reciprocal of the hydrogenion
	// concentration of a solution. pH= log10 1/[H+].
	PHInformation *PHInformation `json:"pHInformation,omitempty"`
}

// FlashPoint contains details on a flash point for a trade item.
type FlashPoint struct {
	// The temperature at which a substance gives off a sufficient vapour
	// to support combustion. This uses a measurement consisting of a unit
	// of measure and value. With the above request it requires the flash
	// point not to be the lowest but the point at which flash point occurs
	// and it could be that temperature and lower for some products. The
	// scientific Measurement Precision code would determine that.
	FlashPointTemperatures *[]TemperatureMeasurement `json:"flashPointTemperature"`
}

// PHInformation describes a PH value. PH is defined as the acidity or
// alkalinity of an aqueous solution. It is defined as the logarithm of the
// reciprocal of the hydrogenion concentration of a solution. pH= log10 1/[H+].
type PHInformation struct {
	// The exact PH amount for a chemical ingredient (not a range).
	ExactPH *float64 `json:"exactPH,omitempty"`
	// The maximum range for PH.
	MaximumPH *float64 `json:"maximumPH,omitempty"`
	// The minimum range value for PH.
	MinimumPH *float64 `json:"minimumPH,omitempty"`
}
