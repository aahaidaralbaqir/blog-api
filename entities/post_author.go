package entities


type PostAuthor struct {
	Post
	Author *Author `json:"author"`
}