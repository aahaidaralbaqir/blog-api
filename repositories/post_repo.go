package repositories

import (
	"github.com/jinzhu/gorm"
	"go-crash-course/entities"
)

type IPostRepository interface {
	Save(*entities.Post) (*entities.Post, error)
	FindAll() ([]*entities.Post, error)
}

type PostRepository struct {
	Conn *gorm.DB
}

func (p *PostRepository) Save(post *entities.Post) (*entities.Post, error) {
	return nil, nil
}

func (p *PostRepository) FindAll() ([]*entities.Post, error) {
	return nil, nil
}

func NewPostRepository(conn *gorm.DB) IPostRepository {
	return &PostRepository{Conn: conn}
}