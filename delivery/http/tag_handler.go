package http

import (
	"encoding/json"
	"go-crash-course/entities"
	"go-crash-course/services"
	"go-crash-course/utils/lib"
	"strconv"

	"github.com/gofiber/fiber"
)

type TagHandler struct {
	TagService *services.TagService
}

func (p *TagHandler) FetchTag(c *fiber.Ctx) {
	response := &lib.Response{}
	data := p.TagService.FetchTag()
	response.ResponseOK("FETCH_TAG_SUCCESS", data, c)
}

func (p *TagHandler) NewTag(c *fiber.Ctx) {
	response := &lib.Response{}
	body := c.Body()
	var tag *entities.Tag
	var _ = json.Unmarshal([]byte(body), &tag)

	result := p.TagService.SaveTag(tag)

	response.ResponseOK("CREATE_TAG_SUCCESS", result, c)
}

func (p *TagHandler) FechTagById(c *fiber.Ctx) {
	response := &lib.Response{}
	paramId := c.Params("id")
	id, _ := strconv.Atoi(paramId)
	data := p.TagService.FetchTagById(id)
	response.ResponseOK("FETCH_TAG_SUCCESS", data, c)
}

func (p *TagHandler) DeleteTag(c *fiber.Ctx) {
	response := &lib.Response{}
	param := c.Params("id")
	id, _ := strconv.Atoi(param)

	err := p.TagService.Destroy(id)

	if err != nil {
		response.ResponseNOK("DELETE_TAG_FAILED", err, c)
	}

	response.ResponseOK("DELETE_TAG_SUCCESS", nil, c)

}

func (p *TagHandler) UpdateTag(c *fiber.Ctx) {
	response := lib.Response{}
	data := c.Body()
	var tag *entities.Tag
	json.Unmarshal([]byte(data), &tag)

	result, err := p.TagService.Update(tag)

	if err != nil {
		response.ResponseNOK("UPDATE_TAG_FAILED", err, c)
	}
	response.ResponseOK("UPDATE_TAG_SUCCES", result, c)
}

func NewTagHandler(r *fiber.App) {
	handler := &TagHandler{
		TagService: services.NewTagService(),
	}

	//r.Get("/posts", handler.FetchPost)
	r.Get("/tags", handler.FetchTag)
	r.Get("/tag/:id", handler.FechTagById)
	r.Post("/tag", handler.NewTag)
	r.Delete("/tag/:id", handler.DeleteTag)
	r.Put("/tag", handler.UpdateTag)
}
