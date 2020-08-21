package repositories

import (
	"github.com/jinzhu/gorm"
	"go-crash-course/database"
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
	var Post []*entities.Post
	p.Conn.Raw("SELECT * FROM posts").Scan(&Post)
	return Post, nil
}

func NewPostRepository() IPostRepository {
	return &PostRepository{Conn: database.GetConnection()}
}