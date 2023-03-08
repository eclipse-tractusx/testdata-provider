package model

import "time"

type Asset struct {
	Properties map[string]string `json:"properties"`
}

func NewAsset(id string) Asset {
	asset := Asset{
		Properties: map[string]string{},
	}
	asset.Properties["asset:prop:id"] = id
	asset.Properties["created"] = time.Now().String()
	asset.Properties["createdBy"] = "testdata-provider"
	return asset
}
