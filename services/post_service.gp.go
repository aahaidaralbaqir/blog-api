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

func (u *PostService) GetPostWithAuthor() []*entities.PostAuthor {
	postsWihAuthor,_ := u.Repository.WithAuthor()
	return postsWihAuthor
}

func (u *PostService) SavePost(data *entities.Post) (post *entities.Post,err error) {
	result, err := u.Repository.Save(data)
	return result,nil
}

func NewPostService() PostService {
	return PostService{
		Repository: repositories.NewPostRepository(),
	}
}
