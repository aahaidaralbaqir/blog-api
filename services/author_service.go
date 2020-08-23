package services

import (
	"go-crash-course/entities"
	"go-crash-course/repositories"
)

type AuthorService struct {
	AuthorRepository repositories.IAuthorRepository
}

func (a *AuthorService)  FetchAuthor() []*entities.Author {
	authors := a.AuthorRepository.FindAll()
	return authors
}

func (a *AuthorService) FetchAuthorById(id int) *entities.Author {
	author := a.AuthorRepository.FindById(id)
	return author
}


func NewAuthorService() *AuthorService {
	return &AuthorService{
		AuthorRepository: repositories.NewAuthorRepository(),
	}
}

