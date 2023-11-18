package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/kratos-client-go"
)

var (
	PORT         = os.Getenv("PORT")
	KRATOS_URL   = os.Getenv("KRATOS_URL")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
)

func main() {

	port, err := strconv.Atoi(PORT)
	if err != nil {
		panic(err)
	}

	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{{URL: KRATOS_URL}}
	client := ory.NewAPIClient(configuration)
	router := gin.Default()

	router.GET("/api/v2/auth/get-registration-form", func(c *gin.Context) {
		req := client.FrontendApi.CreateBrowserRegistrationFlow(c)
		_, resp, err := client.FrontendApi.CreateBrowserRegistrationFlowExecute(req)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get login flow"})
			return
		}
		ui, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read ui body"})
			return
		}
		var uistruct map[string]interface{}
		err = json.Unmarshal(ui, &uistruct)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read ui body"})
			return
		}
		c.JSON(http.StatusOK, uistruct)
	})

	router.Run(fmt.Sprintf(":%d", port))
}
