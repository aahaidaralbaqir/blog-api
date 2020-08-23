package http

import (
	"github.com/gofiber/fiber"
	"go-crash-course/services"
	"go-crash-course/utils/lib"
	"strconv"
)

type AuthorHandler struct {
	AuthorService *services.AuthorService
}

func (p *AuthorHandler) FechAuthor(c *fiber.Ctx) {
	response := &lib.Response{}
	data := p.AuthorService.FetchAuthor()
	response.ResponseOK("FETCH_AUTHOR_SUCCESS",data,c)
}


func (p *AuthorHandler) FechAuthorById(c *fiber.Ctx) {
	response := &lib.Response{}
	paramId := c.Params("id")
	id,_ := strconv.Atoi(paramId)
	data := p.AuthorService.FetchAuthorById(id)
	response.ResponseOK("FETCH_AUTHOR_SUCCESS",data,c)
}


func NewAuthorHandler(r *fiber.App) {
	handler := &AuthorHandler{
		AuthorService: services.NewAuthorService(),
	}

	//r.Get("/posts", handler.FetchPost)
	r.Get("/authors", handler.FechAuthor)
	r.Get("/author/:id", handler.FechAuthorById)
}
