package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/kratos-client-go"
)

var (
	PORT         = os.Getenv("PORT")
	KRATOS_URL   = os.Getenv("KRATOS_URL")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
)

func main() {

	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{{URL: KRATOS_URL}}
	client := ory.NewAPIClient(configuration)
	router := gin.Default()

	router.GET("/api/v2/auth/get-registration-form", func(c *gin.Context) {
		req := client.FrontendApi.CreateBrowserRegistrationFlow(c)
		_, resp, err := client.FrontendApi.CreateBrowserRegistrationFlowExecute(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get login flow"})
		}
		ui, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read ui body"})
		}
		var uistruct map[string]interface{}
		err = json.Unmarshal(ui, &uistruct)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read ui body"})
		}
		c.JSON(http.StatusOK, uistruct)
	})
	router.Run(":8081")
}
