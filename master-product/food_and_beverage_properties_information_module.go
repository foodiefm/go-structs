package masterproduct

// FoodAndBeveragePropertiesInformationModule contains information on
// physiochemical or other properties of food and beverage products.
type FoodAndBeveragePropertiesInformationModule struct {
	// Information on the product's physicochemical characteristics.
	PhysiochemicalCharacteristics *[]PhysiochemicalCharacteristic `json:"physiochemicalCharacteristic,omitempty"`
}

// PhysiochemicalCharacteristic is an information on the product's
// physicochemical characteristics.
type PhysiochemicalCharacteristic struct {
	// Code indicating the type of physiochemical characteristic. Uses code
	// list physiochemicalCharacteristicCode.
	PhysiochemicalCharacteristicCode *string `json:"physiochemicalCharacteristicCode,omitempty"`
	// Measurement value of the physicochemical characteristic.
	PhysiochemicalCharacteristicValues *[]Measurement `json:"physiochemicalCharacteristicValue,omitempty"`
}
