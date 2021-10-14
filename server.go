package main

import (
	"io"
	"os"

	"github.com/aryanicosa/golang-gin/controller"
	"github.com/aryanicosa/golang-gin/middleware"
	"github.com/aryanicosa/golang-gin/service"
	"github.com/gin-gonic/gin"
	"github.com/tpkeeper/gin-dump"
)

var(
	postService service.PostService = service.New()
	postController controller.PostController = controller.New(postService)
)

//setting up the output to a log file
func SetupLogOutput() {
	f, _ := os.Create("gin.log") //create custom log by a file. Error ignored for a while
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //assign a file and standar log output
}

func main() {
	SetupLogOutput()

	server := gin.New() //create new instance of server

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	server.GET("/post", func(ctx *gin.Context) {
		ctx.JSON(200, postController.FindAll())
	})

	server.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(201, postController.Save(ctx))
	})

	server.Run(":8080")
}