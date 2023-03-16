/********************************************************************************
 * Copyright (c) 2023 Contributors to the Eclipse Foundation
 *
 * See the NOTICE file(s) distributed with this work for additional
 * information regarding copyright ownership.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Apache License, Version 2.0 which is available at
 * https://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 ********************************************************************************/

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
