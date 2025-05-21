package main

import (
	"jwt_auth_vkyc/handlers"
	"jwt_auth_vkyc/middleware"
	"jwt_auth_vkyc/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Token generation for testing
	r.GET("/generate-token", func(c *gin.Context) {
		token, err := utils.GenerateJWT("sejalnayak", "customer")
		if err != nil {
			c.JSON(500, gin.H{"error": "token generation failed"})
			return
		}
		c.JSON(200, gin.H{"token": token})
	})

	// Secured route
	secured := r.Group("/v1", middleware.JWTAuthMiddleware())
	{
		secured.POST("/initiate-vkyc", handlers.InitiateVKYC)
	}

	r.Run(":8080")
}
