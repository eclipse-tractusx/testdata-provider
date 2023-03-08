package db

import (
	"database/sql"
	"fmt"
)

type Database struct {
	connection *sql.DB
}

func NewInstance() Database {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", getHost(), getPort(), getUser(), getPassword(), getDbName())

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	return Database{
		connection: db,
	}
}

func (db Database) Save(id string, data []byte) {

	createStmt := `CREATE TABLE IF NOT EXISTS data (id VARCHAR(255) PRIMARY KEY, data BYTEA)`
	insertStmt := `INSERT INTO data (id, data) VALUES ($1, $2)`

	_, err := db.connection.Exec(createStmt)
	if err != nil {
		panic(err)
	}

	_, err = db.connection.Exec(insertStmt, id, data)
	if err != nil {
		panic(err)
	}
}

func (db Database) Load(id string) ([]byte, error) {
	selectStmt := `SELECT data FROM data WHERE id = $1`

	var data []byte
	err := db.connection.QueryRow(selectStmt, id).Scan(&data)
	if err != sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		panic(err)
	}

	return data, nil
}
