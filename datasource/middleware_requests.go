package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRequestCounter() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Request.Header.Get("X-User")
		if user_id == "" {
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:4455/login")
			return
		}
		// c.Set("user_id", user_id)
		// TODO: save user_id and request path to database
		// this will need to calculate service usage.
		c.Next()
	}
}
