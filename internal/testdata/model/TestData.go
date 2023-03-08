package model

type TestData struct {
	TestDataContainer []TestDataContainer100 `json:"https://catenax.io/schema/TestDataContainer/1.0.0"`
	VehicleBlueprint  []VehicleBlueprint100  `json:"https://catenax.io/schema/VehicleBlueprint/1.0.0"`
}

// TODO Comment Dominik: Complete Model

type TestDataContainer100 struct {
}

type VehicleBlueprint100 struct {
}

type CatenaXObject struct {
	CatenaXId                        string `json:"catenaXId"`
	BusinessPartnerNumberLegalEntity string `json:"bpnl"`
}
