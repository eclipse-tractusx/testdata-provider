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

package main

import (
	"cx/internal"
	psql "cx/internal/psql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/data/:id", getData)
	router.POST("/import/:bpnl", importData)

	router.Run("localhost:8080")
}

func getData(c *gin.Context) {

	dataId := c.Param("id")
	db := psql.NewInstance()

	data, err := db.LoadData(dataId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "not found")
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", data)
}

func importData(c *gin.Context) {

	// TODO read CX_Testdata_v1.4.1-AsPlanned.json
	bpnl := c.Param("bpnl")

	err := internal.ImportTestDataByBPNL(nil, bpnl)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
}
