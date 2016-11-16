package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rogeralsing/Better/customer/login"
)

func main() {

	router := gin.New()

	login.CreateLoginService(router)

	router.Run()
}
