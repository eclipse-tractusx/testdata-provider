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

import "time"

type TestData struct {
	HttpsCatenaxIoSchemaTestDataContainer100 []struct {
		CatenaXId   string `json:"catenaXId"`
		Bpnl        string `json:"bpnl"`
		PlainObject []struct {
			BPNOEMC          string    `json:"BPN_OEM_C"`
			BPNOEMA          string    `json:"BPN_OEM_A"`
			BPNOEMB          string    `json:"BPN_OEM_B"`
			BPNIRSTEST       string    `json:"BPN_IRS_TEST"`
			BPNNTIERA        string    `json:"BPN_N_TIER_A"`
			AUTHOR           string    `json:"AUTHOR"`
			BPNOEMBSITEA     string    `json:"BPN_OEM_B_SITE_A"`
			BPNOEMASITEA     string    `json:"BPN_OEM_A_SITE_A"`
			BPNOEMCSITEA     string    `json:"BPN_OEM_C_SITE_A"`
			BPNDISMANTLER    string    `json:"BPN_DISMANTLER"`
			BPNTIERA         string    `json:"BPN_TIER_A"`
			BPNTIERC         string    `json:"BPN_TIER_C"`
			BPNTIERB         string    `json:"BPN_TIER_B"`
			BPNSUBTIERB      string    `json:"BPN_SUB_TIER_B"`
			BPNSUBTIERA      string    `json:"BPN_SUB_TIER_A"`
			BPNSUBTIERC      string    `json:"BPN_SUB_TIER_C"`
			CREATIONDATE     time.Time `json:"CREATION_DATE"`
			BPNTIERCSITEA    string    `json:"BPN_TIER_C_SITE_A"`
			UIDPOOL          string    `json:"UIDPOOL"`
			BPNTIERASITEA    string    `json:"BPN_TIER_A_SITE_A"`
			BPNTIERBSITEA    string    `json:"BPN_TIER_B_SITE_A"`
			BPNSUBTIERBSITEA string    `json:"BPN_SUB_TIER_B_SITE_A"`
			BPNSUBTIERASITEA string    `json:"BPN_SUB_TIER_A_SITE_A"`
			BPNSUBTIERCSITEA string    `json:"BPN_SUB_TIER_C_SITE_A"`
			BPNNTIERASITEA   string    `json:"BPN_N_TIER_A_SITE_A"`
		} `json:"PlainObject,omitempty"`
		UrnBammIoCatenaxPartAsPlanned100PartAsPlanned []struct {
			ValidityPeriod struct {
				ValidFrom time.Time `json:"validFrom"`
				ValidTo   time.Time `json:"validTo"`
			} `json:"validityPeriod"`
			CatenaXId           string `json:"catenaXId"`
			PartTypeInformation struct {
				ManufacturerPartId string `json:"manufacturerPartId"`
				Classification     string `json:"classification"`
				NameAtManufacturer string `json:"nameAtManufacturer"`
			} `json:"partTypeInformation"`
		} `json:"urn:bamm:io.catenax.part_as_planned:1.0.0#PartAsPlanned,omitempty"`
		UrnBammIoCatenaxPartSiteInformationAsPlanned100PartSiteInformationAsPlanned []struct {
			CatenaXId string `json:"catenaXId"`
			Sites     []struct {
				FunctionValidUntil time.Time `json:"functionValidUntil"`
				Function           string    `json:"function"`
				FunctionValidFrom  time.Time `json:"functionValidFrom"`
				CatenaXSiteId      string    `json:"catenaXSiteId"`
			} `json:"sites"`
		} `json:"urn:bamm:io.catenax.part_site_information_as_planned:1.0.0#PartSiteInformationAsPlanned,omitempty"`
		UrnBammIoCatenaxSingleLevelBomAsPlanned102SingleLevelBomAsPlanned []struct {
			CatenaXId  string `json:"catenaXId"`
			ChildParts []struct {
				Quantity struct {
					QuantityNumber  string `json:"quantityNumber"`
					MeasurementUnit struct {
						DatatypeURI  string `json:"datatypeURI"`
						LexicalValue string `json:"lexicalValue"`
					} `json:"measurementUnit"`
				} `json:"quantity"`
				CreatedOn      time.Time `json:"createdOn"`
				LastModifiedOn time.Time `json:"lastModifiedOn"`
				ChildCatenaXId string    `json:"childCatenaXId"`
			} `json:"childParts"`
		} `json:"urn:bamm:io.catenax.single_level_bom_as_planned:1.0.2#SingleLevelBomAsPlanned,omitempty"`
	} `json:"https://catenax.io/schema/TestDataContainer/1.0.0"`
	HttpsCatenaxIoSchemaVehicleBlueprint100 []struct {
		Parent struct {
			Include      []interface{} `json:"include"`
			Code         []interface{} `json:"code"`
			UseId        bool          `json:"useId"`
			ModelVersion string        `json:"modelVersion"`
			Values       struct {
				BPNOEMC          string    `json:"BPN_OEM_C"`
				BPNOEMA          string    `json:"BPN_OEM_A"`
				BPNOEMB          string    `json:"BPN_OEM_B"`
				BPNIRSTEST       string    `json:"BPN_IRS_TEST"`
				BPNNTIERA        string    `json:"BPN_N_TIER_A"`
				AUTHOR           string    `json:"AUTHOR"`
				BPNOEMBSITEA     string    `json:"BPN_OEM_B_SITE_A"`
				BPNOEMASITEA     string    `json:"BPN_OEM_A_SITE_A"`
				BPNOEMCSITEA     string    `json:"BPN_OEM_C_SITE_A"`
				BPNDISMANTLER    string    `json:"BPN_DISMANTLER"`
				BPNTIERA         string    `json:"BPN_TIER_A"`
				BPNTIERC         string    `json:"BPN_TIER_C"`
				BPNTIERB         string    `json:"BPN_TIER_B"`
				BPNSUBTIERB      string    `json:"BPN_SUB_TIER_B"`
				BPNSUBTIERA      string    `json:"BPN_SUB_TIER_A"`
				BPNSUBTIERC      string    `json:"BPN_SUB_TIER_C"`
				CREATIONDATE     time.Time `json:"CREATION_DATE"`
				BPNTIERCSITEA    string    `json:"BPN_TIER_C_SITE_A"`
				UIDPOOL          string    `json:"UIDPOOL"`
				BPNTIERASITEA    string    `json:"BPN_TIER_A_SITE_A"`
				BPNTIERBSITEA    string    `json:"BPN_TIER_B_SITE_A"`
				BPNSUBTIERBSITEA string    `json:"BPN_SUB_TIER_B_SITE_A"`
				BPNSUBTIERASITEA string    `json:"BPN_SUB_TIER_A_SITE_A"`
				BPNSUBTIERCSITEA string    `json:"BPN_SUB_TIER_C_SITE_A"`
				BPNNTIERASITEA   string    `json:"BPN_N_TIER_A_SITE_A"`
			} `json:"values"`
			Count        int    `json:"count"`
			NodeType     string `json:"nodeType"`
			ModelName    string `json:"modelName"`
			Condition    string `json:"condition"`
			InstanceId   string `json:"instanceId"`
			TemplateName string `json:"templateName"`
			Children     []struct {
				Include      []interface{} `json:"include"`
				CatenaXId    string        `json:"catenaXId"`
				Bpnl         string        `json:"bpnl"`
				Code         []interface{} `json:"code"`
				UseId        bool          `json:"useId"`
				ModelVersion string        `json:"modelVersion"`
				Values       struct {
					PartTypeInformationClassification     string `json:"partTypeInformation.classification"`
					PartTypeInformationNameAtManufacturer string `json:"partTypeInformation.nameAtManufacturer"`
					PartTypeInformationManufacturerPartId string `json:"partTypeInformation.manufacturerPartId"`
					ValidityPeriodValidFrom               string `json:"validityPeriod.validFrom"`
					ValidityPeriodValidTo                 string `json:"validityPeriod.validTo"`
				} `json:"values"`
				Count        int    `json:"count"`
				NodeType     string `json:"nodeType"`
				ModelName    string `json:"modelName"`
				Condition    string `json:"condition"`
				InstanceId   string `json:"instanceId"`
				TemplateName string `json:"templateName"`
				Children     []struct {
					Include []struct {
						Condition string        `json:"condition"`
						Code      []interface{} `json:"code"`
						Values    struct {
						} `json:"values"`
						Name    string `json:"name"`
						Count   int    `json:"count"`
						Version string `json:"version"`
					} `json:"include"`
					Code         []string `json:"code"`
					UseId        bool     `json:"useId"`
					ModelVersion string   `json:"modelVersion"`
					Values       struct {
						Sites0CatenaXSiteId      string `json:"sites[0].catenaXSiteId,omitempty"`
						Sites0FunctionValidFrom  string `json:"sites[0].functionValidFrom,omitempty"`
						Sites0FunctionValidUntil string `json:"sites[0].functionValidUntil,omitempty"`
					} `json:"values"`
					Count        int    `json:"count"`
					NodeType     string `json:"nodeType"`
					ModelName    string `json:"modelName"`
					Condition    string `json:"condition"`
					InstanceId   string `json:"instanceId"`
					TemplateName string `json:"templateName"`
					Children     []struct {
						Include      []interface{} `json:"include"`
						CatenaXId    string        `json:"catenaXId"`
						Bpnl         string        `json:"bpnl"`
						Code         []interface{} `json:"code"`
						UseId        bool          `json:"useId"`
						ModelVersion string        `json:"modelVersion"`
						Values       struct {
							PartTypeInformationClassification     string `json:"partTypeInformation.classification"`
							PartTypeInformationNameAtManufacturer string `json:"partTypeInformation.nameAtManufacturer"`
							PartTypeInformationManufacturerPartId string `json:"partTypeInformation.manufacturerPartId"`
							ValidityPeriodValidFrom               string `json:"validityPeriod.validFrom"`
							ValidityPeriodValidTo                 string `json:"validityPeriod.validTo"`
						} `json:"values"`
						Count        int    `json:"count"`
						NodeType     string `json:"nodeType"`
						ModelName    string `json:"modelName"`
						Condition    string `json:"condition"`
						InstanceId   string `json:"instanceId"`
						TemplateName string `json:"templateName"`
						Children     []struct {
							Include []struct {
								Condition string        `json:"condition"`
								Code      []interface{} `json:"code"`
								Values    struct {
								} `json:"values"`
								Name    string `json:"name"`
								Count   int    `json:"count"`
								Version string `json:"version"`
							} `json:"include"`
							Code         []string `json:"code"`
							UseId        bool     `json:"useId"`
							ModelVersion string   `json:"modelVersion"`
							Values       struct {
								Sites0CatenaXSiteId      string `json:"sites[0].catenaXSiteId,omitempty"`
								Sites0FunctionValidFrom  string `json:"sites[0].functionValidFrom,omitempty"`
								Sites0FunctionValidUntil string `json:"sites[0].functionValidUntil,omitempty"`
							} `json:"values"`
							Count        int    `json:"count"`
							NodeType     string `json:"nodeType"`
							ModelName    string `json:"modelName"`
							Condition    string `json:"condition"`
							InstanceId   string `json:"instanceId"`
							TemplateName string `json:"templateName"`
							Children     []struct {
								Include      []interface{} `json:"include"`
								CatenaXId    string        `json:"catenaXId"`
								Bpnl         string        `json:"bpnl"`
								Code         []interface{} `json:"code"`
								UseId        bool          `json:"useId"`
								ModelVersion string        `json:"modelVersion"`
								Values       struct {
									PartTypeInformationClassification     string `json:"partTypeInformation.classification"`
									PartTypeInformationNameAtManufacturer string `json:"partTypeInformation.nameAtManufacturer"`
									PartTypeInformationManufacturerPartId string `json:"partTypeInformation.manufacturerPartId"`
									ValidityPeriodValidFrom               string `json:"validityPeriod.validFrom"`
									ValidityPeriodValidTo                 string `json:"validityPeriod.validTo"`
								} `json:"values"`
								Count        int    `json:"count"`
								NodeType     string `json:"nodeType"`
								ModelName    string `json:"modelName"`
								Condition    string `json:"condition"`
								InstanceId   string `json:"instanceId"`
								TemplateName string `json:"templateName"`
								Children     []struct {
									Include []struct {
										Condition string        `json:"condition"`
										Code      []interface{} `json:"code"`
										Values    struct {
										} `json:"values"`
										Name    string `json:"name"`
										Count   int    `json:"count"`
										Version string `json:"version"`
									} `json:"include"`
									Code         []string `json:"code"`
									UseId        bool     `json:"useId"`
									ModelVersion string   `json:"modelVersion"`
									Values       struct {
										Sites0CatenaXSiteId      string `json:"sites[0].catenaXSiteId,omitempty"`
										Sites0FunctionValidFrom  string `json:"sites[0].functionValidFrom,omitempty"`
										Sites0FunctionValidUntil string `json:"sites[0].functionValidUntil,omitempty"`
									} `json:"values"`
									Count        int    `json:"count"`
									NodeType     string `json:"nodeType"`
									ModelName    string `json:"modelName"`
									Condition    string `json:"condition"`
									InstanceId   string `json:"instanceId"`
									TemplateName string `json:"templateName"`
									Children     []struct {
										Include      []interface{} `json:"include"`
										CatenaXId    string        `json:"catenaXId"`
										Bpnl         string        `json:"bpnl"`
										Code         []interface{} `json:"code"`
										UseId        bool          `json:"useId"`
										ModelVersion string        `json:"modelVersion"`
										Values       struct {
											PartTypeInformationClassification     string `json:"partTypeInformation.classification"`
											PartTypeInformationNameAtManufacturer string `json:"partTypeInformation.nameAtManufacturer"`
											PartTypeInformationManufacturerPartId string `json:"partTypeInformation.manufacturerPartId"`
											ValidityPeriodValidFrom               string `json:"validityPeriod.validFrom"`
											ValidityPeriodValidTo                 string `json:"validityPeriod.validTo"`
										} `json:"values"`
										Count        int    `json:"count"`
										NodeType     string `json:"nodeType"`
										ModelName    string `json:"modelName"`
										Condition    string `json:"condition"`
										InstanceId   string `json:"instanceId"`
										TemplateName string `json:"templateName"`
										Children     []struct {
											Include []struct {
												Condition string        `json:"condition"`
												Code      []interface{} `json:"code"`
												Values    struct {
												} `json:"values"`
												Name    string `json:"name"`
												Count   int    `json:"count"`
												Version string `json:"version"`
											} `json:"include"`
											Code         []string `json:"code"`
											UseId        bool     `json:"useId"`
											ModelVersion string   `json:"modelVersion"`
											Values       struct {
												Sites0CatenaXSiteId      string `json:"sites[0].catenaXSiteId,omitempty"`
												Sites0FunctionValidFrom  string `json:"sites[0].functionValidFrom,omitempty"`
												Sites0FunctionValidUntil string `json:"sites[0].functionValidUntil,omitempty"`
											} `json:"values"`
											Count        int    `json:"count"`
											NodeType     string `json:"nodeType"`
											ModelName    string `json:"modelName"`
											Condition    string `json:"condition"`
											InstanceId   string `json:"instanceId"`
											TemplateName string `json:"templateName"`
											Children     []struct {
												Include      []interface{} `json:"include"`
												CatenaXId    string        `json:"catenaXId"`
												Bpnl         string        `json:"bpnl"`
												Code         []interface{} `json:"code"`
												UseId        bool          `json:"useId"`
												ModelVersion string        `json:"modelVersion"`
												Values       struct {
													PartTypeInformationClassification     string `json:"partTypeInformation.classification"`
													PartTypeInformationNameAtManufacturer string `json:"partTypeInformation.nameAtManufacturer"`
													PartTypeInformationManufacturerPartId string `json:"partTypeInformation.manufacturerPartId"`
													ValidityPeriodValidFrom               string `json:"validityPeriod.validFrom"`
													ValidityPeriodValidTo                 string `json:"validityPeriod.validTo"`
												} `json:"values"`
												Count        int    `json:"count"`
												NodeType     string `json:"nodeType"`
												ModelName    string `json:"modelName"`
												Condition    string `json:"condition"`
												InstanceId   string `json:"instanceId"`
												TemplateName string `json:"templateName"`
												Children     []struct {
													Include      []interface{} `json:"include"`
													Code         []string      `json:"code"`
													UseId        bool          `json:"useId"`
													ModelVersion string        `json:"modelVersion"`
													Values       struct {
														Sites0CatenaXSiteId      string `json:"sites[0].catenaXSiteId,omitempty"`
														Sites0FunctionValidFrom  string `json:"sites[0].functionValidFrom,omitempty"`
														Sites0FunctionValidUntil string `json:"sites[0].functionValidUntil,omitempty"`
													} `json:"values"`
													Count           int           `json:"count"`
													NodeType        string        `json:"nodeType"`
													ModelName       string        `json:"modelName"`
													Condition       string        `json:"condition"`
													InstanceId      string        `json:"instanceId"`
													TemplateName    string        `json:"templateName"`
													Children        []interface{} `json:"children"`
													UseTemplate     bool          `json:"useTemplate"`
													Comment         string        `json:"comment"`
													TemplateVersion string        `json:"templateVersion"`
												} `json:"children"`
												UseTemplate     bool   `json:"useTemplate"`
												Comment         string `json:"comment"`
												TemplateVersion string `json:"templateVersion"`
											} `json:"children"`
											UseTemplate     bool   `json:"useTemplate"`
											Comment         string `json:"comment"`
											TemplateVersion string `json:"templateVersion"`
										} `json:"children"`
										UseTemplate     bool   `json:"useTemplate"`
										Comment         string `json:"comment"`
										TemplateVersion string `json:"templateVersion"`
									} `json:"children"`
									UseTemplate     bool   `json:"useTemplate"`
									Comment         string `json:"comment"`
									TemplateVersion string `json:"templateVersion"`
								} `json:"children"`
								UseTemplate     bool   `json:"useTemplate"`
								Comment         string `json:"comment"`
								TemplateVersion string `json:"templateVersion"`
							} `json:"children"`
							UseTemplate     bool   `json:"useTemplate"`
							Comment         string `json:"comment"`
							TemplateVersion string `json:"templateVersion"`
						} `json:"children"`
						UseTemplate     bool   `json:"useTemplate"`
						Comment         string `json:"comment"`
						TemplateVersion string `json:"templateVersion"`
					} `json:"children"`
					UseTemplate     bool   `json:"useTemplate"`
					Comment         string `json:"comment"`
					TemplateVersion string `json:"templateVersion"`
				} `json:"children"`
				UseTemplate     bool   `json:"useTemplate"`
				Comment         string `json:"comment"`
				TemplateVersion string `json:"templateVersion"`
			} `json:"children"`
			UseTemplate     bool   `json:"useTemplate"`
			Comment         string `json:"comment"`
			TemplateVersion string `json:"templateVersion"`
		} `json:"parent"`
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"https://catenax.io/schema/VehicleBlueprint/1.0.0"`
}
