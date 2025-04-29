package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	// r := gin.Default()
	r := gin.New()

	r.GET("/v1/user/:id", func(c *gin.Context) {
		myuser := struct {
			Name string `json:"name"`
		}{
			Name: c.Param("id"),
		}

		c.Set("Content-Type", "application/json")
		c.JSON(http.StatusOK, myuser)
	})

	r.Run(":8080")
}
