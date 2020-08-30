package repositories

import (
	"database/sql"
	"fmt"
	"go-crash-course/database"
	"go-crash-course/entities"
)

type ITagRepository interface {
	FindById(id int) *entities.Tag
	FindAll() []*entities.Tag
	Store(post *entities.Tag) *entities.Tag
	Destroy(id int) error
	Update(post *entities.Tag) (*entities.Tag,error)
}


type TagRepository struct {
	Conn *sql.DB
}

func (a *TagRepository) FindById(id int) *entities.Tag {
	result := &entities.Tag{}
	err := a.Conn.QueryRow("SELECT * FROM tags WHERE id = ?", id).Scan(&result.ID, &result.Name)

	if err != nil {
		fmt.Print("FIND BY ID TAG REPOSITORY", err.Error())
		return nil
	}

	return result
}

func (a *TagRepository) Store(tag *entities.Tag) *entities.Tag {
	_,err := a.Conn.Exec("INSERT INTO tags values(?,?)",nil,tag.Name)
	if err != nil {
		return nil
	}
	return tag
}

func (a *TagRepository) FindAll() []*entities.Tag {
	var tag []*entities.Tag

	rows, err := a.Conn.Query("SELECT * FROM tags")

	if err != nil {
		fmt.Println("TAG FIND ALL ERROR", err.Error())
	}

	for rows.Next() {
		var each = &entities.Tag{}
		err := rows.Scan(&each.ID, &each.Name)

		if err != nil {
			fmt.Println("ERROR", err.Error())
		}

		tag = append(tag, each)
	}

	return tag
}

func (a *TagRepository) Destroy(id int) error {
	_, err := a.Conn.Exec("DELETE FROM tags WHERE id = ?",id)
	if err != nil {
		return err
	}
	return nil
}

func (a *TagRepository) Update(tag *entities.Tag) (*entities.Tag,error ) {
	_, err := a.Conn.Exec("UPDATE tags SET name = ? WHERE id = ?  ",tag.Name,tag.ID)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func NewTagRepository() ITagRepository {
	return &TagRepository{Conn: database.GetConnection()}
}
