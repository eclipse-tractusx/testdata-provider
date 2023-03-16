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
