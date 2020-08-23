package repositories

import (
	"database/sql"
	"fmt"
	"go-crash-course/database"
	"go-crash-course/entities"
)

type IPostRepository interface {
	Save(*entities.Post) (*entities.Post, error)
	FindAll() ([]*entities.Post, error)
	WithAuthor() ([]*entities.PostAuthor, error)
}

type PostRepository struct {
	Conn *sql.DB
}

func (p *PostRepository) Save(post *entities.Post) (*entities.Post, error) {
	return nil, nil
}

func (p *PostRepository) WithAuthor() ([]*entities.PostAuthor, error) {
	var PostAuthor []*entities.PostAuthor
	rows,err := p.Conn.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Println("WITH AUTHOR", err.Error())
	}

	for rows.Next() {
		var each = &entities.PostAuthor{}
		var err = rows.Scan(&each.ID,&each.Description,&each.Title,&each.AuthorId)
		if err != nil {
			fmt.Println(err.Error())
		}

		PostAuthor = append(PostAuthor,each)
	}
	return PostAuthor, nil
}

func (p *PostRepository) FindAll() ([]*entities.Post, error) {

	return nil, nil
}

func NewPostRepository() IPostRepository {
	return &PostRepository{Conn: database.GetConnection()}
}
