package structs

import (
	"encoding/json"
	"time"
)

type MasterProductData struct {
	ID        string                 `json:"id"`
	ExtID     string                 `json:"ext_id"`
	ProductID string                 `json:"product_id"`
	Gtin      string                 `json:"gtin"`
	Name      string                 `json:"name"`
	Header    MasterProductHeaders   `json:"Header"`
	TradeItem MasterProductTradeItem `json:"tradeItem"`
}

type MasterProductHeaders struct {
	XDataFormatVersion string `json:"x_dataFormatVersion"`
}

type MasterProductTradeItem struct {
	GTIN                               string                              `json:"gtin"`                              //Global trade item number
	XTradeItemIdentification           TradeItemIdentification             `json:"x_tradeItemIdentification"`         // Extra information on Global Trade Item Number to identify a trade item.
	AdditionalTradeItemIdentifications []AdditionalTradeItemIdentification `json:"additionalTradeItemIdentification"` // Alternative means to the Global Trade Item Number to identify a trade item.
	InformationProviderOfTradeItem     InformationProviderOfTradeItem      `json:"informationProviderOfTradeItem"`    // The identification of a party, by GLN, in a specific party role.
	ManufacturerOfTradeItems           []ManufacturerOfTradeItem           `json:"manufacturerOfTradeItem"`           // Party name and identification information for the manufacturer(s) of the trade item.
	GdsnTradeItemClassification        GdsnTradeItemClassification         `json:"gdsnTradeItemClassification"`       // Information specifying the product class to which a trade item belongs and the classification system being applied.
	ReferencedTradeItems               []ReferencedTradeItem               `json:"referencedTradeItem"`               // A trade item referenced by this trade item for example replaced or replaced by.
	TargetMarkets                      TradeItemTargetMarket               `json:"targetMarket"`                      // Target Market associated with a Trade Item.
	TradeItemContactInformations       []TradeItemContactInformation       `json:"tradeItemContactInformation"`       // Contact details for a Trade Item.
	TradeItemSynchronisationDates      TradeItemSynchronisationDates       `json:"tradeItemSynchronisationDates"`     // Dates relevant to the process of trade item synchronisation for example publication date.
	TradeItemInformation               TradeItemInformation                `json:"tradeItemInformation"`              // Detailed information on the trade item.
}

// TradeItemIdentification contains extra information on Global Trade Item Number to identify a trade item.
type TradeItemIdentification struct {
	// A trade item identifier that is in addition to the GTIN.
	ID string `json:"$"`
	// This code will be used to cross-reference the Vendors internal trade item number to the GTIN in a one to one relationship.
	AdditionalTradeItemIdentificationTypeCode string `json:"@additionalTradeItemIdentificationTypeCode"`
	// Start Date-Time of the given GTIN in ISO Format
	StartDateTime time.Time `json:"@startDateTime"`
	// End Date-Time of the given GTIN in ISO Format
	EndDateTime time.Time `json:"@endDateTime"`
}

// AdditionalTradeItemIdentification describes alternative means to the Global Trade Item Number to identify a trade item.
type AdditionalTradeItemIdentification struct {
	// A trade item identifier that is in addition to the GTIN.
	ID string `json:"$"`
	// This code will be used to cross-reference the Vendors internal trade item number to the GTIN in a one to one relationship.
	AdditionalTradeItemIdentificationTypeCode string `json:"@additionalTradeItemIdentificationTypeCode"`
	// Start Date-Time of the given GTIN in ISO Format
	StartDateTime time.Time `json:"@startDateTime"`
	// End Date-Time of the given GTIN in ISO Format
	EndDateTime time.Time `json:"@endDateTime"`
	// The snapshot of the code list at a certain point in time.
	Version string `json:"@version"`
}

// InformationProviderOfTradeItem identifies a party, by GLN, in a specific party role.
type InformationProviderOfTradeItem struct {
	// The Global Location Number (GLN) is a structured Identification of a physical location,
	// legal or functional entity within an enterprise. The GLN is the primary party identifier.
	// Each party identified in the trading relationship must have a primary party Identification.
	GLN string `json:"gln"`
	// The name of the party expressed in text.
	PartyName string `json:"partyName"`
	// The address associated with the party. This could be the full company address.
	PartyAddress string `json:"partyAddress"`
}

// ManufacturerOfTradeItem contains party name and identification information for the manufacturer(s) of the trade item.
type ManufacturerOfTradeItem struct {
	// The Global Location Number (GLN) is a structured Identification of a physical
	// location, legal or functional entity within an enterprise. The GLN is the primary
	// party identifier. Each party identified in the trading relationship must have
	// a primary party Identification.
	GLN string `json:"gln"`
	// The name of the party expressed in text.
	PartyName string `json:"partyName"`
	// The address associated with the party. This could be the full company address.
	PartyAddress string `json:"partyAddress"`
}

// GdsnTradeItemClassification specify the product class to which a trade item belongs and the classification system being applied.
type GdsnTradeItemClassification struct {
	// Code specifying a product category according to the GS1 Global Product Classification (GPC) standard.
	GpcCategoryCode string `json:"gpcCategoryCode"`
	// Category code based on alternate classification schema chosen in addition to the Global Product Classification (GPC).
	AdditionalTradeItemClassifications []AdditionalTradeItemClassification `json:"additionalTradeItemClassification"`
}

// AdditionalTradeItemClassification contains category code based on alternate classification
// schema chosen in addition to the Global Product Classification (GPC).
type AdditionalTradeItemClassification struct {
	// Additional classification system code.
	AdditionalTradeItemClassificationSystemCode string `json:"additionalTradeItemClassificationSystemCode"`
	// A code list value for an Additional Trade Item Classification Type.
	AdditionalTradeItemClassificationValues []AdditionalTradeItemClassificationValue `json:"additionalTradeItemClassificationValue"`
}

// AdditionalTradeItemClassificationValue is a code list value for an Additional Trade Item Classification Type.
type AdditionalTradeItemClassificationValue struct {
	// Category code based on alternate classification schema chosen in addition to GS1 classification.
	AdditionalTradeItemClassificationCodeValue string `json:"additionalTradeItemClassificationCodeValue"`
}

// ReferencedTradeItem is a trade item referenced by this trade item for example replaced or replaced by.
type ReferencedTradeItem struct {
	// The identification of the referenced trade item.
	GTIN string `json:"gtin"`
	// A code depicting the type of trade item that is referenced for a specific purpose for example
	// substitute, replaced by, equivalent trade items.
	ReferencedTradeItemTypeCode string `json:"referencedTradeItemTypeCode"`
}

// TradeItemTargetMarket is target market associated with a Trade Item.
type TradeItemTargetMarket struct {
	// The code that identifies the target market. The taget market is at country level or
	// higher geographical definition and is where a trade item is intended to be sold.
	TargetMarketCountryCode string `json:"targetMarketCountryCode"`
}

// TradeItemContactInformation is a contact details for a Trade Item.
type TradeItemContactInformation struct {
	// The general category of the contact party for a trade item for example Purchasing.
	ContactTypeCode string `json:"contactTypeCode"`
	// The address associated with the contact type. For example, in case of a contact
	// type of CONSUMER_SUPPORT, this could be the full company address as expressed
	// on the trade item packaging or label.
	ContactAddress string `json:"contactAddress"`
	// A description of the contact for the trade item.
	ContactDescriptions []ContactDescription `json:"contactDescription"`
	// The name of the company or person associated with the contact type. For example,
	// in case of a contact type of CONSUMER_SUPPORT, this could be the company name as
	// expressed on the trade item packaging or label.
	ContactName string `json:"contactName"`
	// The communication channel for example phone number for a target market for a Trade Item.
	TargetMarketCommunicationChannels []TargetMarketCommunicationChannel `json:"targetMarketCommunicationChannel"`
}

// ContactDescription is a description of the contact for the trade item.
type ContactDescription struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// TargetMarketCommunicationChannel is the communication channel for example phone number for a target market for a Trade Item.
type TargetMarketCommunicationChannel struct {
	// A target market associated with a communication channel for example Canada.
	TargetMarkets []CommunicationChannelTargetMarket `json:"targetMarket"`
	// The channel or manner in which a communication can be made, such as telephone or email.
	CommunicationChannels []CommunicationChannel `json:"communicationChannel"`
}

// CommunicationChannelTargetMarket  associated with a communication channel for example Canada.
type CommunicationChannelTargetMarket struct {
	// The code that identifies the target market. The target market is at country level or higher geographical
	// definition and is where a trade item is intended to be sold.
	TargetMarketCountryCode string `json:"targetMarketCountryCode"`
}

// CommunicationChannel is the channel or manner in which a communication can be made, such as telephone or email.
type CommunicationChannel struct {
	// The channel or manner in which a communication can be made, such as telephone or email.
	CommunicationChannelCode string `json:"communicationChannelCode"`
	// The channel or manner in which a communication can be made, such as telephone or email.
	CommunicationValue string `json:"communicationValue"`
	// The channel or manner in which a communication can be made, such as telephone or email.
	CommunicationChannelName string `json:"communicationChannelName"`
}

// TradeItemSynchronisationDates contains relevant dates to the process of trade item synchronisation for example publication date.
type TradeItemSynchronisationDates struct {
	// Indicates the point in time where the last modification on a Trade Item was made.
	LastChangeDateTime time.Time `json:"lastChangeDateTime"`
}

// TradeItemInformation is a detailed information on the trade item.
type TradeItemInformation struct {
	// DG private use module. dgPrivateUseModule must be a hash or nil. When hash,
	// dgPrivateUseModule may contain any number of any kind of keys with any kind of values.
	Extension TradeItemExtension `json:"extensions"`
}

// TradeItemExtension is a DG private use module. dgPrivateUseModule must be a hash or nil. When hash,
// dgPrivateUseModule may contain any number of any kind of keys with any kind of values.
type TradeItemExtension struct {
	// A module containing details on products traditionally containing alcohol.
	AlcoholInformationModule AlcoholInformationModule `json:"alcoholInformationModule"`
	// A module containing information on allergens for a trade item.
	AllergenInformationModule AllergenInformationModule `json:"allergenInformationModule"`
	// A module contain instructions on how the consumer is to use or store a trade item.
	ConsumerInstructionsModule ConsumerInstructionsModule `json:"consumerInstructionsModule"`
	// A module detailing substances that can harm people.
	DangerousSubstanceInformationModule DangerousSubstanceInformationModule `json:"dangerousSubstanceInformationModule"`
	// A module contain a product dietary suitability.
	DietInformationModule DietInformationModule `json:"dietInformationModule"`
	// Information on any farming or processing performed on and agricultural trade item.
	FarmingAndProcessingInformationModule FarmingAndProcessingInformationModule `json:"farmingAndProcessingInformationModule"`
	// Information on the constituent ingredient make up of the product.
	FoodAndBeverageIngredientModule FoodAndBeverageIngredientModule `json:"foodAndBeverageIngredientModule"`
	// Information on way the product can be prepared or served.
	FoodAndBeveragePreparationServingModule FoodAndBeveragePreparationServingModule `json:"foodAndBeveragePreparationServingModule"`
	// Information on physiochemical or other properties of food and beverage products.
	FoodAndBeveragePropertiesInformationModule FoodAndBeveragePropertiesInformationModule `json:"foodAndBeveragePropertiesInformationModule"`
	// Information on a trade item meant to convey features and benefits and targeted customer.
	MarketingInformationModule MarketingInformationModule `json:"marketingInformationModule"`
	// A module providing Information on ingredients for items that are not food for
	// example detergents, medicines.
	NonfoodIngredientModule NonfoodIngredientModule `json:"nonfoodIngredientModule"`
	// Information about content of nutrients. Multiple sets of nutrient information
	// can be specified with varying state, serving size and daily value intake base.
	NutritionalInformationModule NutritionalInformationModule `json:"nutritionalInformationModule"`
	// Packaging information for a trade item.
	PackagingInformationModule PackagingInformationModule `json:"packagingInformationModule"`
	// A module containing details on markings on the packaging of the trade item for
	//  example dates, environment.
	PackagingMarkingModule PackagingMarkingModule `json:"packagingMarkingModule"`
	// Information on the activity (e.g. bottling) taken place for a trade item
	// as well as the associated geographic area.
	PlaceOfItemActivityModule PlaceOfItemActivityModule `json:"placeOfItemActivityModule"`
	// A module used to express characteristics for a product for example values for
	//  a property such as numberOfPlys.
	ProductCharacteristicsModule ProductCharacteristicsModule `json:"productCharacteristicsModule"`
	// A module containing information usually contained on a safety data sheet
	//  or on a material safety data sheet as it is referred to in some target
	//  markets.
	SafetyDataSheetModule SafetyDataSheetModule `json:"safetyDataSheetModule"`
	// Sales information regarding price and selling conditions/restrictions
	// of the Trade Item to the consumer.
	SalesInformationModule SalesInformationModule `json:"salesInformationModule"`
	// A module carrying general descriptions of the trade item including
	//  brand, form, variant.
	TradeItemDescriptionModule TradeItemDescriptionModule `json:"tradeItemDescriptionModule"`
	// A module containing information on the amount of time the item can or should
	// be used, sold, etc.
	TradeItemLifespanModule TradeItemLifespanModule `json:"tradeItemLifespanModule"`
	// A module containing measurement information for the trade item.
	TradeItemMeasurementsModule TradeItemMeasurementsModule `json:"tradeItemMeasurementsModule"`
	// Information on temperature considerations for trade item.
	TradeItemTemperatureInformationModule TradeItemTemperatureInformationModule `json:"tradeItemTemperatureInformationModule"`
	// A module with information specific to variable weight or dimension trade items.
	VariableTradeItemInformationModule VariableTradeItemInformationModule `json:"variableTradeItemInformationModule"`
	// Associated code lists
	DGCodeListModule DGCodeListModule `json:"dgCodeListModule"`
	// Product media properties
	DGMediaModule DGMediaModule `json:"dgMediaModule"`
	// Product presentation properties
	DGPresentationModule DGPresentationModule `json:"dgPresentationModule"`
	// Free data.
	DGPrivateUseModule *json.RawMessage `json:"dgPrivateUseModule"`
	// A module containing freely defined product attributes.
	DGProductAttributeModule DGProductAttributeModule `json:"dgProductAttributeModule"`
}

// AlcoholInformationModule is a module containing details on products traditionally containing alcohol.
type AlcoholInformationModule struct {
	// Details on products traditionally containing alcohol.
	AlcoholInformation AlcoholInformation `json:"alcoholInformation"`
}

// AlcoholInformation describes details on products traditionally containing alcohol.
type AlcoholInformation struct {
	// Percentage of alcohol contained in the base unit trade item.
	PercentageOfAlcoholByVolume float64 `json:"percentageOfAlcoholByVolume"`
	// Indication of the amount of sugar contained in the beverage for example if sugar remaining equals 6.5 g/l then enter 6.5 GL.
	AlcoholicBeverageSugarContents []AlcoholicBeverageSugarContent `json:"alcoholicBeverageSugarContent"`
}

// AlcoholicBeverageSugarContent indicates of the amount of sugar contained in the beverage for example if sugar remaining equals 6.5 g/l then enter 6.5 GL.
type AlcoholicBeverageSugarContent struct {
	// Measurement value.
	Measurement float64 `json:"$"`
	// Unit of measure code. Uses code list measurementUnitCode.
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}

// AllergenInformationModule is a module containing information on allergens for a trade item.
type AllergenInformationModule struct {
	// Information on substances that might cause allergic reactions and substances subject to intolerance
	// when consumed. The allergy information refers to specified regulations that apply to the target market
	// to which the item information is published.
	AllergenRelatedInformations []AllergenRelatedInformation `json:"allergenRelatedInformation"`
}

// AllergenRelatedInformation contains information on substances that might cause allergic reactions
//  and substances subject to intolerance when consumed. The allergy information refers to specified
// regulations that apply to the target market to which the item information is published.
type AllergenRelatedInformation []struct {
	// Agency that controls the allergen definition.
	AllergenSpecificationAgency string `json:"allergenSpecificationAgency"`
	// Free text field containing the name and version of the regulation or standard that
	// contains the definition of the allergen.
	AllergenSpecificationName string `json:"allergenSpecificationName"`
	// Textual description of the presence or absence of allergens as governed by local rules
	// and regulations, specified as one string.
	AllergenStatements []AllergenStatement `json:"allergenStatement"`
	// Description of the presence or absence of allergens as governed by local rules and regulations, specified per allergen.
	Allergens []Allergen `json:"allergen"`
}

// AllergenStatement is a textual description of the presence or absence of allergens as governed
// by local rules and regulations, specified as one string.
type AllergenStatement struct {
	// Name of allergen
	Name string `json:"$"`
	// Language code
	LanguageCode string `json:"@languageCode"`
	// Substring emphasis. Emphases may overlap.
	XEmphasis []XEmphasis `json:"x_emphasis"`
}

// XEmphasis is a substring emphasis. Emphases may overlap.
type XEmphasis struct {
	// Emphasis starting index in characters from the beginning
	// of the string. Index starts at zero.
	StartAt int `json:"startAt"`
	// Emphasis length in characters.
	Length int `json:"length"`
}

// Allergen is a description of the presence or absence of allergens as governed by
// local rules and regulations, specified per allergen.
type Allergen struct {
	// Code indicating the type of allergen. Uses code list allergenTypeCode.
	AllergenTypeCode string `json:"allergenTypeCode"`
	// Code indicating the level of presence of the allergen.
	LevelOfContainmentCode string `json:"levelOfContainmentCode"`
}

// ConsumerInstructionsModule is a module contain instructions on how the consumer is to
// use or store a trade item.
type ConsumerInstructionsModule struct {
	// Instructions on how the consumer is to use or store a trade item.
	ConsumerInstructions ConsumerInstructions `json:"consumerInstructions"`
}

// ConsumerInstructions contains instructions on how the consumer is to use or store a trade item.
type ConsumerInstructions struct {
	// Expresses in text the consumer storage instructions of a product which are normally held on the
	// label or accompanying the product. This information may or may not be labeled on the pack. Instructions
	// may refer to a suggested storage temperature, a specific storage requirement.
	ConsumerStorageInstructions []ConsumerStorageInstruction `json:"consumerStorageInstructions"`
	// Expresses in text the consumer usage instructions of a product which are normally held on the label or accompanying the product. This information may or may not be labeled on the pack. Instructions may refer to a the how the consumer is to use the product, This does not include storage, food preparations, and drug dosage and preparation instructions.
	ConsumerUsageInstructions []ConsumerUsageInstruction `json:"consumerUsageInstructions"`
}

// ConsumerStorageInstruction expresses in text the consumer storage instructions of a product
// which are normally held on the label or accompanying the product. This information may or may
// not be labeled on the pack. Instructions may refer to a suggested storage temperature, a specific
//  storage requirement.
type ConsumerStorageInstruction struct {
	Instruction  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// ConsumerUsageInstruction expresses in text the consumer usage instructions of a product which are normally
//  held on the label or accompanying the product. This information may or may not be labeled on the pack.
// Instructions may refer to a the how the consumer is to use the product, This does not include storage, food
// preparations, and drug dosage and preparation instructions.
type ConsumerUsageInstruction struct {
	Instruction  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// DangerousSubstanceInformationModule is a module detailing substances that can harm people.
type DangerousSubstanceInformationModule struct {
	// Details on substances that can harm people, other living organisms, property, or the environment.
	DangerousSubstanceInformations []DangerousSubstanceInformation `json:"dangerousSubstanceInformation"`
}

// DangerousSubstanceInformation contains details on substances that can harm people, other living
// organisms, property, or the environment.
type DangerousSubstanceInformation struct {
	// Properties of a dangerous substance.
	DangerousSubstanceProperties []DangerousSubstanceProperty `json:"dangerousSubstanceProperties"`
}

// DangerousSubstanceProperty details properties of a dangerous substance.
type DangerousSubstanceProperty struct {
	// The name of the type of dangerous substance contained in the trade item.
	DangerousSubstanceName string `json:"dangerousSubstanceName"`
	// An indicator whether or not a trade item is classified and labelled as containing
	// a dangerous substance.
	IsDangerousSubstance bool `json:"isDangerousSubstance"`
	// The abbreviation codes for labelling obligations and special risks (health risks
	// of skin, respiratory organs, swallow, eyes, reproduction) for handling of the substance.
	RiskPhraseCodes []RiskPhraseCode `json:"riskPhraseCode"`
	// Safety phrases are defined as safety advice concerning dangerous substances and preparations.
	SafetyPhraseCodes []SafetyPhraseCode `json:"safetyPhraseCode"`
}

// RiskPhraseCode is the abbreviation codes for labelling obligations and special risks
// (health risks of skin, respiratory organs, swallow, eyes, reproduction) for handling
// of the substance.
type RiskPhraseCode struct {
	// The name of the agency that manages a code list.
	ExternalAgencyName string `json:"externalAgencyName"`
	// The name of the code list maintained by an external agency.
	ExternalCodeListName         string                            `json:"externalCodeListName"`
	EnumerationValueInformations []RiskEnumerationValueInformation `json:"enumerationValueInformation"`
}

// RiskEnumerationValueInformation cotains about risk phares codes
type RiskEnumerationValueInformation struct {
	// Code List Value maintained by an external code list agency.
	EnumerationValue string `json:"enumerationValue"`
}

// SafetyPhraseCode defines safety advice concerning dangerous substances and preparations.
type SafetyPhraseCode struct {
	// The name of the agency that manages a code list.
	ExternalAgencyName string `json:"externalAgencyName"`
	// The name of the code list maintained by an external agency.
	ExternalCodeListName         string                              `json:"externalCodeListName"`
	EnumerationValueInformations []SafetyEnumerationValueInformation `json:"enumerationValueInformation"`
}

// SafetyEnumerationValueInformation of safety phrases
type SafetyEnumerationValueInformation struct {
	// Code List Value maintained by an external code list agency.
	EnumerationValue string `json:"enumerationValue"`
}

// DietInformationModule is a module contain a product dietary suitability.
type DietInformationModule struct {
	// The diet the product is suitable for.
	DietInformation DietInformation `json:"dietInformation"`
}

// DietInformation is the diet the product is suitable for.
type DietInformation struct {
	// Expresses in text the dietary description of a product which are normally held
	// on the label or accompanying the product. This information may or may not be labeled
	// on the pack. Instructions may refer to a suggested lifestyle or dietary preference.
	DietTypeDescriptions []DietTypeDescription `json:"dietTypeDescription"`
	DietTypeInformations []DietTypeInformation `json:"dietTypeInformation"`
}

// DietTypeDescription expresses in text the dietary description of a product which are
// normally held on the label or accompanying the product. This information may or may not
// be labeled on the pack. Instructions may refer to a suggested lifestyle or dietary preference.
type DietTypeDescription struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// DietTypeInformation Expresses in text the suggested dietary suitability of a product which
// are normally held on the label or accompanying the product. This information may or may not
// be labeled on the pack.
type DietTypeInformation struct {
	DietTypeCode    string `json:"dietTypeCode"`
	DietTypeSubcode string `json:"dietTypeSubcode"`
}

// FarmingAndProcessingInformationModule contains information on any farming or processing
// performed on and agricultural trade item.
type FarmingAndProcessingInformationModule struct {
	// Details on the trade item regarding the extent of organic production.
	TradeItemOrganicInformation TradeItemOrganicInformation `json:"tradeItemOrganicInformation"`
	// 	Information on farming and processing for a trade item.
	TradeItemFarmingAndProcessing TradeItemFarmingAndProcessing `json:"tradeItemFarmingAndProcessing"`
	// Attribute value pair information.
	AVPList AVPList `json:"avpList"`
}

// TradeItemOrganicInformation details on the trade item regarding the extent of organic production.
type TradeItemOrganicInformation struct {
	// Indication of the place where the agricultural raw materials of which the product is composed
	// have been farmed. It applies only to the trade item, not ingredient by ingredient. Uses code
	// list organicProductPlaceOfFarmingCode
	OrganicProductPlaceOfFarmingCode string `json:"organicProductPlaceOfFarmingCode"`
	// Any claim to indicate the organic status of a trade item or of one or more of its components.
	OrganicClaims []OrganicClaim `json:"organicClaim"`
}

// OrganicClaim contains any claim to indicate the organic status of a trade item or of one or more of its components.
type OrganicClaim struct {
	// A Governing body that creates and maintains standards related to organic products. Uses code list organicClaimAgencyCode
	OrganicClaimAgencyCode []string `json:"organicClaimAgencyCode"`
	// The percent of actual organic materials per weight of the trade item. This is usually claimed on the product
	OrganicPercentClaim int `json:"organicPercentClaim"`
}

// TradeItemFarmingAndProcessing contains information on farming and processing for a trade item.
type TradeItemFarmingAndProcessing struct {
	// A statement of the presence or absence of genetically modified protein or DNA. Uses code
	// list geneticallyModifiedDeclarationCode
	GeneticallyModifiedDeclarationCode string `json:"geneticallyModifiedDeclarationCode"`
	// Code value indicating the preservation technique used to preserve the product from
	// deterioration. Uses code list preservationTechniqueCode.
	PreservationTechniqueCode []string `json:"preservationTechniqueCode"`
}

// AVPList is attribute value pair information.
type AVPList struct {
	// Attribute values
	StringAVPs []StringAVP `json:"stringAVP"`
}

// StringAVP presents Attribute values
type StringAVP []struct {
	AttributeValue string `json:"$"`
	//Normalised attribute name
	AttributeName string `json:"@attributeName"`
}

// FoodAndBeverageIngredientModule contains information on the constituent ingredient make up of the product.
type FoodAndBeverageIngredientModule struct {
	// Information on the constituent ingredient make up of the product specified as one string.
	IngredientStatements []IngredientStatement `json:"ingredientStatement"`
	// The fruit juice content of the trade item expressed as a percentage.
	JuiceContentPercent float64 `json:"juiceContentPercent"`
	// Information on presence or absence of additives or genetic modifications contained in the trade item.
	AdditiveInformations []AdditiveInformation `json:"additiveInformation"`
	// Information on the constituent ingredient make up of the product split out per ingredient.
	FoodAndBeverageIngredients []FoodAndBeverageIngredient `json:"foodAndBeverageIngredient"`
	// Free text field for any additional ingredient information.
	XAdditionalIngredientStatements []XAdditionalIngredientStatement `json:"x_additionalIngredientStatement"`
	// Denotes that the product in question is either a food item or a beverage.
	XIsFoodOrBeverage bool `json:"x_isFoodOrBeverage"`
}

// IngredientStatement contains information on the constituent ingredient make up of the
// product specified as one string.
type IngredientStatement struct {
	Name         string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// AdditiveInformation contains information on presence or absence of additives or genetic
// modifications contained in the trade item.
type AdditiveInformation struct {
	// The name of any additive or genetic modification contained or not contained in the trade item.
	AdditiveName string `json:"additiveName"`
	// Code indicating the level of presence of the additive. Uses code list levelOfContainmentCode
	LevelOfContainmentCode string `json:"levelOfContainmentCode"`
}

// FoodAndBeverageIngredient contains information on the constituent ingredient make up of
// the product split out per ingredient.
type FoodAndBeverageIngredient struct {
	// Value indicating the ingredient order.
	IngredientSequence string `json:"ingredientSequence"`
	// Indication of the percentage of the ingredient contained in the product.
	IngredientContentPercentage float64 `json:"ingredientContentPercentage"`
	// Text field indicating one ingredient or ingredient group (according to regulations of
	// the target market). Ingredients include any additives (colorings, preservatives, e-numbers,
	// etc) that are encompassed.
	IngredientNames []IngredientName `json:"ingredientName"`
	// Denotes that the ingredient should have it's text emphasised.
	IsIngredientEmphasised bool `json:"isIngredientEmphasised"`
	// Details on any methods and techniques used by a manufacturer or supplier to
	// the trade item, ingredients or raw materials.
	IngredientFarmingProcessing IngredientFarmingProcessing `json:"ingredientFarmingProcessing"`
	// Information on the organic nature of ingredient.
	IngredientOrganicInformation IngredientOrganicInformation `json:"ingredientOrganicInformation"`
	// Information on the activity (e.g. bottling) taken place for an ingredient as well as the associated geographic area.
	IngredientPlaceOfActivities []IngredientPlaceOfActivity `json:"ingredientPlaceOfActivity"`
}

// IngredientName is text field indicating one ingredient or ingredient group (according to regulations of the target market). Ingredients include any additives (colorings, preservatives, e-numbers, etc) that are encompassed.
type IngredientName []struct {
	Name         string `json:"$"`
	LanguageCode string `json:"@languageCode"`
	// Substring emphasis. Emphases may overlap.
	XEmphasis []IngredientXEmphasis `json:"x_emphasis"`
}

// IngredientXEmphasis is a substring emphasis. Emphases may overlap.
type IngredientXEmphasis struct {
	// Emphasis starting index in characters from the beginning of the string. Index starts at zero.
	StartAt int `json:"startAt"`
	// Emphasis length in characters
	Length int `json:"length"`
}

// IngredientFarmingProcessing details on any methods and techniques used by a manufacturer
// or supplier to the trade item, ingredients or raw materials.
type IngredientFarmingProcessing struct {
	// A statement of the presence or absence of genetically modified protein or DNA.
	// Uses code list geneticallyModifiedDeclarationCode.
	GeneticallyModifiedDeclarationCode string `json:"geneticallyModifiedDeclarationCode"`
	// Code value indicating the preservation technique used to preserve the product from
	// deterioration. Uses code list preservationTechniqueCode.
	PreservationTechniqueCode []string `json:"preservationTechniqueCode"`
}

// IngredientOrganicInformation contains information on the organic nature of ingredient.
type IngredientOrganicInformation struct {
	// Indication of the place where the agricultural raw materials of which the product is
	// composed have been farmed. It applies only to the trade item, not ingredient by ingredient.
	// Uses code list organicProductPlaceOfFarmingCode.
	OrganicProductPlaceOfFarmingCode string `json:"organicProductPlaceOfFarmingCode"`
	// Any claim to indicate the organic status of a trade item or of one or more of its components.
	OrganicClaim []IngredientOrganicClaim `json:"organicClaim"`
}

// IngredientOrganicClaim Any claim to indicate the organic status of a trade item or
// of one or more of its components.
type IngredientOrganicClaim struct {
	// A Governing body that creates and maintains standards related to organic products.
	// Uses code list organicClaimAgencyCode.
	OrganicClaimAgencyCode []string `json:"organicClaimAgencyCode"`
	// The percent of actual organic materials per weight of the trade item. This is
	// usually claimed on the product
	OrganicPercentClaim int `json:"organicPercentClaim"`
}

// IngredientPlaceOfActivity contains information on the activity (e.g. bottling)
//  taken place for an ingredient as well as the associated geographic area.
type IngredientPlaceOfActivity struct {
	// A description of the country the item may have originated from or has been processed.
	CountryOfOriginStatements []CountryOfOriginStatement `json:"countryOfOriginStatement"`
	// The place a trade item originates from. This is to be specifically used to enable things
	//  such as cities, mountain ranges, regions that do not comply with ISO standards.
	ProvenanceStatements []ProvenanceStatement `json:"provenanceStatement"`
	// The country the item may have originated from or has been processed
	CountryOfOrigins []CountryOfOrigin `json:"countryOfOrigin"`
	// Details on the activity (e.g. bottling) taken place for a trade item as well as
	// the associated geographic area.
	ProductActivityDetails []ProductActivityDetail `json:"productActivityDetails"`
}

// CountryOfOriginStatement is a description of the country the item may have originated from or has been processed.
type CountryOfOriginStatement struct {
	Value        string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// ProvenanceStatement is the place a trade item originates from. This is to be specifically
// used to enable things such as cities, mountain ranges, regions that do not comply with ISO standards.
type ProvenanceStatement struct {
	Value        string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// CountryOfOrigin is the country the item may have originated from or has been processed
type CountryOfOrigin struct {
	// Code specifying a country. Use code list countryCode.
	CountryCode string `json:"countryCode"`
}

// ProductActivityDetail contains details on the activity (e.g. bottling) taken place for a
// trade item as well as the associated geographic area.
type ProductActivityDetail struct {
	// A code depicting the type of activity being performed on a trade item. Uses code
	// list productActivityTypeCode
	ProductActivityTypeCode string `json:"productActivityTypeCode"`
	// Country where activity happens
	CountryOfActivities []CountryOfActivity `json:"countryOfActivity"`
	// An external code value that depicts a specific zone or region for example a FAO Catch Zone.
	ProductActivityRegionZoneCodeReferences []ProductActivityRegionZoneCodeReference `json:"productActivityRegionZoneCodeReference"`
	// Free text field used to describe the activity region.
	XStatements []XStatement `json:"x_statement"`
}

// CountryOfActivity contains country where activity happens
type CountryOfActivity struct {
	// Code specifying a country. Use code list countryCode
	CountryCode string `json:"countryCode"`
}

// ProductActivityRegionZoneCodeReference is an external code value that depicts a specific
//  zone or region for example a FAO Catch Zone.
type ProductActivityRegionZoneCodeReference struct {
	// The name of the agency that manages a code list.
	ExternalAgencyName string `json:"externalAgencyName"`
	// The name of the code list maintained by an external agency.
	ExternalCodeListName string `json:"externalCodeListName"`
	// The version of the code list maintained by an external agency
	ExternalCodeListVersion string `json:"externalCodeListVersion"`
	// Code list values
	EnumerationValueInformation []EnumerationValueInformation `json:"enumerationValueInformation"`
}

// EnumerationValueInformation code list values
type EnumerationValueInformation struct {
	// Code List Value maintained by an external code list agency.
	EnumerationValue string `json:"enumerationValue"`
}

// XStatement contains free text field used to describe the activity region.
type XStatement struct {
	Statement    string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// XAdditionalIngredientStatement is a free text field for any additional ingredient information.
type XAdditionalIngredientStatement []struct {
	Statement    string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// FoodAndBeveragePreparationServingModule is information on way the product can be prepared or served.
type FoodAndBeveragePreparationServingModule struct {
	// Preparation and serving information for a food and beverage item.
	PreparationServings []PreparationServing `json:"preparationServing"`
}

// PreparationServing contains preparation and serving information for a food and beverage item.
type PreparationServing struct {
	// An indication of the ease of preparation for semi-prepared products.
	// The convenience level indicates the level of preparation in percentage
	// required to prepare and helps the consumer to assess how long it will take
	// to prepare the meal.
	ConvenienceLevelPercent int `json:"convenienceLevelPercent"`
	// Textual instruction on how to prepare the product before serving.
	PreparationInstructions []PreparationInstruction `json:"preparationInstructions"`
	// A code specifying the technique used to make the product ready for consumption. Uses code list preparationTypeCode.
	PreparationTypeCode string `json:"preparationTypeCode"`
	// Free text field for serving suggestion.
	ServingSuggestions []ServingSuggestion `json:"servingSuggestion"`
	// Information on the yield of a product.
	ProductYieldInformations []ProductYieldInformation `json:"productYieldInformation"`
}

// PreparationInstruction textual instruction on how to prepare the product before serving.
type PreparationInstruction struct {
	Instruction  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// ServingSuggestion is a ree text field for serving suggestion.
type ServingSuggestion struct {
	Suggestion   string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// ProductYieldInformation is a information on the yield of a product.
type ProductYieldInformation struct {
	// Measurement
	ProductYield ProductYieldMeasurement `json:"productYield"`
	// Code indicating the type of yield measurement. Uses code list productYieldTypeCode.
	ProductYieldTypeCode string `json:"productYieldTypeCode"`
}

// ProductYieldMeasurement represents measurement
type ProductYieldMeasurement struct {
	Measurement         int    `json:"$"`
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}

// FoodAndBeveragePropertiesInformationModule contains information on
// physiochemical or other properties of food and beverage products.
type FoodAndBeveragePropertiesInformationModule struct {
	// Information on the product's physicochemical characteristics.
	PhysiochemicalCharacteristics []PhysiochemicalCharacteristic `json:"physiochemicalCharacteristic"`
}

// PhysiochemicalCharacteristic is an information on the product's
// physicochemical characteristics.
type PhysiochemicalCharacteristic struct {
	// Code indicating the type of physiochemical characteristic. Use code list
	// physiochemicalCharacteristicCode.
	PhysiochemicalCharacteristicCode string `json:"physiochemicalCharacteristicCode"`
	// Measurement value of the physicochemical characteristic.
	PhysiochemicalCharacteristicValues []PhysiochemicalCharacteristicValue `json:"physiochemicalCharacteristicValue"`
}

// PhysiochemicalCharacteristicValue is a measurement value.
type PhysiochemicalCharacteristicValue struct {
	Measurement         float64 `json:"$"`
	MeasurementUnitCode string  `json:"@measurementUnitCode"`
}

// MarketingInformationModule contains information of a trade item meant to
// convey features and benefits and targeted customer.
type MarketingInformationModule struct {
	// Information on a trade item meant to convey features and benefits.
	MarketingInformation MarketingInformation `json:"marketingInformation"`
}

// MarketingInformation contains information of a trade item meant to convey features and benefits.
type MarketingInformation struct {
	// Marketing message associated to the Trade item.
	TradeItemMarketingMessages []TradeItemMarketingMessage `json:"tradeItemMarketingMessage"`
	// Words or phrases that enables web search engines to find trade items on the internet
	// for example Shampoo, Lather, Baby.
	TradeItemKeyWords []TradeItemKeyWord `json:"tradeItemKeyWords"`
	// An indicator whether or not the Trade Item is excluded and hidden from promotions.
	// When not defined, assumed to be false.
	XHideTradeItemFromPromotions bool `json:"x_hideTradeItemFromPromotions"`
}

// TradeItemMarketingMessage contains marketing message associated to the Trade item.
type TradeItemMarketingMessage struct {
	Message      string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// TradeItemKeyWord contains words or phrases that enables web search engines
//  to find trade items on the internet for example Shampoo, Lather, Baby.
type TradeItemKeyWord struct {
	KeyWord      string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// NonfoodIngredientModule is a module providing Information on ingredients for
// items that are not food for example detergents, medicines.
type NonfoodIngredientModule struct {
	//  Ingredient statement for non-food items.
	NonfoodIngredientStatements []NonfoodIngredientStatement `json:"nonfoodIngredientStatement"`
	// Specifies a non-food ingredient of concern for a trade item as a code.
	// Uses code list nonfoodIngredientOfConcernCode.
	NonfoodIngredientOfConcernCode []string `json:"nonfoodIngredientOfConcernCode"`
	// Information on presence or absence of additives or genetic modifications
	// contained in the trade item.
	AdditiveInformations []NonFoodAdditiveInformation `json:"additiveInformation"`
	// Information on ingredients for items that are not food for example
	// detergents, medicines.
	NonfoodIngredients []NonfoodIngredient `json:"nonfoodIngredient"`
}

// NonfoodIngredientStatement is a ingredient statement for non-food items.
type NonfoodIngredientStatement struct {
	Statement    string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// NonFoodAdditiveInformation contains information on presence or absence of additives
type NonFoodAdditiveInformation struct {
	// Name of additive ingredient
	AdditiveName string `json:"additiveName"`
	// Code indicating the level of presence of the additive. Uses code list levelOfContainmentCode.
	LevelOfContainmentCode string `json:"levelOfContainmentCode"`
}

// NonfoodIngredient contains information on ingredients for items that are not food
// for example detergents, medicines.
type NonfoodIngredient struct {
	// The name of the non-food ingredient.
	IngredientName string `json:"ingredientName"`
	// Denotes the nonfood ingredient that should have it's text emphasised in
	// some fashion on the item's packaging.
	IsNonfoodIngredientEmphasized bool `json:"isNonfoodIngredientEmphasized"`
	// Substring emphasis for ingredientName.
	XEmphasis []XEmphasis `json:"x_emphasis"`
}

// NutritionalInformationModule contains information about content of nutrients.
//  Multiple sets of nutrient information can be specified with varying state,
// serving size and daily value intake base.
type NutritionalInformationModule struct {
	// Free text field for any additional nutritional claims.
	NutritionalClaims []NutritionalClaim `json:"nutritionalClaim"`
	// Details on a nutritional claim for a trade item permitted by known regulations for a target market.
	NutritionalClaimDetails []NutritionalClaimDetail `json:"nutritionalClaimDetail"`
	// Nutrient information for a trade item.
	NutrientHeaders []NutrientHeader `json:"nutrientHeader"`
}

// NutritionalClaim is a free text field for any additional nutritional claims.
type NutritionalClaim struct {
	Claim        string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// NutritionalClaimDetail contains Details on a nutritional claim for a trade
// item permitted by known regulations for a target market.
type NutritionalClaimDetail struct {
	// A code depicting the degree to which a trade item contains a specific nutrient
	//  or ingredient in relation to a health claim. Uses code list nutritionalClaimTypeCode.
	NutritionalClaimTypeCode string `json:"nutritionalClaimTypeCode"`
	// The type of nutrient, ingredient, vitamins and minerals that the nutritional claim is
	//  in reference to for example fat, copper, milk. Uses code list nutritionalClaimNutrientElementCode.
	NutritionalClaimNutrientElementCode string `json:"nutritionalClaimNutrientElementCode"`
}

// NutrientHeader contains nutrient  information for a trade item.
type NutrientHeader struct {
	// Code specifying the preparation state or type the nutrient information
	//  applies to, for example, unprepared, boiled, fried. Uses code
	//  list preparationStateCode.
	PreparationStateCode string `json:"preparationStateCode"`
	// Free text field specifying the daily value intake base for on which
	// the daily value intake per nutrient has been based.
	DailyValueIntakeReferences DailyValueIntakeReference `json:"dailyValueIntakeReference"`
	// Unit of measure code. Uses code list measurementUnitCode.
	NutrientBasisQuantity NutrientBasisQuantity `json:"nutrientBasisQuantity"`
	// Measurement value specifying the serving size in which the information
	//  per nutrient has been stated.
	ServingSizes []ServingSize `json:"servingSize"`
	// A free text field specifying the serving size for which the nutrient information has been stated.
	ServingSizeDescriptions []ServingSizeDescription `json:"servingSizeDescription"`
	// Nutrient detail for a trade item.
	NutrientDetails []NutrientDetail `json:"nutrientDetail"`
}

// DailyValueIntakeReference is a free text field specifying the daily value intake base
// for on which the daily value intake per nutrient has been based.
type DailyValueIntakeReference struct {
	Value        string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// NutrientBasisQuantity is a unit of measure code. Uses code list measurementUnitCode.
type NutrientBasisQuantity struct {
	Measurement         int    `json:"$"`
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}

// ServingSize is a measurement value specifying the serving size in which the
//  information per nutrient has been stated.
type ServingSize struct {
	Measurement         int    `json:"$"`
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}

// ServingSizeDescription is a free text field specifying the serving size
// for which the nutrient information has been stated.
type ServingSizeDescription struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// NutrientDetail describes nutrient detail for a trade item.
type NutrientDetail struct {
	// Nutrient type code. Uses code list nutrientTypeCode.
	NutrientTypeCode string `json:"nutrientTypeCode"`
	// The percentage of the recommended daily intake of a nutrient as
	// recommended by authorities of the target market. Is expressed relative
	//  to the serving size and base daily value intake.
	DailyValueIntakePercent float64 `json:"dailyValueIntakePercent"`
	// Code indicating whether the specified nutrient content is exact or
	// approximate. One should follow local regulatory guidelines when
	// selecting a precision. Uses code list measurementPrecisionCode.
	MeasurementPrecisionCode string `json:"measurementPrecisionCode"`
	// Measurement value indicating the amount of nutrient contained
	//  in the product. Is expressed relative to the serving size.
	QuantityContaineds []QuantityContained `json:"quantityContained"`
}

// QuantityContained is a measurement value indicating the amount of nutrient
// contained in the product. Is expressed relative to the serving size.
type QuantityContained struct {
	Measurement         int    `json:"$"`
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}

// PackagingInformationModule  contains packaging information for a trade item.
type PackagingInformationModule struct {
	// Details on packaging for a trade item.
	Packagings []Packaging `json:"packaging"`
}

// Packaging details for a trade item.
type Packaging struct {
	// The process the packaging could undertake for recyclable & sustainability programs.
	PackagingRecyclingProcessTypeCode []string `json:"packagingRecyclingProcessTypeCode"`
	// The dominant means used to transport, store, handle or display the trade
	// item as defined by the data source. This packaging is not used to describe
	//  any manufacturing process.Uses code list packagingTypeCode.
	PackagingTypeCode string `json:"packagingTypeCode"`
	// Details on packaging material for a trade item's packaging.
	PackagingMaterials []PackagingMaterial `json:"packagingMaterial"`
}

// PackagingMaterial is details on packaging material for a trade item's packaging.
type PackagingMaterial struct {
	// The materials used for the packaging of the trade item. Uses code list packagingMaterialTypeCode.
	PackagingMaterialTypeCode string `json:"packagingMaterialTypeCode"`
	// Determines whether packaging material is recoverable. Recoverable materials are those which
	// are capable of beingreused or returned to use in the form of raw materials.
	IsPackagingMaterialRecoverable bool `json:"isPackagingMaterialRecoverable"`
}

// PackagingMarkingModule is a module containing details on markings on the
// packaging of the trade item for example dates, environment.
type PackagingMarkingModule struct {
	// Details on markings on the packaging of the trade item.
	PackagingMarking PackagingMarking `json:"packagingMarking"`
}

// PackagingMarking is details on markings on the packaging of the trade item.
type PackagingMarking struct {
	// A marking that the trade item received recognition, endorsement, certification by
	// following guidelines by the label issuing agency. Uses code list
	// packagingMarkedLabelAccreditationCode.
	PackagingMarkedLabelAccreditationCode []string `json:"packagingMarkedLabelAccreditationCode"`
}

// PlaceOfItemActivityModule contains information on the activity (e.g. bottling)
// taken place for a trade item as well as the associated geographic area.
type PlaceOfItemActivityModule struct {
	// Information on the activity (e.g. bottling) taken place for a trade
	// item as well as the associated geographic area.
	PlaceOfProductActivity PlaceOfProductActivity `json:"placeOfProductActivity"`
}

// PlaceOfProductActivity contains information on the activity (e.g. bottling)
//  taken place for a trade item as well as the associated geographic area.
type PlaceOfProductActivity struct {
	// A description of the country the item may have originated from or has been processed.
	CountryOfOriginStatements []CountryOfOriginStatement `json:"countryOfOriginStatement"`
	// The place a trade item originates from. This is to be specifically used to enable
	// things such as cities, mountain ranges, regions that do not comply with ISO standards.
	ProvenanceStatements []ProvenanceStatement `json:"provenanceStatement"`
	// The country the item may have originated from or has been processed.
	CountryOfOrigins []CountryOfOrigin `json:"countryOfOrigin"`
	// Details on the activity (e.g. bottling) taken place for a trade item
	// as well as the associated geographic area.
	ProductActivityDetails []ProductActivityDetail `json:"productActivityDetails"`
}

// ProductCharacteristicsModule is a module used to express characteristics
// for a product for example values for a property such as numberOfPlys.
type ProductCharacteristicsModule struct {
	// A characteristic for a product for example values for a property such as
	// numberOfPlys along with its associated value.
	ProductCharacteristics []ProductCharacteristic `json:"productCharacteristics"`
}

// ProductCharacteristic describes characteristic for a product for example
//  values for a property such as numberOfPlys along with its associated value.
type ProductCharacteristic struct {
	// The name of the product characteristic being described.Uses code list productCharacteristicCode.
	ProductCharacteristicCode string `json:"productCharacteristicCode"`
	// The product characteristic value expressed as a description (text with language).
	ProductCharacteristicValueDescriptions []ProductCharacteristicValueDescription `json:"productCharacteristicValueDescription"`
	// The product characteristic value expressed as a string (text value with no language).
	ProductCharacteristicValueString []string `json:"productCharacteristicValueString"`
}

// ProductCharacteristicValueDescription expresses a product characteristic as a
// description (text with language).
type ProductCharacteristicValueDescription []struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// SafetyDataSheetModule is a module containing information usually contained
//  on a safety data sheet or on a material safety data sheet as it is referred
//  to in some target markets.
type SafetyDataSheetModule struct {
	// Trade item information usually contained on a safety data sheet
	//  or on a material safety data sheet as it is referred to in some
	//  target markets.
	SafetyDataSheetInformations []SafetyDataSheetInformation `json:"safetyDataSheetInformation"`
}

// SafetyDataSheetInformation contains trade item information usually contained on a safety
//  data sheet or on a material safety data sheet as it is referred to in some target markets.
type SafetyDataSheetInformation struct {
	// An indicator whether the Trade Item is regulated for shipment by any agency.
	IsRegulatedForTransportation bool `json:"isRegulatedForTransportation"`
	// Details related to the Globally Harmonized System of Classification and Labelling of Chemicals.
	GHSDetail GHSDetail `json:"gHSDetail"`
	// Information on Physical or Chemical Properties for a trade item for example water solubility.
	PhysicalChemicalPropertyInformation PhysicalChemicalPropertyInformation `json:"physicalChemicalPropertyInformation"`
}

// GHSDetail Details related to the Globally Harmonized System of Classification and Labelling of Chemicals.
type GHSDetail struct {
	// Words such as "Danger" or "Warning" used to emphasize hazards and indicate
	//  the relative level of severity of the hazard. For GHS these are assigned to
	//  a GHS hazard class and category. Some lower level hazard categories do not use
	//  signal words. Uses code list gHSSignalWordsCode.
	GHSSignalWordsCode string `json:"gHSSignalWordsCode"`
	// A code depicting the symbols which convey health, physical and environmental
	// hazard information, assigned to a hazard class and category for example GHS.
	// Pictograms include the harmonized hazard symbols plus other graphic elements,
	// such as borders, background patterns or colours that are intended to convey
	//  specific information. Examples of all the pictograms and downloadable files
	//  for GHS can be accessed on the UN website for the GHS. Uses code list
	//  gHSSymbolDescriptionCode.
	GHSSymbolDescriptionCode []string `json:"gHSSymbolDescriptionCode"`
	// Standard phrases describing the nature of a hazard per GHS.
	HazardStatements []HazardStatement `json:"hazardStatement"`
	// Measures listed on a hazardous label to minimize or prevent adverse
	// effects related to GHS.
	PrecautionaryStatements []PrecautionaryStatement `json:"precautionaryStatement"`
}

// HazardStatement contains standard phrases describing the nature of a hazard per GHS.
type HazardStatement struct {
	// Standard phrases assigned to a hazard class and category that describe the
	// nature of the hazard.
	HazardStatementsCode string `json:"hazardStatementsCode"`
	// A description of standard phrases assigned to a hazard class and category
	// that describe the nature of the hazard.
	HazardStatementsDescriptions []HazardStatementsDescription `json:"hazardStatementsDescription"`
}

// HazardStatementsDescription is a description of standard phrases assigned to
// a hazard class and category that describe the nature of the hazard.
type HazardStatementsDescription struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// PrecautionaryStatement contains measures listed on a hazardous label to minimize
//  or prevent adverse effects related to GHS.
type PrecautionaryStatement struct {
	// Measures listed on a hazardous label to minimize or prevent adverse effects.
	//  For GHS, the precautionary statements have been linked to each GHS hazard
	//  statement and type of hazard. Precautionary statements for GHS cover prevention,
	//  response in cases of accidental spillage or exposure, storage, and disposal.
	PrecautionaryStatementsCode string `json:"precautionaryStatementsCode"`
	// A description of the measures listed on a hazardous label to minimize or
	//  prevent adverse effects.
	PrecautionaryStatementsDescriptions []PrecautionaryStatementsDescription `json:"precautionaryStatementsDescription"`
}

// PrecautionaryStatementsDescription is a description of the measures listed on
//  a hazardous label to minimize or prevent adverse effects.
type PrecautionaryStatementsDescription struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// PhysicalChemicalPropertyInformation contains information on Physical or
//  Chemical Properties for a trade item for example water solubility.
type PhysicalChemicalPropertyInformation struct {
	// Details on a flash point for a trade item.
	FlashPoints []FlashPoint `json:"flashPoint"`
	// PH is defined as the acidity or alkalinity of an aqueous solution.
	// It is defined as the logarithm of the reciprocal of the hydrogenion
	//  concentration of a solution. pH= log10 1/[H+].
	PHInformation PHInformation `json:"pHInformation"`
}

// FlashPoint contains details on a flash point for a trade item.
type FlashPoint struct {
	// The temperature at which a substance gives off a sufficient
	//  vapour to support combustion. This uses a measurement consisting
	//  of a unit of measure and value. With the above request it requires
	//  the flash point not to be the lowest but the point at which flash point
	//  occurs and it could be that temperature and lower for some products. The
	//  scientific Measurement Precision code would determine that.
	FlashPointTemperatures []FlashPointTemperature `json:"flashPointTemperature"`
}

// FlashPointTemperature the temperature at which a substance gives off a sufficient
//  vapour to support combustion. This uses a measurement consisting of a unit of
//  measure and value. With the above request it requires the flash point not to be
//  the lowest but the point at which flash point occurs and it could be that temperature
//  and lower for some products. The scientific Measurement Precision code would determine that.
type FlashPointTemperature struct {
	Temperature                    int    `json:"$"`
	TemperatureMeasurementUnitCode string `json:"@temperatureMeasurementUnitCode"`
}

// PHInformation describes a PH value
// PH is defined as the acidity or alkalinity of an aqueous solution. It is defined
// as the logarithm of the reciprocal of the hydrogenion concentration of a solution.
// pH= log10 1/[H+].
type PHInformation struct {
	// The exact PH amount for a chemical ingredient (not a range).
	ExactPH int `json:"exactPH"`
	// The maximum range for PH.
	MaximumPH float64 `json:"maximumPH"`
	// The minimum range value for PH.
	MinimumPH float64 `json:"minimumPH"`
}

// SalesInformationModule describes sales information regarding price and selling
// conditions/restrictions of the Trade Item to the consumer.
type SalesInformationModule struct {
	// Restrictions or requirements on the retailer for sales of the Trade Item
	// to the consumer.
	SalesInformation SalesInformation `json:"salesInformation"`
}

// SalesInformation describes restrictions or requirements on the retailer for
// sales of the Trade Item to the consumer.
type SalesInformation struct {
	// A code depicting restrictions imposed on the Trade Item regarding how
	// it can be sold to the consumer for example Prescription Required. Uses
	// code list consumerSalesConditionCode.
	ConsumerSalesConditionCode []string `json:"consumerSalesConditionCode"`
	// Indicator to show how a product is sold. Uses code list priceByMeasureTypeCode.
	PriceByMeasureTypeCode string `json:"priceByMeasureTypeCode"`
	// The quantity of the product at usage. Applicable for concentrated products
	// and products where the comparison price is calculated based on a measurement
	// other than netContent.
	PriceComparisonMeasurements []PriceComparisonMeasurement `json:"priceComparisonMeasurement"`
	// Describes the measurement used for selling unit of the Trade Item to the end consumer.
	SellingUnitOfMeasure string `json:"sellingUnitOfMeasure"`
	// Defines compliancy with EU 1169 regulation.
	XEu1169Compliance XEu1169Compliance `json:"x_eu1169Compliance"`
	// An indicator whether or not the Trade Item is excluded from loyalty programs.
	XIsExcludedFromLoyaltyPrograms bool `json:"x_isExcludedFromLoyaltyPrograms"`
	// Defines how much the quantity of a Trade Item is changed when additional items
	// are added or removed from shopping basket.
	XSellingContentIncrement int `json:"x_sellingContentIncrement"`
	// Defines the initial quantity of Trade Item when the first instance of the item
	// is added to shopping basket.
	XSellingContentInitial int `json:"x_sellingContentInitial"`
	// Defines the measurement unit code used for selling of the Trade Item to the end
	// consumer. Uses code list sellingUnitOfMeasure.
	XSellingUnitOfMeasureCode string `json:"x_sellingUnitOfMeasureCode"`
}

// PriceComparisonMeasurement is the quantity of the product at usage. Applicable
// for concentrated products and products where the comparison price is calculated
// based on a measurement other than netContent.
type PriceComparisonMeasurement struct {
	Measurement         float64 `json:"$"`
	MeasurementUnitCode string  `json:"@measurementUnitCode"`
}

// XEu1169Compliance defines compliancy with EU 1169 regulation.
type XEu1169Compliance struct {
	// Regulation compliancy. Uses code list x_complianceCode.
	XComplianceCode string `json:"x_complianceCode"`
}

// TradeItemDescriptionModule a module carrying general descriptions of the trade item including
// brand, form, variant.
type TradeItemDescriptionModule struct {
	// Description Information for the trade item.
	TradeItemDescriptionInformation TradeItemDescriptionInformation `json:"tradeItemDescriptionInformation"`
}

// TradeItemDescriptionInformation is description information for the trade item.
type TradeItemDescriptionInformation struct {
	//Additional variants necessary to communicate to the industry to
	//  help define the product.
	AdditionalTradeItemDescriptions []AdditionalTradeItemDescription `json:"additionalTradeItemDescription"`
	// A free form short length description of the trade item that can
	// be used to identify the trade item at point of sale.
	DescriptionShorts []DescriptionShort `json:"descriptionShort"`
	// Describes use of the product or service by the consumer. Should help
	// clarify the product classification associated with the GTIN.
	FunctionalNames []FunctionalName `json:"functionalName"`
	// An understandable and useable description of a trade item using brand
	// and other descriptors. This attribute is filled with as little abbreviation
	//  as possible while keeping to a reasonable length. This should be a meaningful
	//  description of the trade item with full spelling to facilitate message processing.
	//  Retailers can use this description as the base to fully understand the brand,
	//  flavour, scent etc. of the specific GTIN in order to accurately create a product
	//  description as needed for their internal systems.
	TradeItemDescriptions []TradeItemDescription `json:"tradeItemDescription"`
	// Free text field used to identify the variant of the product. Variants are
	//  the distinguishing characteristics that differentiate products with the
	//  same brand and size including such things as the particular flavor, fragrance, taste.
	VariantDescriptions []VariantDescription `json:"variantDescription"`
	// Information on brands and sub-brands for a trade item.
	BrandNameInformation BrandNameInformation `json:"brandNameInformation"`
}

// AdditionalTradeItemDescription contains additional variants
//  necessary to communicate to the industry to help define the product.
type AdditionalTradeItemDescription struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// DescriptionShort is a free form short length description of the trade item that can
// be used to identify the trade item at point of sale.
type DescriptionShort struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// FunctionalName describes use of the product or service by the consumer.
// Should help clarify the product classification associated with the GTIN.
type FunctionalName struct {
	Name         string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// TradeItemDescription is an understandable and useable description of
//  a trade item using brand and other descriptors. This attribute is
// filled with as little abbreviation as possible while keeping to a reasonable
//  length. This should be a meaningful description of the trade item with full
//  spelling to facilitate message processing. Retailers can use this description
//  as the base to fully understand the brand, flavour, scent etc. of the specific
//  GTIN in order to accurately create a product description as needed for their
//  internal systems.
type TradeItemDescription struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// VariantDescription free text field used to identify the variant of the product.
// Variants are the distinguishing characteristics that differentiate products with
//  the same brand and size including such things as the particular flavor, fragrance, taste.
type VariantDescription struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// BrandNameInformation contains information on brands and sub-brands for a trade item.
type BrandNameInformation struct {
	// The recognisable name used by a brand owner to uniquely identify a line of trade
	//  item or services. This is recognizable by the consumer.
	BrandName string `json:"brandName"`
	// The recognisable name used by a brand owner to uniquely identify a line of trade
	//  item or services expressed in a different language than the primary brand name (brandName).
	LanguageSpecificBrandNames []LanguageSpecificBrandName `json:"languageSpecificBrandName"`
	// A second level of brand expressed in a different language than the primary sub-brand name (subBrand).
	LanguageSpecificSubbrandNames []LanguageSpecificSubbrandName `json:"languageSpecificSubbrandName"`
	// Second level of brand. Can be a trademark. It is the primary differentiating factor
	// that a brand owner wants to communicate to the consumer or buyer.
	SubBrand string `json:"subBrand"`
}

// LanguageSpecificBrandName is the recognisable name used by a brand owner to uniquely identify
// a line of trade item or services expressed in a different language than the primary brand name (brandName).
type LanguageSpecificBrandName struct {
	Name         string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// LanguageSpecificSubbrandName is a second level of brand expressed in a different
//  language than the primary sub-brand name (subBrand).
type LanguageSpecificSubbrandName struct {
	Name         string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// TradeItemLifespanModule is a module containing information on the amount
// of time the item can or should be used, sold, etc.
type TradeItemLifespanModule struct {
	// Information on the amount of time the item can or should be used, sold, etc.
	TradeItemLifespan TradeItemLifespan `json:"tradeItemLifespan"`
}

// TradeItemLifespan contains information on the amount of time the item can or
//  should be used, sold, etc.
type TradeItemLifespan struct {
	// The period of day, guaranteed by the manufacturer, before the expiration date of the product, based on the production.
	MinimumTradeItemLifespanFromTimeOfProduction int `json:"minimumTradeItemLifespanFromTimeOfProduction"`
	// The number of days the trade item that had been opened can remain on the shelf and must then be removed.
	OpenedTradeItemLifespan int `json:"openedTradeItemLifespan"`
}

// TradeItemMeasurementsModule is a module containing measurement
// information for the trade item.
type TradeItemMeasurementsModule struct {
	// Measurement information for the trade item.
	TradeItemMeasurements TradeItemMeasurements `json:"tradeItemMeasurements"`
}

// TradeItemMeasurements is measurement information for the trade item.
type TradeItemMeasurements struct {
	Depth  GDSNDepth  `json:"depth"`
	Height GDSNHeight `json:"height"`
	Width  GDSNWidth  `json:"width"`
	//The amount of the trade item contained by a package, usually as claimed on the label. For example, Water 750ml - net content = "750 MLT" ; 20 count pack of diapers, net content = "20 ea.". In case of multi-pack, indicates the net content of the total trade item. For fixed value trade items use the value claimed on the package, to avoid variable fill rate issue that arises with some trade item which are sold by volume or weight, and whose actual content may vary slightly from batch to batch. In case of variable quantity trade items, indicates the average quantity.
	NetContent []GDSNNetContent `json:"netContent"`
	// Information on the weight of a trade item.
	TradeItemWeight TradeItemWeight `json:"tradeItemWeight"`
}

// GDSNDepth presents depth value of product.
type GDSNDepth struct {
	Value               float64 `json:"$"`
	MeasurementUnitCode string  `json:"@measurementUnitCode"`
}

// GDSNHeight presents height value of product.
type GDSNHeight struct {
	Value               float64 `json:"$"`
	MeasurementUnitCode string  `json:"@measurementUnitCode"`
}

// GDSNWidth presents width value of product.
type GDSNWidth struct {
	Value               float64 `json:"$"`
	MeasurementUnitCode string  `json:"@measurementUnitCode"`
}

// GDSNNetContent is the amount of the trade item contained by a package,
//  usually as claimed on the label. For example, Water 750ml - net
// content = "750 MLT" ; 20 count pack of diapers, net content = "20 ea.".
//  In case of multi-pack, indicates the net content of the total trade item.
//  For fixed value trade items use the value claimed on the package, to avoid
// variable fill rate issue that arises with some trade item which are sold
// by volume or weight, and whose actual content may vary slightly from batch
//  to batch. In case of variable quantity trade items, indicates the average quantity.
type GDSNNetContent struct {
	Measurement         int    `json:"$"`
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}

// TradeItemWeight is information on the weight of a trade item.
type TradeItemWeight struct {
	DrainedWeight GDSNDrainedWeight `json:"drainedWeight"`
	GrossWeight   GDSNDrainedWeight `json:"grossWeight"`
	NetWeight     GDSNNetWeight     `json:"netWeight"`
}

type GDSNDrainedWeight struct {
	Measurement         int    `json:"$"`
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}

type GDSNGrossWeight struct {
	Measurement         int    `json:"$"`
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}

type GDSNNetWeight struct {
	Measurement         int    `json:"$"`
	MeasurementUnitCode string `json:"@measurementUnitCode"`
}

// TradeItemTemperatureInformationModule is information on temperature considerations for trade item.
type TradeItemTemperatureInformationModule struct {
	// The condition of the product sold to the end consumer. Uses code
	// list tradeItemTemperatureConditionTypeCode.
	TradeItemTemperatureConditionTypeCode string `json:"tradeItemTemperatureConditionTypeCode"`
	// Details on permissible temperatures of a trade item during various points of the supply chain.
	TradeItemTemperatureInformations []TradeItemTemperatureInformation `json:"tradeItemTemperatureInformation"`
}

// TradeItemTemperatureInformation describes details on permissible temperatures of
//  a trade item during various points of the supply chain.
type TradeItemTemperatureInformation struct {
	MaximumTemperature          GDSNTemperature `json:"maximumTemperature"`
	MaximumToleranceTemperature GDSNTemperature `json:"maximumToleranceTemperature"`
	MinimumTemperature          GDSNTemperature `json:"minimumTemperature"`
	MinumumToleranceTemperature GDSNTemperature `json:"minumumToleranceTemperature"`
	// Code qualifying the type of a temperature requirement for example Storage.
	// Uses code list temperatureQualifierCode.
	TemperatureQualifierCode string `json:"temperatureQualifierCode"`
}

// GDSNTemperature provides temperature measurement value and associated unit of measure code.
type GDSNTemperature struct {
	Temperature                    int    `json:"$"`
	TemperatureMeasurementUnitCode string `json:"@temperatureMeasurementUnitCode"`
}

// VariableTradeItemInformationModule is a module with information specific to variable weight or dimension trade items.
type VariableTradeItemInformationModule struct {
	// Information specific to variable weight or dimension trade items.
	VariableTradeItemInformation VariableTradeItemInformation `json:"variableTradeItemInformation"`
}

// VariableTradeItemInformation is information specific to variable weight or dimension trade items.
type VariableTradeItemInformation struct {
	// Indicates that an article is not a fixed quantity, but that the quantity is variable. Can be weight,
	// length, volume. trade item is used or traded in continuous rather than discrete quantities.
	IsTradeItemAVariableUnit bool `json:"isTradeItemAVariableUnit"`
	// Indicator to show whether product is loose or pre-packed. Uses code list variableTradeItemTypeCode.
	VariableTradeItemTypeCode string `json:"variableTradeItemTypeCode"`
	// Indication of the percentage value that the actual weight of the trade item may differ from the average
	// or estimated weight given. For example, Roast beef off the bone 3.5 kg, Gross weight 3500 Grams,
	//  Range = 14 %. This means that this item may be produced with weight values ranging from 3.010 kg to 3.990 kg.
	VariableWeightAllowableDeviationPercentage int `json:"variableWeightAllowableDeviationPercentage"`
}

// DGCodeListModule lists associated code lists
type DGCodeListModule struct {
	CodeLists []CodeList `json:"codeList"`
}

// CodeList presents GDSN code list
type CodeList struct {
	// Code list name
	CodeListName string `json:"codeListName"`
	// Whether the code list is an external code list.
	IsExternalCodeList bool             `json:"isExternalCodeList"`
	CodeListRecords    []CodeListRecord `json:"codeListRecord"`
	// The name of the agency that manages a code list.
	ExternalAgencyName string `json:"externalAgencyName,omitempty"`
	// The name of the code list maintained by an external agency.
	ExternalCodeListName string `json:"externalCodeListName,omitempty"`
	// The version of the code list maintained by an external agency.
	ExternalCodeListVersion string `json:"externalCodeListVersion,omitempty"`
}

// CodeListRecord presents single record of code list
type CodeListRecord struct {
	Code        string                `json:"code"`
	Name        []CodeListRecordField `json:"name"`
	Description []CodeListRecordField `json:"description"`
	Label       []CodeListRecordField `json:"label"`
}

// CodeListRecordField presents code list record value
type CodeListRecordField struct {
	Value        string `json:"$"`
	LanguegeCode string `json:"@languegeCode"`
}

// DGMediaModule contains product media properties
type DGMediaModule struct {
	// Media files associated with the product.
	Media []GDSNMedia `json:"media"`
}

// GDSNMedia presents media files associated with the product.
type GDSNMedia struct {
	MediaSequence          int                     `json:"mediaSequence"`
	MediaLanguageCodes     []string                `json:"mediaLanguageCode"`
	MediaNames             []GDSNMediaName         `json:"mediaName"`
	MediaStorageKey        string                  `json:"mediaStorageKey"`
	MediaMimeType          string                  `json:"mediaMimeType"`
	MediaDimensionWidth    int                     `json:"mediaDimensionWidth"`
	MediaDimensionHeight   int                     `json:"mediaDimensionHeight"`
	MediaFileName          string                  `json:"mediaFileName"`
	MediaTypeCode          string                  `json:"mediaTypeCode"`
	MediaTypeVariantCode   string                  `json:"mediaTypeVariantCode"`
	IsReadyForPublishing   bool                    `json:"isReadyForPublishing"`
	MediaStateDescriptions []MediaStateDescription `json:"mediaStateDescription"`
	MediaProvider          MediaProvider           `json:"mediaProvider"`
}

// GDSNMediaName presents name of media
type GDSNMediaName struct {
	Name         string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// MediaStateDescription contains description of media
type MediaStateDescription struct {
	Description  string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// MediaProvider is the identification of a party, by GLN, in a specific party role.
type MediaProvider struct {
	// The Global Location Number (GLN) is a structured Identification of a physical
	// location, legal or functional entity within an enterprise. The GLN is the primary
	// party identifier. Each party identified in the trading relationship must have a
	//  primary party Identification.
	Gln string `json:"gln"`
	// The name of the party expressed in text.
	PartyName string `json:"partyName"`
	// The address associated with the party. This could be the full company address.
	PartyAddress string `json:"partyAddress"`
}

// DGPresentationModule contains product presentation properties
type DGPresentationModule struct {
	// Limits product visibility to given periods. Presentation
	// categories may further limit product visibility. Periods can be open-ended.
	ProductConsumerVisibilities []ProductConsumerVisibility `json:"productConsumerVisibility"`
	// Sets the visibility of product. Overrides visibility given by productConsumerVisibility.
	ProductAbsoluteConsumerVisibility bool `json:"productAbsoluteConsumerVisibility"`
	// Categories the product is associated with.
	PresentationCategories []PresentationCategory `json:"presentationCategory"`
}

// ProductConsumerVisibility limits product visibility to given periods. Presentation
// categories may further limit product visibility. Periods can be open-ended.
type ProductConsumerVisibility struct {
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
}

// PresentationCategory presents a category
type PresentationCategory struct {
	// Category tree name
	TreeName string `json:"treeName"`
	// Category external ID
	ExtID string `json:"extId"`
	// Restricts category association to given periods. Periods can be open-ended.
	ValidityPeriods ValidityPeriod `json:"validityPeriod"`
}

// ValidityPeriod restricts category association to given periods. Periods can be open-ended.
type ValidityPeriod struct {
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
}

// DGProductAttributeModule is a module containing freely defined product attributes.
type DGProductAttributeModule struct {
	// Product attribute groups used to collect attributes into meaningful sets.
	ProductAttributeGroups []ProductAttributeGroup `json:"productAttributeGroup"`
}

// ProductAttributeGroup describes product attribute group used to collect attributes into meaningful sets.
type ProductAttributeGroup struct {
	// Product-unique data provider assigned identifer.
	ProductAttributeGroupExtID string `json:"productAttributeGroupExtId"`
	// Value indicating the group order.
	ProductAttributeGroupSequence string `json:"productAttributeGroupSequence"`
	// Attribute group name used for presentation.
	ProductAttributeGroupNames []ProductAttributeGroupName `json:"productAttributeGroupName"`
	// Product attributes.
	ProductAttributes []ProductAttribute `json:"productAttribute"`
}

// ProductAttributeGroupName is attribute group name used for presentation.
type ProductAttributeGroupName struct {
	Name         string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

// ProductAttribute describes attribute declaration. Note that while neither
//  productAttributeValueString, productAttributeValueNumeric, nor
//  productAttributeValueBoolean is required, exactly one of these must be provided.
type ProductAttribute struct {
	// Group-unique data provider assigned identifier.
	ProductAttributeExtID string `json:"productAttributeExtId"`
	// Value indicating the attribute order.
	ProductAttributeSequence string `json:"productAttributeSequence"`
	// Code specifying the attribute type. Uses code list productAttributeTypeCode.
	ProductAttributeTypeCode string `json:"productAttributeTypeCode"`
	// An indicator whether or not the attribute is and can be used as a facet attribute.
	IsFacetAttribute             bool                          `json:"isFacetAttribute"`
	ProductAttributeNames        []ProductAttributeName        `json:"productAttributeName"`
	ProductAttributeValueStrings []ProductAttributeValueString `json:"productAttributeValueString"`
	ProductAttributeValueNumeric float64                       `json:"productAttributeValueNumeric"`
	ProductAttributeValueBoolean bool                          `json:"productAttributeValueBoolean"`
}

// ProductAttributeName presents attribute name
type ProductAttributeName struct {
	Name         string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}

type ProductAttributeValueString struct {
	Value        string `json:"$"`
	LanguageCode string `json:"@languageCode"`
}
