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
