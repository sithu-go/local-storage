package grpccontrollers

import (
	"context"
	"fmt"
	"local-storage/filegrpc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context) {

	filename := c.Param("filename")

	imgGetRequest := filegrpc.ImageGetRequest{
		Filename: filename,
	}

	imageClient := filegrpc.NewFileClient(filegrpc.Conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	imageResponse, err := imageClient.GetImage(ctx, &imgGetRequest)

	if err != nil {
		fmt.Println("Error getting file", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Writer.Write(imageResponse.GetFileData())

}
