package controller

import (
	"github.com/aryanicosa/golang-gin/entity"
	"github.com/aryanicosa/golang-gin/service"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	FindAll() []entity.Post //slice to store post
	Save(ctx *gin.Context) entity.Post
}

type controller struct {
	service service.PostService
}

//constructor function
func New(service service.PostService) PostController {
	return &controller {
		service: service,
	}
}

func (c *controller) FindAll() []entity.Post {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Post {
	var post entity.Post
	ctx.BindJSON(&post)
	c.service.Save(post)
	return post
}