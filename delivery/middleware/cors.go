package middleware

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func SetupCors(app *fiber.App){
	config := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH"},
		AllowCredentials: false,
		MaxAge: 10,
	}
	app.Use(cors.New(config))
}