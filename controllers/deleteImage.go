package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DeleteImage(c *gin.Context) {
	dst := "images/"
	filename := c.Param("filename")

	err := os.Remove(dst + filename)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't found image",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Image deleted",
	})
}
