package entities

type Post struct {
	ID          int `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Title       string `json:"title" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"type:text"`
	AuthorId    int `json:"author_id" gorm:"type:integer`
}