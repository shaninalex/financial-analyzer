/*
API to communicate with frontend. Allow user to:
public:
- login
- register

private:
- change user data
- receive all nesessery data to build ui
- receive and manage reports list
*/
package api

import "github.com/gin-gonic/gin"

type Api struct {
	router gin.Engine
}

func InitializeApi() (*Api, error) {

	api := &Api{
		router: *gin.Default(),
	}

	return api, nil
}
