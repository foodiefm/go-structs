package masterproduct

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// MasterProduct is a structure containing all global Master Product details.
type MasterProduct struct {
	MasterProductTerse
	Details *MasterProductDetails `json:"details,omitempty"`
}

// MasterProductTerse is a terse version of a Master Product.
type MasterProductTerse struct {
	ID               *uuid.UUID `json:"id"`
	ProductID        *uuid.UUID `json:"product_id"`
	OrganizationID   *uuid.UUID `json:"organization_id"`
	DecommissionedAt *time.Time `json:"decommissioned_at,omitempty"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`

	ExtID string `json:"ext_id"`
	GTIN  string `json:"gtin"`
	Name  string `json:"name"`
}

// MasterProductDetails is the details portion of Master Product.
type MasterProductDetails struct {
	Header    MasterProductHeader    `json:"Header"`
	TradeItem MasterProductTradeItem `json:"tradeItem"`
}

// MasterProductHeader describes data structure level details of Master
// Product.
type MasterProductHeader struct {
	XDataFormatVersion string `json:"x_dataFormatVersion"`
}

// MasterProductTradeItem contains GDSN product details.
type MasterProductTradeItem struct {
	// Global trade item number.
	GTIN *string `json:"gtin,omitempty"`
	// Extra information on Global Trade Item Number to identify a trade item.
	XTradeItemIdentification *TradeItemIdentification `json:"x_tradeItemIdentification,omitempty"`
	// Alternative means to the Global Trade Item Number to identify a trade item.
	AdditionalTradeItemIdentifications *[]TradeItemIdentification `json:"additionalTradeItemIdentification,omitempty"`
	// The identification of a party, by GLN, in a specific party role.
	InformationProviderOfTradeItem *PartyInRole `json:"informationProviderOfTradeItem,omitempty"`
	// Party name and identification information for the manufacturer(s) of the trade item.
	ManufacturerOfTradeItems *[]PartyInRole `json:"manufacturerOfTradeItem,omitempty"`
	// Information specifying the product class to which a trade item belongs and the classification system being applied.
	GDSNTradeItemClassification *GDSNTradeItemClassification `json:"gdsnTradeItemClassification,omitempty"`
	// A trade item referenced by this trade item for example replaced or replaced by.
	ReferencedTradeItems *[]ReferencedTradeItem `json:"referencedTradeItem,omitempty"`
	// Target Market associated with a Trade Item.
	TargetMarkets *[]TargetMarket `json:"targetMarket,omitempty"`
	// Contact details for a Trade Item.
	TradeItemContactInformations *[]TradeItemContactInformation `json:"tradeItemContactInformation,omitempty"`
	// Dates relevant to the process of trade item synchronisation for example publication date.
	TradeItemSynchronisationDates *TradeItemSynchronisationDates `json:"tradeItemSynchronisationDates,omitempty"`
	// Detailed information on the trade item.
	TradeItemInformation TradeItemInformation `json:"tradeItemInformation"`
}

// TradeItemIdentification contains extra information on Global Trade Item
// Number to identify a trade item.
type TradeItemIdentification struct {
	// A trade item identifier that is in addition to the GTIN.
	ID string `json:"$"`
	// This code will be used to cross-reference the Vendors internal trade
	// item number to the GTIN in a one to one relationship.
	AdditionalTradeItemIdentificationTypeCode string `json:"@additionalTradeItemIdentificationTypeCode"`
	// Start Date-Time of the given GTIN in ISO Format.
	StartDateTime *time.Time `json:"@startDateTime,omitempty"`
	// End Date-Time of the given GTIN in ISO Format.
	EndDateTime *time.Time `json:"@endDateTime,omitempty"`
}

// PartyInRole identifier a party, by GLN, in a specific party role.
type PartyInRole struct {
	// The Global Location Number (GLN) is a structured Identification of a
	// physical location, legal or functional entity within an enterprise.
	// The GLN is the primary party identifier. Each party identified in
	// the trading relationship must have a primary party Identification.
	GLN *string `json:"gln,omitempty"`
	// The name of the party expressed in text.
	PartyName *string `json:"partyName,omitempty"`
	// The address associated with the party. This could be the full
	// company address.
	PartyAddress *string `json:"partyAddress,omitempty"`
}

// GDSNTradeItemClassification specify the product class to which a trade item
// belongs and the classification system being applied.
type GDSNTradeItemClassification struct {
	// Code specifying a product category according to the GS1 Global
	// Product Classification (GPC) standard.
	GPCCategoryCode string `json:"gpcCategoryCode"`
	// Category code based on alternate classification schema chosen in
	// addition to the Global Product Classification (GPC).
	AdditionalTradeItemClassifications *[]AdditionalTradeItemClassification `json:"additionalTradeItemClassification,omitempty"`
}

// AdditionalTradeItemClassification contains category code based on alternate
// classification schema chosen in addition to the Global Product
// Classification (GPC).
type AdditionalTradeItemClassification struct {
	// Additional classification system code.
	AdditionalTradeItemClassificationSystemCode *string `json:"additionalTradeItemClassificationSystemCode,omitempty"`
	// A code list value for an Additional Trade Item Classification Type.
	AdditionalTradeItemClassificationValues *[]AdditionalTradeItemClassificationValue `json:"additionalTradeItemClassificationValue,omitempty"`
}

// AdditionalTradeItemClassificationValue is a code list value for an
// Additional Trade Item Classification Type.
type AdditionalTradeItemClassificationValue struct {
	// Category code based on alternate classification schema chosen in
	// addition to GS1 classification.
	AdditionalTradeItemClassificationCodeValue string `json:"additionalTradeItemClassificationCodeValue"`
}

// ReferencedTradeItem is a trade item referenced by this trade item for
// example replaced or replaced by.
type ReferencedTradeItem struct {
	// The identification of the referenced trade item.
	GTIN string `json:"gtin"`
	// A code depicting the type of trade item that is referenced for a
	// specific purpose for example substitute, replaced by, equivalent
	// trade items.
	ReferencedTradeItemTypeCode string `json:"referencedTradeItemTypeCode"`
}

// TargetMarket is target market associated with a Trade Item.
type TargetMarket struct {
	// The code that identifies the target market. The taget market is at
	// country level or higher geographical definition and is where a trade
	// item is intended to be sold.
	TargetMarketCountryCode string `json:"targetMarketCountryCode"`
}

// TradeItemContactInformation is a contact details for a Trade Item.
type TradeItemContactInformation struct {
	// The general category of the contact party for a trade item for
	// example Purchasing.
	ContactTypeCode string `json:"contactTypeCode"`
	// The address associated with the contact type. For example, in case
	// of a contact type of CONSUMER_SUPPORT, this could be the full
	// company address as expressed on the trade item packaging or label.
	ContactAddress *string `json:"contactAddress,omitempty"`
	// A description of the contact for the trade item.
	ContactDescriptions *[]MLString `json:"contactDescription,omitempty"`
	// The name of the company or person associated with the contact type.
	// For example, in case of a contact type of CONSUMER_SUPPORT, this
	// could be the company name as expressed on the trade item packaging
	// or label.
	ContactName *string `json:"contactName,omitempty"`
	// The communication channel for example phone number for a target
	// market for a Trade Item.
	TargetMarketCommunicationChannels *[]TargetMarketCommunicationChannel `json:"targetMarketCommunicationChannel,omitempty"`
}

// TargetMarketCommunicationChannel is the communication channel for example
// phone number for a target market for a Trade Item.
type TargetMarketCommunicationChannel struct {
	// A target market associated with a communication channel for example
	// Canada.
	TargetMarkets *[]TargetMarket `json:"targetMarket.omitempty"`
	// The channel or manner in which a communication can be made, such as
	// telephone or email.
	CommunicationChannels []CommunicationChannel `json:"communicationChannel,omitempty"`
}

// CommunicationChannel is the channel or manner in which a communication can
// be made, such as telephone or email.
type CommunicationChannel struct {
	// The channel or manner in which a communication can be made, such as
	// telephone or email.
	CommunicationChannelCode string `json:"communicationChannelCode"`
	// The channel or manner in which a communication can be made, such as
	// telephone or email.
	CommunicationValue string `json:"communicationValue"`
	// The channel or manner in which a communication can be made, such as
	// telephone or email.
	CommunicationChannelName *string `json:"communicationChannelName,omitempty"`
}

// TradeItemSynchronisationDates contains relevant dates to the process of
// trade item synchronisation for example publication date.
type TradeItemSynchronisationDates struct {
	// Indicates the point in time where the last modification on a Trade
	// Item was made.
	LastChangeDateTime *time.Time `json:"lastChangeDateTime,omitempty"`
}

// TradeItemInformation is a detailed information on the trade item.
type TradeItemInformation struct {
	Extension TradeItemExtension `json:"extensions"`
}

// TradeItemExtension contains references to additional information modules.
type TradeItemExtension struct {
	// A module containing details on products traditionally containing
	// alcohol.
	AlcoholInformationModule *AlcoholInformationModule `json:"alcoholInformationModule,omitempty"`
	// A module containing information on allergens for a trade item.
	AllergenInformationModule *AllergenInformationModule `json:"allergenInformationModule,omitempty"`
	// A module contain instructions on how the consumer is to use or store
	// a trade item.
	ConsumerInstructionsModule *ConsumerInstructionsModule `json:"consumerInstructionsModule,omitempty"`
	// A module detailing content and processing related information
	// specific to dairy, fish, meat and poultry products.
	DairyFishMeatPoultryItemModule *DairyFishMeatPoultryItemModule `json:"dairyFishMeatPoultryItemModule"`
	// A module detailing substances that can harm people.
	DangerousSubstanceInformationModule *DangerousSubstanceInformationModule `json:"dangerousSubstanceInformationModule,omitempty"`
	// A module contain a product dietary suitability.
	DietInformationModule *DietInformationModule `json:"dietInformationModule,omitempty"`
	// Information on any farming or processing performed on and
	// agricultural trade item.
	FarmingAndProcessingInformationModule *FarmingAndProcessingInformationModule `json:"farmingAndProcessingInformationModule,omitempty"`
	// Information on the constituent ingredient make up of the product.
	FoodAndBeverageIngredientModule *FoodAndBeverageIngredientModule `json:"foodAndBeverageIngredientModule,omitempty"`
	// Information on way the product can be prepared or served.
	FoodAndBeveragePreparationServingModule *FoodAndBeveragePreparationServingModule `json:"foodAndBeveragePreparationServingModule,omitempty"`
	// Information on physiochemical or other properties of food and
	// beverage products.
	FoodAndBeveragePropertiesInformationModule *FoodAndBeveragePropertiesInformationModule `json:"foodAndBeveragePropertiesInformationModule,omitempty"`
	// Information on a trade item meant to convey features and benefits
	// and targeted customer.
	MarketingInformationModule *MarketingInformationModule `json:"marketingInformationModule,omitempty"`
	// A module providing Information on ingredients for items that are not
	// food for example detergents, medicines.
	NonfoodIngredientModule *NonfoodIngredientModule `json:"nonfoodIngredientModule,omitempty"`
	// Information about content of nutrients. Multiple sets of nutrient
	// information can be specified with varying state, serving size and
	// daily value intake base.
	NutritionalInformationModule *NutritionalInformationModule `json:"nutritionalInformationModule,omitempty"`
	// Packaging information for a trade item.
	PackagingInformationModule *PackagingInformationModule `json:"packagingInformationModule,omitempty"`
	// A module containing details on markings on the packaging of the
	// trade item for example dates, environment.
	PackagingMarkingModule *PackagingMarkingModule `json:"packagingMarkingModule,omitempty"`
	// Information on the activity (e.g. bottling) taken place for a trade
	// item as well as the associated geographic area.
	PlaceOfItemActivityModule *PlaceOfItemActivityModule `json:"placeOfItemActivityModule,omitempty"`
	// A module used to express characteristics for a product for example
	// values for a property such as numberOfPlys.
	ProductCharacteristicsModule *ProductCharacteristicsModule `json:"productCharacteristicsModule,omitempty"`
	// A module containing information usually contained on a safety data
	// sheet or on a material safety data sheet as it is referred to in
	// some target markets.
	SafetyDataSheetModule *SafetyDataSheetModule `json:"safetyDataSheetModule,omitempty"`
	// Sales information regarding price and selling
	// conditions/restrictions of the Trade Item to the consumer.
	SalesInformationModule *SalesInformationModule `json:"salesInformationModule,omitempty"`
	// A module carrying general descriptions of the trade item including
	// brand, form, variant.
	TradeItemDescriptionModule *TradeItemDescriptionModule `json:"tradeItemDescriptionModule,omitempty"`
	// A module containing information on the amount of time the item can
	// or should be used, sold, etc.
	TradeItemLifespanModule *TradeItemLifespanModule `json:"tradeItemLifespanModule,omitempty"`
	// A module containing measurement information for the trade item.
	TradeItemMeasurementsModule *TradeItemMeasurementsModule `json:"tradeItemMeasurementsModule,omitempty"`
	// Information on temperature considerations for trade item.
	TradeItemTemperatureInformationModule *TradeItemTemperatureInformationModule `json:"tradeItemTemperatureInformationModule,omitempty"`
	// A module with information specific to variable weight or dimension
	// trade items.
	VariableTradeItemInformationModule *VariableTradeItemInformationModule `json:"variableTradeItemInformationModule,omitempty"`
	// Associated code lists.
	DGCodeListModule *DGCodeListModule `json:"dgCodeListModule,omitempty"`
	// Product media properties.
	DGMediaModule *DGMediaModule `json:"dgMediaModule,omitempty"`
	// Product presentation properties.
	DGPresentationModule *DGPresentationModule `json:"dgPresentationModule,omitempty"`
	// Private use module. dgPrivateUseModule must be a hash or nil. When
	// hash, dgPrivateUseModule may contain any number of any kind of keys
	// with any kind of values.
	DGPrivateUseModule *json.RawMessage `json:"dgPrivateUseModule,omitempty"`
	// A module containing freely defined product attributes.
	DGProductAttributeModule *DGProductAttributeModule `json:"dgProductAttributeModule,omitempty"`
}
