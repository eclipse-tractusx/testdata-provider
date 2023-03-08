package model

type AssetInformation struct {
	Identification      string               `json:"identification"`
	IdShort             string               `json:"idShort"`
	Description         []Description        `json:"description"`
	GlobalAssetId       GlobalAssetId        `json:"globalAssetId"`
	SpecificAssetIds    []SpecificAssetId    `json:"specificAssetIds"`
	SubmodelDescriptors []SubmodelDescriptor `json:"submodelDescriptors"`
}

// Description Remark Dominik
//
// From our example I cannot see what type the description is, and currently I'm too lazy to look it up :-)
type Description struct {
}

type GlobalAssetId struct {
	Value []string `json:"value"`
}

type SpecificAssetId struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SubmodelDescriptor struct {
	Description    []Description `json:"description"`
	IdShort        string        `json:"idShort"`
	Identification string        `json:"identification"`
	SemanticId     SemanticId    `json:"semanticId"`
	Endpoints      []Endpoint    `json:"endpoints"`
}

type SemanticId struct {
	Value []string `json:"value"`
}

type Endpoint struct {
	Interface           string              `json:"interface"`
	ProtocolInformation ProtocolInformation `json:"protocolInformation"`
}

// ProtocolInformation Remark Dominik
//
// Yes, why always string? We should do this better :-)
type ProtocolInformation struct {
	EndpointAddress         string `json:"endpointAddress"`
	EndpointProtocol        string `json:"endpointProtocol"`
	EndpointProtocolVersion string `json:"endpointProtocolVersion"`
}
