package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//LoginRequest is the type that defines the import type.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//LoginResponse is the type that defines the import type.
type LoginResponse struct {
	Session string `json:"session"`
}

//Login will check stuff
func login(request LoginRequest) (LoginResponse, error) {
	login := LoginResponse{"Success"}
	return login, nil
}

func CreateLoginService(router *gin.Engine) {
	router.POST("/Login", func(c *gin.Context) {

		var request LoginRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, err.Error)
			return
		}

		response, err := login(request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error)
			return
		}

		c.JSON(http.StatusOK, response)
	})
}
