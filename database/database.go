package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var conn *sql.DB
var err error

func Connection() {
	databaseURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
	)

	conn, err = sql.Open("mysql", databaseURI)
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

func GetConnection() *sql.DB {
	return conn
}

func CreateMigration() {
	fmt.Println("MIGRATION => `migration successfully`")
}
