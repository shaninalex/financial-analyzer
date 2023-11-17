package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/kratos-client-go"
)

func main() {

	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			// TODO: move to env variable
			URL: "http://127.0.0.1:4433", // Kratos Admin API
		},
	}
	client := ory.NewAPIClient(configuration)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		req := client.FrontendApi.CreateBrowserLoginFlow(c)
		_, resp, err := client.FrontendApi.CreateBrowserLoginFlowExecute(req)
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
		c.JSON(http.StatusOK, uistruct["ui"])
	})
	router.Run(":8005")
}
