package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
