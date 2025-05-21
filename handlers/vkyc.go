package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitiateVKYC(c *gin.Context) {
	userID := c.GetString("user_id")
	role := c.GetString("role")

	c.JSON(http.StatusOK, gin.H{
		"message":   "vKYC session initiated",
		"user_id":   userID,
		"role":      role,
		"vkyc_link": "https://vkyc.nayak.com/v1/video?token=securetoken#step1",
	})
}
