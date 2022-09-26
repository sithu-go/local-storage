package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context) {
	filename := c.Param("filename")
	data, err := os.ReadFile("images/" + filename)
	if len(data) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesage": "not found",
		})
	}
	if err != nil {
		fmt.Println("error reading", err)
	}
	// c.Writer.Write(data)
	c.File("images/" + filename)

}
