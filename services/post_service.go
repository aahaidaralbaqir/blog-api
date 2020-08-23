package services

import (
	"go-crash-course/entities"
	"go-crash-course/repositories"
)

type PostService struct {
	PostRepository repositories.IPostRepository
	AuthorRepository repositories.IAuthorRepository
}

func (u *PostService) GetPost() []*entities.Post {
	posts, _ := u.PostRepository.FindAll()
	return posts
}

func (u *PostService) GetPostWithAuthor() []*entities.PostAuthor {
	posts, _ := u.PostRepository.WithAuthor()
	for _,post := range posts {
		post.Author = u.AuthorRepository.FindById(post.AuthorId)
	}
	return posts
}

func (u *PostService) SavePost(data *entities.Post) (post *entities.Post, err error) {
	result, err := u.PostRepository.Save(data)
	return result, nil
}

func NewPostService() *PostService {
	return &PostService{
		PostRepository: repositories.NewPostRepository(),
		AuthorRepository: repositories.NewAuthorRepository(),
	}
}
