package grpccontrollers

import (
	"context"
	"fmt"
	"io"
	"local-storage/filegrpc"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateImage(c *gin.Context) {

	file, err := c.FormFile("file")
	ext := filepath.Ext(file.Filename)
	if err != nil {
		fmt.Println("Error Parsing file", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileMultipart, err := file.Open()
	if err != nil {
		fmt.Println("Error Opening file", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileBytes, err := io.ReadAll(fileMultipart)

	if err != nil {
		fmt.Println("Error reading file", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	imgRequest := filegrpc.ImageCreateRequest{
		FileData: fileBytes,
		Ext:      ext,
	}

	fileClient := filegrpc.NewFileClient(filegrpc.Conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	imageResponse, err := fileClient.CreateImage(ctx, &imgRequest)
	if err != nil {
		fmt.Println("Error creating file", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, imageResponse)
}
