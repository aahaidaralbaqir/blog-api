package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json="id"`
	Title string `json="title"`
	Text  string `json="text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{
		1,
		"Title 1",
		"Text One",
	}}
}

func getPosts(r http.ResponseWriter, w *http.Request) {
	r.Header().Set("Content-type", "applicaton/json")
	result, err := json.Marshal(posts)

	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{"error": "Something went wrong","flag" : "MARSHAL_ERROR"}`))
		return
	}
	r.WriteHeader(http.StatusOK)
	r.Write(result)
}

func addPost(r http.ResponseWriter, w *http.Request) {
	var post Post
	err := json.NewDecoder(w.Body).Decode(&post)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{"error": "Something went wrong","flag" : "UNMARSHAL_ERROR"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	result, err := json.Marshal(post)
	r.WriteHeader(http.StatusOK)
	r.Write(result)
}
