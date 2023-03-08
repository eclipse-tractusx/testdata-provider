package model

type PolicyDefinition struct {
	Id     string `json:"id"`
	Policy Policy `json:"policy"`
}

type Policy struct {
	Prohibitions []Prohibition `json:"prohibitions"`
	Obligations  []Obligation  `json:"obligations"`
	Permissions  []Permission  `json:"permissions"`
}

type Prohibition struct {
}

type Obligation struct {
}

type Permission struct {
	EdcType     string       `json:"edcType"`
	Actions     Action       `json:"action"`
	Constraints []Constraint `json:"constraints"`
}

type Constraint struct {
}

type Action struct {
	Type string `json:"type"`
}

func NewAllowAllPolicy(id string) PolicyDefinition {
	policy := PolicyDefinition{
		Id: id,
		Policy: Policy{
			Prohibitions: []Prohibition{},
			Obligations:  []Obligation{},
			Permissions: []Permission{
				{
					EdcType: "dataspaceconnector:permission",
					Actions: Action{
						Type: "USE",
					},
				},
			},
		},
	}
	return policy
}
