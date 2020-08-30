package http

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"go-crash-course/entities"
	"go-crash-course/services"
	"go-crash-course/utils/lib"
	"strconv"
)

type PostHandler struct {
	PostService *services.PostService
}

func (p *PostHandler) FetchPost(c *fiber.Ctx) {
	response := &lib.Response{}
	data := p.PostService.GetPost()
	response.ResponseOK("RECEIVE_FETCH_POST_SUCCESS",data,c)
}


func (p *PostHandler) FetchPostById(c *fiber.Ctx) {
	response := &lib.Response{}
	param := c.Params("id")
	id,_ := strconv.Atoi(param)
	data := p.PostService.FindById(id)
	response.ResponseOK("FETCH_POST_SINGLE_SUCCESS",data,c)
}

func (p *PostHandler) FetchPostWithAuthor(c *fiber.Ctx) {
	response := &lib.Response{}
	data := p.PostService.GetPostWithAuthor()
	response.ResponseOK("RECEIVE_FETCH_POST_AUTHOR_SUCCESS",data,c)
}

func (p *PostHandler) NewPost(c *fiber.Ctx) {
	response := &lib.Response{}
	body := c.Body()
	var post *entities.Post
	var _ = json.Unmarshal([]byte(body), &post)

	result := p.PostService.SavePost(post)

	response.ResponseOK("CREATE_POST_SUCCESS",result,c)
}

func (p *PostHandler) DeletePost (c *fiber.Ctx) {
	response := &lib.Response{}
	param := c.Params("id")
	id,_ := strconv.Atoi(param)

	err := p.PostService.Destroy(id)

	if err != nil {
		response.ResponseNOK("DELETE_POST_FAILED",err,c)
	}

	response.ResponseOK("DELETE_POST_SUCCESS",nil,c)

}


func (p *PostHandler) UpdatePost (c *fiber.Ctx) {
	response := lib.Response{}
	data := c.Body()
	var post *entities.Post
	json.Unmarshal([]byte(data),&post)

	result, err := p.PostService.Update(post)

	if err != nil {
		response.ResponseNOK("UPDATE_POST_FAILED",err,c)
	}
	response.ResponseOK("UPDATE_POST_SUCCES",result,c)
}

func NewPostHandler(r *fiber.App) {
	handler := &PostHandler{
		PostService: services.NewPostService(),
	}

	//r.Get("/posts", handler.FetchPost)
	r.Get("/posts", handler.FetchPostWithAuthor)
	r.Get("/post/:id", handler.FetchPostById)
	r.Delete("/post/:id", handler.DeletePost)
	r.Post("/post",handler.NewPost)
	r.Put("/post",handler.UpdatePost)


}
