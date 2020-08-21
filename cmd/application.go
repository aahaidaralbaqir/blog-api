package cmd

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/spf13/viper"
	"go-crash-course/database"
)

type Application struct {}


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

func (a *Application) Start(port int)  {
	app := fiber.New()
	app.Listen(port)
}