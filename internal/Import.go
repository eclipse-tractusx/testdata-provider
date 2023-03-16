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

package internal

import (
	"cx/internal/edc"
	psql "cx/internal/psql"
	"cx/internal/psql/model"
	testdata2 "cx/internal/testdata"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

func ImportTestDataByBPNL(testdata []byte, bpnl string) error {

	id, err := databaseImport(testdata, bpnl)
	if err != nil {
		return err
	}

	err = edcImport(id)
	if err != nil {
		return err
	}

	err = aasImport(id)
	if err != nil {
		return err
	}

	return nil
}

func aasImport(importId string) error {

	// TODO Write After AAS Import

	return nil
}

func edcImport(importId string) error {

	hostName := os.Getenv("HOST_NAME")
	db := psql.NewInstance()
	testDataImport, err := db.LoadImport(importId)
	if err != nil {
		return err
	}

	for _, entry := range testDataImport.Entries {

		dataEndpoint := "http://" + hostName + ":8080/data/" + entry.ID

		edcApi := edc.NewApi()
		edcApi.CreateAsset(entry.EdcAssetId, dataEndpoint)
		entry.EdcAssetCreated = true
		err := db.Save(&testDataImport)
		if err != nil {
			return err
		}

		edcApi.CreatePolicy(entry.EdcPolicyId)
		entry.EdcPolicyCreated = true
		err = db.Save(&testDataImport)
		if err != nil {
			return err
		}

		edcApi.CreateContractDefinition(entry.EdcContractDefinitionId, entry.EdcAssetId, entry.EdcPolicyId, entry.EdcPolicyId)
		entry.EdcContractDefinitionCreated = true
		err = db.Save(&testDataImport)
		if err != nil {
			return err
		}
	}

	return nil
}

func databaseImport(testdata []byte, bpnl string) (string, error) {
	testData, _ := testdata2.Parse(testdata)
	testDataContainer := testData.HttpsCatenaxIoSchemaTestDataContainer100

	testDataImport := model.TestDataImport{
		Entries: []model.TestDataImportEntry{},
	}

	for _, assetInformation := range testDataContainer {

		// filter for target BPNL
		if bpnl != assetInformation.Bpnl {
			continue
		}

		partsAsPlanned := assetInformation.UrnBammIoCatenaxPartAsPlanned100PartAsPlanned
		partSiteInformation := assetInformation.UrnBammIoCatenaxPartSiteInformationAsPlanned100PartSiteInformationAsPlanned
		singleLevelBomAsPlanned := assetInformation.UrnBammIoCatenaxSingleLevelBomAsPlanned102SingleLevelBomAsPlanned
		importId := time.Now().Format(time.RFC3339) + ":" + assetInformation.CatenaXId

		for index, part := range partsAsPlanned {
			data, _ := json.Marshal(part)
			id := importId + ":PartsAsPlanned:" + strconv.Itoa(index)
			entry := createEntry(id, data)
			testDataImport.Entries = append(testDataImport.Entries, entry)
		}
		for index, part := range partSiteInformation {
			data, _ := json.Marshal(part)
			id := importId + ":PartSiteInformation:" + strconv.Itoa(index)
			entry := createEntry(id, data)
			testDataImport.Entries = append(testDataImport.Entries, entry)
		}
		for index, part := range singleLevelBomAsPlanned {
			data, _ := json.Marshal(part)
			id := importId + ":SingleLevelBomAsPlanned:" + strconv.Itoa(index)
			entry := createEntry(id, data)
			testDataImport.Entries = append(testDataImport.Entries, entry)
		}
	}

	err := psql.NewInstance().Save(&testDataImport)
	return testDataImport.ID, err
}

func createEntry(id string, data []byte) model.TestDataImportEntry {
	assetId := id + ":asset"
	shellDescriptorId := id + ":shellDescriptor"
	policyId := id + ":policy"
	contractDefinitionId := id + ":contractDefinition"

	entry := model.TestDataImportEntry{
		EdcAssetId:                   assetId,
		EdcAssetCreated:              false,
		EdcPolicyId:                  policyId,
		EdcPolicyCreated:             false,
		EdcContractDefinitionId:      contractDefinitionId,
		EdcContractDefinitionCreated: false,
		ShellDescriptorId:            shellDescriptorId,
		ShellDescriptorIdCreated:     false,
		Data:                         data,
	}
	return entry
}
