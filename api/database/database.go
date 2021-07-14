package database

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm/schema"
	"time"

	//"log"
	//"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName string = "neo"
	Password string = "neoneo"
	Addr     string = "127.0.0.1"
	Port     int = 3306
	Database string = "nesplash"
	MaxLifetime int = 10
	MaxOpenConns int = 100
	MaxIdleConns int = 10
)

func GetDataBase() *gorm.DB {
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", UserName, Password, Addr, Port, Database)
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("connection to mysql failed : ", err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(MaxLifetime)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}