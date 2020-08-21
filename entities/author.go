package entities

import (
	"github.com/jinzhu/gorm"
)

type Author struct {
	*gorm.Model
	Id int `json:"id"`
	Name string `json:"name"`
	Age string `json:"age"`
}