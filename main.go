package main

import (
	"local-storage/controllers"
	"local-storage/filegrpc"
	"local-storage/grpccontrollers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	filegrpc.Connectgrpc()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	images := router.Group("/images")
	{
		images.GET("/", controllers.GetImages)
		images.POST("/", controllers.CreateImage)
		images.GET("/:filename", controllers.GetImage)
		images.DELETE("/:filename", controllers.DeleteImage)

	}
	grpcIMG := router.Group("/grpc/images")
	{
		grpcIMG.POST("/", grpccontrollers.CreateImage)
		grpcIMG.GET("/:filename", grpccontrollers.GetImage)
		grpcIMG.DELETE("/:filename", grpccontrollers.DeleteImage)
	}
	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
