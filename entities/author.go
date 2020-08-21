package entities


type Author struct {
	ID int `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Name string `json:"name" gorm:"type:varchar(100)"`
}