package main

import "github.com/gin-gonic/gin"
import "github.com/rogeralsing/Better/registration/register"

func main() {
	router := gin.New()

	register.CreateRegisterService(router)

	router.Run(":8081")
}
