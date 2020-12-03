package masterproduct

// DGCodeListModule lists associated code lists
type DGCodeListModule struct {
	CodeLists *[]CodeList `json:"codeList,omitempty"`
}

// CodeList presents GDSN code list
type CodeList struct {
	// Code list name.
	CodeListName string `json:"codeListName"`
	// Whether the code list is an external code list.
	IsExternalCodeList *bool             `json:"isExternalCodeList,omitempty"`
	CodeListRecords    *[]CodeListRecord `json:"codeListRecord,omitempty"`
	// The name of the agency that manages a code list.
	ExternalAgencyName *string `json:"externalAgencyName,omitempty"`
	// The name of the code list maintained by an external agency.
	ExternalCodeListName *string `json:"externalCodeListName,omitempty"`
	// The version of the code list maintained by an external agency.
	ExternalCodeListVersion *string `json:"externalCodeListVersion,omitempty"`
}

// CodeListRecord presents a single record of code list.
type CodeListRecord struct {
	Code        string      `json:"code"`
	Name        *[]MLString `json:"name,omitempty"`
	Description *[]MLString `json:"description,omitempty"`
	Label       *[]MLString `json:"label,omitempty"`
}
