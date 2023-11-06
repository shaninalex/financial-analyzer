package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) AlphavantageOverview(c *gin.Context) {
	ticker := c.Query("ticker")
	if ticker == "" {
		ErrorResponse(c, http.StatusBadRequest, "ticker query is required")
		return
	}

	data, err := api.Datasource.Alphavantage.Overview(ticker)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
