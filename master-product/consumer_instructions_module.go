package masterproduct

// ConsumerInstructionsModule is a module contain instructions on how the
// consumer is to use or store a trade item.
type ConsumerInstructionsModule struct {
	// Instructions on how the consumer is to use or store a trade item.
	ConsumerInstructions *ConsumerInstructions `json:"consumerInstructions.omitempty"`
}

// ConsumerInstructions contains instructions on how the consumer is to use or
// store a trade item.
type ConsumerInstructions struct {
	// Expresses in text the consumer storage instructions of a product
	// which are normally held on the label or accompanying the product.
	// This information may or may not be labeled on the pack. Instructions
	// may refer to a suggested storage temperature, a specific storage
	// requirement.
	ConsumerStorageInstructions *[]MLString `json:"consumerStorageInstructions.omitempty"`
	// Expresses in text the consumer usage instructions of a product which
	// are normally held on the label or accompanying the product. This
	// information may or may not be labeled on the pack. Instructions may
	// refer to a the how the consumer is to use the product, This does not
	// include storage, food preparations, and drug dosage and preparation
	// instructions.
	ConsumerUsageInstructions *[]MLString `json:"consumerUsageInstructions.omitempty"`
}
