package masterproduct

import (
	"time"
)

// DGPresentationModule contains product presentation properties.
type DGPresentationModule struct {
	// Sets the visibility of product. Overrides visibility given by
	// productConsumerVisibility.
	ProductAbsoluteConsumerVisibility *bool `json:"productAbsoluteConsumerVisibility,omitempty"`
	// Limits product visibility to given periods. Presentation categories
	// may further limit product visibility. Periods can be open-ended.
	ProductConsumerVisibilities *[]TimePeriod `json:"productConsumerVisibility,omitempty"`
	// Categories the product is associated with.
	PresentationCategories *[]PresentationCategory `json:"presentationCategory,omitempty"`
}

// PresentationCategory presents a category
type PresentationCategory struct {
	// Category tree name.
	TreeName string `json:"treeName"`
	// Category external ID.
	ExtID string `json:"extId"`
	// Restricts category association to given periods. Periods can be
	// open-ended.
	ValidityPeriods *[]TimePeriod `json:"validityPeriod,omitempty"`
}

// TimePeriod defines a time range. A period can be open ended.
type TimePeriod struct {
	StartDateTime *time.Time `json:"startDateTime.omitempty"`
	EndDateTime   *time.Time `json:"endDateTime,omitempty"`
}
