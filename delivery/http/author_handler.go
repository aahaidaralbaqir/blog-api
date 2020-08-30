package http

import (
	"encoding/json"
	"github.com/gofiber/fiber"
	"go-crash-course/entities"
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

func (p *AuthorHandler) NewAuthor(c *fiber.Ctx) {
	response := &lib.Response{}
	body := c.Body()
	var author *entities.Author
	var _ = json.Unmarshal([]byte(body), &author)

	result := p.AuthorService.SaveAuthor(author)

	response.ResponseOK("CREATE_AUTHOR_SUCCESS",result,c)
}


func (p *AuthorHandler) FechAuthorById(c *fiber.Ctx) {
	response := &lib.Response{}
	paramId := c.Params("id")
	id,_ := strconv.Atoi(paramId)
	data := p.AuthorService.FetchAuthorById(id)
	response.ResponseOK("FETCH_AUTHOR_SUCCESS",data,c)
}

func (p *AuthorHandler) DeleteAuthor (c *fiber.Ctx) {
	response := &lib.Response{}
	param := c.Params("id")
	id,_ := strconv.Atoi(param)

	err := p.AuthorService.Destroy(id)

	if err != nil {
		response.ResponseNOK("DELETE_AUTHOR_FAILED",err,c)
	}

	response.ResponseOK("DELETE_AUTHOR_SUCCESS",nil,c)

}

func (p *AuthorHandler) UpdateAuthor (c *fiber.Ctx) {
	response := lib.Response{}
	data := c.Body()
	var author *entities.Author
	json.Unmarshal([]byte(data),&author)

	result, err := p.AuthorService.Update(author)

	if err != nil {
		response.ResponseNOK("UPDATE_AUTHOR_FAILED",err,c)
	}
	response.ResponseOK("UPDATE_AUTHOR_SUCCES",result,c)
}


func NewAuthorHandler(r *fiber.App) {
	handler := &AuthorHandler{
		AuthorService: services.NewAuthorService(),
	}

	//r.Get("/posts", handler.FetchPost)
	r.Get("/authors", handler.FechAuthor)
	r.Get("/author/:id", handler.FechAuthorById)
	r.Post("/author", handler.NewAuthor)
	r.Delete("/author/:id",handler.DeleteAuthor)
	r.Put("/author",handler.UpdateAuthor)
}
