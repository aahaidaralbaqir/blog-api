package repositories

import (
	"database/sql"
	"fmt"
	"go-crash-course/database"
	"go-crash-course/entities"
)

type IPostRepository interface {
	Save(*entities.Post) (*entities.Post)
	FindAll() ([]*entities.Post, error)
	WithAuthor() ([]*entities.PostAuthor, error)
	FindById(id int) (*entities.Post,error)
	Destroy(id int) error
	Update(post *entities.Post) (*entities.Post,error)
}

type PostRepository struct {
	Conn *sql.DB
}

func (p *PostRepository) Save(post *entities.Post) (*entities.Post) {
	_,err := p.Conn.Exec("INSERT INTO posts values(?,?,?,?)",nil,post.Title,post.Description,post.AuthorId)
	if err != nil {
		return nil
	}
	return post
}

func (p *PostRepository) FindById(id int) (*entities.Post,error) {
	result := &entities.Post{}
	err := p.Conn.QueryRow("SELECT * FROM posts WHERE id = ?",id).Scan(&result.ID,&result.Title,&result.Description,&result.AuthorId)
	if err != nil {
		fmt.Println("ERROR",err.Error())
		return nil,err
	}

	return result,nil
}


func (p *PostRepository) WithAuthor() ([]*entities.PostAuthor, error) {
	var PostAuthor []*entities.PostAuthor
	rows, err := p.Conn.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Println("WITH AUTHOR", err.Error())
	}

	for rows.Next() {
		var each = &entities.PostAuthor{}
		var err = rows.Scan(&each.ID, &each.Description, &each.Title, &each.AuthorId)
		if err != nil {
			fmt.Println(err.Error())
		}

		PostAuthor = append(PostAuthor, each)
	}
	return PostAuthor, nil
}

func (p *PostRepository) Destroy(id int) error  {
	_, err := p.Conn.Exec("DELETE FROM posts WHERE id = ?",id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostRepository) Update(post *entities.Post) (*entities.Post, error) {
	_, err := p.Conn.Exec("UPDATE posts SET title = ?, description = ? ,  author_id = ? WHERE id = ?",post.Title,post.Description,post.AuthorId,post.ID)
	if err != nil {
		return nil, err
	}
	return post,nil
}

func (p *PostRepository) FindAll() ([]*entities.Post, error) {

	return nil, nil
}

func NewPostRepository() IPostRepository {
	return &PostRepository{Conn: database.GetConnection()}
}
