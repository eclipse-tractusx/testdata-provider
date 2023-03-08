package edc

import (
	"bytes"
	"cx/internal/edc/model"
	json2 "encoding/json"
	"io"
	"net/http"
	"os"
	"path"
)

type Api struct {
	url     string
	authKey string
}

type CreateAssetWithAddress struct {
	Asset       model.Asset       `json:"asset"`
	DataAddress model.DataAddress `json:"dataAddress"`
}

func NewApi() Api {
	api := Api{
		url:     os.Getenv("EDC_URL"),
		authKey: os.Getenv("EDC_AUTH_KEY"),
	}
	return api
}

func (api Api) CreateAsset(assetId string, assetUrl string) {
	asset := model.NewAsset(assetId)
	dataAddress := model.NewProxyAddress(assetUrl)
	assetWithAddress := CreateAssetWithAddress{
		Asset:       asset,
		DataAddress: dataAddress,
	}
	payload, err := json2.Marshal(assetWithAddress)
	if err != nil {
		panic(err)
	}

	endpoint := path.Join(api.url, "data", "assets")

	api.sendPost(endpoint, payload)
}

func (api Api) createPolicy(id string) {
	policy := model.NewAllowAllPolicy(id)
	payload, err := json2.Marshal(policy)
	if err != nil {
		panic(err)
	}

	endpoint := path.Join(api.url, "data", "policydefinitions")

	api.sendPost(endpoint, payload)
}

func (api Api) createContractDefinition(id string, assetId string, accessPolicyId string, contractPolicyId string) {
	contract := model.NewContractDefinition(id, assetId, accessPolicyId, contractPolicyId)
	payload, err := json2.Marshal(contract)
	if err != nil {
		panic(err)
	}

	endpoint := path.Join(api.url, "data", "contractdefinitions")

	api.sendPost(endpoint, payload)
}

func (api Api) sendPost(endpoint string, json []byte) {
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Api-Key", api.authKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		panic("Failed to create asset")
	}
}
