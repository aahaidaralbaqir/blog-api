package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DBConn *gorm.DB
)

type Database struct{
	Dsn string
}

func (d *Database) SetDSN(dsn string) {
	d.Dsn = dsn
}

func (d *Database) Setup() *gorm.DB {
	var err error
	DBConn, err = gorm.Open("mysql", d.Dsn)

	if err != nil {
		panic("Database -> error")
	}
	fmt.Println("Database connection successfull opened")
	return DBConn
}


func (d *Database) GetInstance() *gorm.DB {
	if DBConn == nil {
		if d.Dsn == "" {
			panic("please configure the dsn")
		}
		d.Setup()
	}
	return DBConn
}