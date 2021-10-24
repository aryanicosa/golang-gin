package main

import (
	"io"
	"net/http"
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

	server.Static("/css,", "./templates/assets/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	//grouping Api Request
	apiRoutes := server.Group("/api") 
	{
		apiRoutes.GET("/post", func(ctx *gin.Context) {
			ctx.JSON(200, postController.FindAll())
		})
	
		apiRoutes.POST("/post", func(ctx *gin.Context) {
			err := postController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{ "message": "Post input is valid!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/post", postController.ShowAll)
	}
	

	server.Run(":8080")
}