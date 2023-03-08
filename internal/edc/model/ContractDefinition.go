package model

type ContractDefinition struct {
	Id               string      `json:"id"`
	Criteria         []Criterion `json:"criteria"`
	AccessPolicyId   string      `json:"accessPolicyId"`
	ContractPolicyId string      `json:"contractPolicyId"`
}

type Criterion struct {
	OperandLeft  string `json:"operandLeft"`
	Operator     string `json:"operator"`
	OperandRight string `json:"operandRight"`
}

func NewContractDefinition(id string, assetId string, accessPolicyId string, contractPolicyId string) ContractDefinition {
	contract := ContractDefinition{
		Id: id,
		Criteria: []Criterion{
			{
				OperandLeft:  "asset:prop:id",
				Operator:     "=",
				OperandRight: assetId,
			},
		},
		AccessPolicyId:   accessPolicyId,
		ContractPolicyId: contractPolicyId,
	}
	return contract
}
