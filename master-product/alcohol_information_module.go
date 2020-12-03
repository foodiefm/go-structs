package masterproduct

// AlcoholInformationModule is a module containing details on products
// traditionally containing alcohol.
type AlcoholInformationModule struct {
	// Details on products traditionally containing alcohol.
	AlcoholInformation *AlcoholInformation `json:"alcoholInformation,omitempty"`
}

// AlcoholInformation describes details on products traditionally containing
// alcohol.
type AlcoholInformation struct {
	// Percentage of alcohol contained in the base unit trade item.
	PercentageOfAlcoholByVolume *float64 `json:"percentageOfAlcoholByVolume,omitempty"`
	// Indication of the amount of sugar contained in the beverage for
	// example if sugar remaining equals 6.5 g/l then enter 6.5 GL.
	AlcoholicBeverageSugarContents *[]AlcoholicBeverageSugarContent `json:"alcoholicBeverageSugarContent,omitempty"`
}

// AlcoholicBeverageSugarContent indicates of the amount of sugar contained in
// the beverage for example if sugar remaining equals 6.5 g/l then enter
// 6.5 GL.
type AlcoholicBeverageSugarContent struct {
	// Measurement value.
	Measurement float64 `json:"$"`
	// Unit of measure code. Uses code list measurementUnitCode.
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}
