package main

import (
	"github.com/aryanicosa/golang-gin/controller"
	"github.com/aryanicosa/golang-gin/service"
	"github.com/gin-gonic/gin"
)

var(
	postService service.PostService = service.New()
	postController controller.PostController = controller.New(postService)
)

func main() {
	server := gin.Default()

	//first endpoint
	// server.GET("/test", func(ctx *gin.Context) { //create a lambda function  named ctx reference to gin.Context
	// 	ctx.JSON(200, gin.H{
	// 		"message" : "OK",
	// 	})
	// })

	server.GET("/post", func(ctx *gin.Context) {
		ctx.JSON(200, postController.FindAll())
	})

	server.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(201, postController.Save(ctx))
	})

	server.Run(":8080")
}