package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Request is the type that defines the import type.
type Request struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func register(Request) error {
	return nil
}

// CreateRegisterService is the...
func CreateRegisterService(router *gin.Engine) {
	router.POST("/Register", func(c *gin.Context) {

		var request Request
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, err.Error)
			return
		}

		if err := register(request); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error)
			return
		}
		c.JSON(http.StatusOK, nil)
	})
}
