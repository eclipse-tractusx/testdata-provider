package db

import (
	"os"
	"strconv"
)

func getHost() string {
	return os.Getenv("DB_HOST")
}

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	return port
}

func getUser() string {
	return os.Getenv("DB_USER")
}

func getPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func getDbName() string {
	return os.Getenv("DB_NAME")
}
