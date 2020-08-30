package cmd

import (
	"fmt"
	"github.com/gofiber/fiber"
	_ "github.com/gofiber/fiber/middleware"
	"github.com/spf13/viper"
	"go-crash-course/database"
	"go-crash-course/delivery/http"
	"go-crash-course/delivery/middleware"
)

type Application struct{}

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Service run in DEBUG Mode")
	}

}

func (a *Application) ConfigureDatabase() {
	database.CreateConnection()
}

func (a *Application) ConfigureMigration() {
	database.CreateMigration()
}

func (a *Application) ConfigureMiddleware(app *fiber.App){
	middleware.SetupCors(app)
}

func (a *Application) ConfigureRoutes(app *fiber.App) {
	http.NewPostHandler(app)
	http.NewAuthorHandler(app)
}

func (a *Application) Start(port int) {
	app := fiber.New()
	a.ConfigureRoutes(app)
	a.ConfigureMiddleware(app)
	app.Listen(port)
}
