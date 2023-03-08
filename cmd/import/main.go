package main

import (
	psql "cx/internal/psql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/data/:id", getData)
	router.POST("/import", importData)

	router.Run("localhost:8080")
}

func getData(c *gin.Context) {

	dataId := c.Param("id")
	db := psql.NewInstance()

	data, err := db.Load(dataId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "not found")
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", data)
}

func importData(c *gin.Context) {

	// TODO
}
