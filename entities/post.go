package entities

type Post struct {
	ID          int `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Title       int `json:"title" gorm:"type:varchar(100)"`
	Description int `json:"description" gorm:"type:text"`
	AuthorId    int `json:"author_id" gorm:"type:integer`
}