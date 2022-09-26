package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateImage(c *gin.Context) {
	dst := "images/"
	file, _ := c.FormFile("file")
	ext := filepath.Ext(file.Filename)

	filename := createFileName(ext)

	// Upload the file to specific dst.
	err := c.SaveUploadedFile(file, fmt.Sprintf("%v/%v", dst, filename))

	if err != nil {
		fmt.Println("Error Saving file", err)
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/images/%v", filename))
}

func createFileName(ext string) string {
	return time.Now().Format("20060102150405") + ext
}
