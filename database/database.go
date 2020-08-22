package database

import (
	"fmt"
	"go-crash-course/entities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var conn *gorm.DB
var err error

func Connection() {
	databaseURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
	)

	conn, err = gorm.Open("mysql", databaseURI)
	fmt.Println(databaseURI)
	if err != nil {
		fmt.Println("DATABASE => `Error When Connecting to database`", err.Error())
	}
	fmt.Println("DATABASE => `Success connecting to database`")
}

func CreateConnection() {
	if err != nil {
		panic("CREATE CONNECTION => Error While Create Connection")
	}
	Connection()
}

func GetConnection() *gorm.DB {
	return conn
}

func CreateMigration() {
	conn.AutoMigrate(&entities.Post{}, &entities.Author{})
	fmt.Println("MIGRATION => `migration successfully`")
}
