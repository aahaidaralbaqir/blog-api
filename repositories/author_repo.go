package repositories

import (
	"database/sql"
	"fmt"
	"go-crash-course/database"
	"go-crash-course/entities"
)

type IAuthorRepository interface {
	FindById(id int) *entities.Author
	FindAll() []*entities.Author
	Store(post *entities.Author) *entities.Author
	Destroy(id int) error
	Update(post *entities.Author) (*entities.Author,error)
}


type AuthorRepository struct {
	Conn *sql.DB
}

func (a *AuthorRepository) FindById(id int) *entities.Author {
	result := &entities.Author{}
	err := a.Conn.QueryRow("SELECT * FROM authors WHERE id = ?", id).Scan(&result.ID, &result.Name)

	if err != nil {
		fmt.Print("FIND BY ID AUTHOR REPOSITORY", err.Error())
		return nil
	}

	return result
}

func (a *AuthorRepository) Store(author *entities.Author) *entities.Author {
	_,err := a.Conn.Exec("INSERT INTO authors values(?,?,?,?,?,?)",nil,author.Name,author.Username,author.Password,author.Bio,author.Skills)
	if err != nil {
		return nil
	}
	return author
}

func (a *AuthorRepository) FindAll() []*entities.Author {
	var Author []*entities.Author

	rows, err := a.Conn.Query("SELECT * FROM authors")

	if err != nil {
		fmt.Println("AuthorFindAll", err.Error())
	}

	for rows.Next() {
		var each = &entities.Author{}
		err := rows.Scan(&each.ID, &each.Name,&each.Username,&each.Password,&each.Bio,&each.Skills)

		if err != nil {
			fmt.Println("ERROR", err.Error())
		}

		Author = append(Author, each)
	}

	return Author
}

func (a *AuthorRepository) Destroy(id int) error {
	_, err := a.Conn.Exec("DELETE FROM authors WHERE id = ?",id)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthorRepository) Update(author *entities.Author) (*entities.Author,error ) {
	_, err := a.Conn.Exec("UPDATE authors SET name = ? WHERE id = ?  ",author.Name,author.ID)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func NewAuthorRepository() IAuthorRepository {
	return &AuthorRepository{Conn: database.GetConnection()}
}
