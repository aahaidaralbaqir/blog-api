package cmd

import (
	"fmt"
	"go-crash-course/database"
	"go-crash-course/delivery/http"

	"github.com/gofiber/fiber"
	"github.com/spf13/viper"
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

func (a *Application) ConfigureRoutes(app *fiber.App) {
	http.NewPostHandler(app)
	http.NewAuthorHandler(app)
}

func (a *Application) Start(port int) {
	app := fiber.New()
	a.ConfigureRoutes(app)
	app.Listen(port)
}
