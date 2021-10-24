package controller

import (
	"net/http"

	"github.com/aryanicosa/golang-gin/entity"
	"github.com/aryanicosa/golang-gin/service"
	"github.com/aryanicosa/golang-gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostController interface {
	FindAll() []entity.Post //slice to store post
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.PostService
}

var validate *validator.Validate //create custom validator

//constructor function
func New(service service.PostService) PostController {
	validate = validator.New()
	validate.RegisterValidation("is-okay", validators.ValidateOkayTitle) //validating post request
	return &controller {
		service: service,
	}
}

func (c *controller) FindAll() []entity.Post {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var post entity.Post
	err := ctx.ShouldBindJSON(&post)
	if err != nil {
		return err
	}
	err = validate.Struct(post)
	if err != nil {
		return err
	}
	//ctx.BindJSON(&post)
	c.service.Save(post)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	posts := c.service.FindAll()

	data := gin.H{
		"title" : "Post page",
		"posts"	: posts,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}