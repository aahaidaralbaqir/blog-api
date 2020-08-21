package entities

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	*gorm.Model
	Id          int `json:"id"`
	Title       int `json:"title"`
	Description int `json:"description"`
	Author      int `json:"author"`
}