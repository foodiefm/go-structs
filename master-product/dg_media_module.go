package masterproduct

// DGMediaModule contains product media properties.
type DGMediaModule struct {
	// Media files associated with the product.
	Media *[]DGMedia `json:"media,omitempty"`
}

// DGMedia presents media files associated with the product.
type DGMedia struct {
	MediaSequence      *int      `json:"mediaSequence,omitempty"`
	MediaLanguageCodes *[]string `json:"mediaLanguageCode,omitempty"`
	// MediaNames present the name of media.
	MediaNames           *[]MLString `json:"mediaName,omitempty"`
	MediaStorageKey      string      `json:"mediaStorageKey"`
	MediaMimeType        string      `json:"mediaMimeType"`
	MediaDimensionWidth  *int        `json:"mediaDimensionWidth,omitempty"`
	MediaDimensionHeight *int        `json:"mediaDimensionHeight,omitempty"`
	MediaFileName        *string     `json:"mediaFileName,omitempty"`
	MediaTypeCode        string      `json:"mediaTypeCode"`
	MediaTypeVariantCode *string     `json:"mediaTypeVariantCode,omitempty"`
	IsReadyForPublishing *bool       `json:"isReadyForPublishing,omitempty"`
	// MediaStateDescriptions contain description of the media.
	MediaStateDescriptions *[]MLString `json:"mediaStateDescription,omitempty"`
	// MediaProvider is the identification of a party, by GLN, in a
	// specific party role.
	MediaProvider *PartyInRole `json:"mediaProvider,omitempty"`
}
