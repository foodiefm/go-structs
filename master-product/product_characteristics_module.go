package masterproduct

// ProductCharacteristicsModule is a module used to express characteristics for
// a product for example values for a property such as numberOfPlys.
type ProductCharacteristicsModule struct {
	// A characteristic for a product for example values for a property
	// such as numberOfPlys along with its associated value.
	ProductCharacteristics *[]ProductCharacteristic `json:"productCharacteristics,omitempty"`
}

// ProductCharacteristic describes characteristic for a product for example
// values for a property such as numberOfPlys along with its associated value.
type ProductCharacteristic struct {
	// The name of the product characteristic being described.Uses code
	// list productCharacteristicCode.
	ProductCharacteristicCode string `json:"productCharacteristicCode"`
	// The product characteristic value expressed as a description (text
	// with language).
	ProductCharacteristicValueDescriptions *[]MLString `json:"productCharacteristicValueDescription,omitempty"`
	// The product characteristic value expressed as a string (text value
	// with no language).
	ProductCharacteristicValueString *[]string `json:"productCharacteristicValueString,omitempty"`
}
