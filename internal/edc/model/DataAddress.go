package model

type DataAddress struct {
	Properties map[string]string `json:"properties"`
}

func NewProxyAddress(url string) DataAddress {
	address := DataAddress{
		Properties: map[string]string{},
	}
	address.Properties["type"] = "HttpData"
	address.Properties["baseUrl"] = url
	return address
}
