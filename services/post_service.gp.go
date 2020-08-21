package services

import (
	"go-crash-course/entities"
	"go-crash-course/repositories"
)

type PostService struct {
	Repository repositories.IPostRepository
}

func (u *PostService) GetPost() []*entities.Post {
	posts,_ := u.Repository.FindAll()
	return posts
}

func NewPostService() PostService {
	return PostService{
		Repository: repositories.NewPostRepository(),
	}
}
