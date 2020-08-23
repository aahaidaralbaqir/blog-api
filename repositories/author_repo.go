package repositories

import (
	"database/sql"
	"fmt"
	"go-crash-course/database"
	"go-crash-course/entities"
)

type IAuthorRepository interface {
	FindById (id int) *entities.Author
	FindAll() []*entities.Author
}

type AuthorRepository struct {
	Conn *sql.DB
}

func (a *AuthorRepository) FindById(id int) *entities.Author {
	result := &entities.Author{}
	err := a.Conn.QueryRow("SELECT * FROM authors WHERE id = ?",id).Scan(&result.ID,&result.Name)

	if err != nil {
		fmt.Print("FIND BY ID AUTHOR REPOSITORY",err.Error())
		return nil
	}

	return result
}

func (a *AuthorRepository) FindAll() []*entities.Author {
	var Author []*entities.Author

	rows,err := a.Conn.Query("SELECT * FROM authors")

	if err != nil {
		fmt.Println("AuthorFindAll",err.Error())
	}

	for rows.Next() {
		var each = &entities.Author{}
		err := rows.Scan(&each.ID,&each.Name)

		if err != nil {
			fmt.Println("ERROR",err.Error())
		}

		Author = append(Author,each)
	}

	return Author
}


func NewAuthorRepository() IAuthorRepository {
	return &AuthorRepository{Conn: database.GetConnection()}
}



