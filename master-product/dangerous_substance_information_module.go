package masterproduct

// DangerousSubstanceInformationModule is a module detailing substances that
// can harm people.
type DangerousSubstanceInformationModule struct {
	// Details on substances that can harm people, other living organisms,
	// property, or the environment.
	DangerousSubstanceInformations *[]DangerousSubstanceInformation `json:"dangerousSubstanceInformation.omitempty"`
}

// DangerousSubstanceInformation contains details on substances that can harm
// people, other living organisms, property, or the environment.
type DangerousSubstanceInformation struct {
	// Properties of a dangerous substance.
	DangerousSubstanceProperties *[]DangerousSubstanceProperty `json:"dangerousSubstanceProperties.omitempty"`
}

// DangerousSubstanceProperty details properties of a dangerous substance.
type DangerousSubstanceProperty struct {
	// The name of the type of dangerous substance contained in the trade
	// item.
	DangerousSubstanceName *string `json:"dangerousSubstanceName.omitempty"`
	// An indicator whether or not a trade item is classified and labelled
	// as containing a dangerous substance.
	IsDangerousSubstance *bool `json:"isDangerousSubstance.omitempty"`
	// The abbreviation codes for labelling obligations and special risks
	// (health risks of skin, respiratory organs, swallow, eyes,
	// reproduction) for handling of the substance.
	RiskPhraseCodes *[]ExternalCodeValue `json:"riskPhraseCode.omitempty"`
	// Safety phrases are defined as safety advice concerning dangerous
	// substances and preparations.
	SafetyPhraseCodes *[]ExternalCodeValue `json:"safetyPhraseCode.omitempty"`
}
