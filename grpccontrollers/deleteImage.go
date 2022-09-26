package grpccontrollers

import (
	"context"
	"fmt"
	"local-storage/filegrpc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DeleteImage(c *gin.Context) {

	filename := c.Param("filename")

	imgdeleteRequest := filegrpc.ImageDeleteRequest{
		Filename: filename,
	}

	imageClient := filegrpc.NewFileClient(filegrpc.Conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	imageResponse, err := imageClient.DeleteImage(ctx, &imgdeleteRequest)

	if err != nil {
		fmt.Println("Error deleting file", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// c.Writer.Write(imageResponse.())
	c.JSON(http.StatusOK, imageResponse)

}
