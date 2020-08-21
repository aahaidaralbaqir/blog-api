package cmd

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go-crash-course/database"
)

type Application struct {}

func(a *Application) configureDatabase() *gorm.DB {
	HOST := viper.GetString("database.host")
	PORT := viper.GetString("database.port")
	USER := viper.GetString("database.user")
	PASSWORD := viper.GetString("database.password"	)
	DBNAME := viper.GetString("database.dbname")
	DSN := fmt.Sprintf("%s:%s@(%s:%s)/%s",USER,PASSWORD,HOST,PORT,DBNAME)
	connection := new(database.Database)
	connection.SetDSN(DSN)
	return connection.GetInstance()
}

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

func (a *Application) Start(port int)  {
	app := fiber.New()
	_ := a.configureDatabase()
	app.Listen(port)
}