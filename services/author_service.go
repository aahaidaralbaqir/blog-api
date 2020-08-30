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

func (a *AuthorService) SaveAuthor(data *entities.Author) *entities.Author {
	result := a.AuthorRepository.Store(data)
	return result
}

func (a *AuthorService) FetchAuthorById(id int) *entities.Author {
	author := a.AuthorRepository.FindById(id)
	return author
}

func (a *AuthorService) Destroy(id int) error {
	err := a.AuthorRepository.Destroy(id)

	if err != nil {
		return err
	}
	return nil
}


func (a *AuthorService) Update(author *entities.Author) (*entities.Author,error){
	result,err := a.AuthorRepository.Update(author)
	if err != nil {
		return nil,err
	}
	return result, nil
}

func NewAuthorService() *AuthorService {
	return &AuthorService{
		AuthorRepository: repositories.NewAuthorRepository(),
	}
}

