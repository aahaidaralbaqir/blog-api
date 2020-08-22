package	lib

import (
	"github.com/gofiber/fiber"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data", omitempty`
	StatusCode int `json:"statusCode", omitempty`
	Errors interface{} `json:"errors",omitempty`
	Flag string `json:"flag",omitempty`
}

func (r *Response) ResponseOK(flag string,data interface{}, c *fiber.Ctx) {
	r.Data = data
	r.StatusCode = http.StatusOK
	r.Flag = flag
	c.Status(http.StatusOK).JSON(r)
}

func (r *Response) ResponseNOK(flag string,errors interface{}, c *fiber.Ctx) {
	r.Errors = errors
	r.Flag = flag
	r.StatusCode = http.StatusInternalServerError
	c.Status(http.StatusInternalServerError).JSON(r)

}