package service

import (
	"github.com/aryanicosa/golang-gin/entity"
)

type PostService interface {
	Save(entity.Post) entity.Post
	FindAll() []entity.Post
}

type postService struct {
	post []entity.Post
}

//implement the interface
func New() PostService {
	return &postService{}
}

func (service *postService) Save(post entity.Post) entity.Post {
	service.post = append(service.post, post)
	return post
}

func (service *postService) FindAll() []entity.Post {
	return service.post
}