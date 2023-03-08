package edc

import "os"

func getConnectorUrl() string {
	return os.Getenv("CONNECTOR_URL")
}

func getConnectorApiKey() string {
	return os.Getenv("CONNECTOR_API_KEY")
}
