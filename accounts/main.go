package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ory "github.com/ory/kratos-client-go"
)

type App struct {
	Router gin.Engine
}

func main() {

	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4434", // Kratos Admin API
		},
	}
	client := ory.NewAPIClient(configuration)
	router := gin.Default()
	k := NewMiddleware()

	router.Use(k.Session())
	router.GET("/details", func(c *gin.Context) {

		user_id := c.MustGet("user_id").(string)
		req := client.IdentityApi.GetIdentity(c, user_id)
		identity, _, err := client.IdentityApi.GetIdentityExecute(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get identity"})
		}

		c.JSON(http.StatusOK, identity)
	})
	router.Run(":8001")
}
