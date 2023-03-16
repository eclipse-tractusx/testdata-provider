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

package db

import (
	"cx/internal/psql/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
}

func NewInstance() Database {
	connStr := fmt.Sprintf("user=%s port=%d user=%s password=%s dbname=%s sslmode=disable", getHost(), getPort(), getUser(), getPassword(), getDbName())

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return Database{
		connection: db,
	}
}

func (db Database) Save(data *model.TestDataImport) error {

	result := db.connection.Create(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db Database) LoadData(id string) ([]byte, error) {

	query := model.TestDataImportEntry{
		ID: id,
	}

	result := db.connection.First(&query)
	if result.Error != nil {
		return query.Data, result.Error
	}

	return query.Data, nil
}

func (db Database) LoadImport(id string) (model.TestDataImport, error) {

	query := model.TestDataImport{
		ID: id,
	}

	result := db.connection.First(&query)
	if result.Error != nil {
		return query, result.Error
	}

	return query, nil
}
