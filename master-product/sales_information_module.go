package masterproduct

// SalesInformationModule describes sales information regarding price and
// selling conditions/restrictions of the Trade Item to the consumer.
type SalesInformationModule struct {
	// Restrictions or requirements on the retailer for sales of the Trade
	// Item to the consumer.
	SalesInformation *SalesInformation `json:"salesInformation,omitempty"`
}

// SalesInformation describes restrictions or requirements on the retailer for
// sales of the Trade Item to the consumer.
type SalesInformation struct {
	// A code depicting restrictions imposed on the Trade Item regarding
	// how it can be sold to the consumer for example Prescription
	// Required. Uses code list consumerSalesConditionCode.
	ConsumerSalesConditionCode *[]string `json:"consumerSalesConditionCode,omitempty"`
	// Indicator to show how a product is sold. Uses code list
	// priceByMeasureTypeCode.
	PriceByMeasureTypeCode *string `json:"priceByMeasureTypeCode,omitempty"`
	// The quantity of the product at usage. Applicable for concentrated
	// products and products where the comparison price is calculated based
	// on a measurement other than netContent.
	PriceComparisonMeasurements *[]Measurement `json:"priceComparisonMeasurement,omitempty"`
	// Describes the measurement used for selling unit of the Trade Item to
	// the end consumer.
	SellingUnitOfMeasure *string `json:"sellingUnitOfMeasure,omitempty"`
	// Defines compliancy with EU 1169 regulation.
	XEU1169Compliance *XEU1169Compliance `json:"x_eu1169Compliance,omitempty,omitempty"`
	// An indicator whether or not the Trade Item is excluded from loyalty
	// programs.
	XIsExcludedFromLoyaltyPrograms *bool `json:"x_isExcludedFromLoyaltyPrograms,omitempty"`
	// Defines how much the quantity of a Trade Item is changed when
	// additional items are added or removed from shopping basket.
	XSellingContentIncrement *float64 `json:"x_sellingContentIncrement,omitempty"`
	// Defines the initial quantity of Trade Item when the first instance
	// of the item is added to shopping basket.
	XSellingContentInitial *float64 `json:"x_sellingContentInitial,omitempty"`
	// Defines sales restrictions tags.
	XSalesRestrictionTags *[]string `json:"x_salesRestrictionTags,omitempty"`
	// Defines the measurement unit code used for selling of the Trade Item
	// to the end consumer. Uses code list sellingUnitOfMeasure.
	XSellingUnitOfMeasureCode *string `json:"x_sellingUnitOfMeasureCode,omitempty"`
	// Devices measure types by which Trade Item is sold.
	XSoldByMeasureTypeCode *string `json:"x_soldByMeasureTypeCode,omitempty"`
}

// XEU1169Compliance defines compliancy with EU 1169 regulation.
type XEU1169Compliance struct {
	// Regulation compliancy. Uses code list x_complianceCode.
	XComplianceCode string `json:"x_complianceCode"`
}
