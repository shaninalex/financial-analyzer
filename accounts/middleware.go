package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/kratos-client-go"
)

type kratosMiddleware struct {
	ory *ory.APIClient
}

func NewMiddleware() *kratosMiddleware {
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4434", // Kratos Admin API
		},
	}
	return &kratosMiddleware{
		ory: ory.NewAPIClient(configuration),
	}
}

func (k *kratosMiddleware) Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := k.validateSession(c.Request)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:4455/login")
			return
		}
		if !*session.Active {
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:4455/login")
			return
		}

		user_id := c.Request.Header.Get("X-User")
		if user_id == "" {
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:4455/login")
			return
		}

		c.Set("user_id", user_id)

		c.Next()
	}
}

func (k *kratosMiddleware) validateSession(r *http.Request) (*ory.Session, error) {
	cookie, err := r.Cookie("ory_kratos_session")
	if err != nil {
		return nil, err
	}
	if cookie == nil {
		return nil, errors.New("no session found in cookie")
	}
	resp, _, err := k.ory.FrontendApi.ToSession(context.Background()).Cookie(cookie.String()).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
