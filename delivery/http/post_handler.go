package http

import (
	"github.com/gofiber/fiber"
	"go-crash-course/entities"
	"go-crash-course/services"
	"go-crash-course/utils/lib"
)

type PostHandler struct {
	PostService *services.PostService
}

func (p *PostHandler) FetchPost(c *fiber.Ctx) {
	response := &lib.Response{}
	data := p.PostService.GetPost()
	response.ResponseOK("RECEIVE_FETCH_POST_SUCCESS",data,c)
}

func (p *PostHandler) FetchPostWithAuthor(c *fiber.Ctx) {
	response := &lib.Response{}
	data := p.PostService.GetPostWithAuthor()
	response.ResponseOK("RECEIVE_FETCH_POST_AUTHOR_SUCCESS",data,c)
}

func (p *PostHandler) NewPost(c *fiber.Ctx) {
	response := &lib.Response{}
	post := new(entities.Post)
	c.BodyParser(post)

	result, err := p.PostService.SavePost(post)

	if err != nil {
		response.ResponseNOK("RECEIVE_CREATE_POST_ERROR",err.Error(),c)
	}

	response.ResponseOK("RECEIVE_CREATE_POST_SUCCESS",result,c)
}

func NewPostHandler(r *fiber.App) {
	handler := &PostHandler{
		PostService: services.NewPostService(),
	}

	//r.Get("/posts", handler.FetchPost)
	r.Get("/posts", handler.FetchPostWithAuthor)
	r.Post("/post",handler.NewPost)

}
