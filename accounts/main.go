package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router gin.Engine
}

func AccountDetails(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "test"})
}

func main() {
	router := gin.Default()
	k := NewMiddleware()

	router.Use(k.Session())
	router.GET("/details", AccountDetails)
	router.Run(":8001")
}
