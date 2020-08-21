package http

import (
	"github.com/gofiber/fiber"
	"go-crash-course/services"
	"net/http"
)

type ResponseError struct {
	message string `json:"message"`
}

type PostHandler struct {
	PostService services.PostService
}

func (p *PostHandler) FetchPost(c *fiber.Ctx) {
	post := p.PostService.GetPost()

	c.Status(http.StatusOK).JSON(post)
}

func NewPostHandler(r *fiber.App) {
	handler := &PostHandler{
		PostService: services.NewPostService(),
	}

	r.Get("/posts",handler.FetchPost)

}

