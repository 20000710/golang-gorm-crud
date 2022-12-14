package config

import (
	"fmt"
	"os"
	"rakamin/structs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func DBInit() *gorm.DB {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=godb password=pgadmin123")
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	db = conn

	db.Debug().AutoMigrate(structs.Person{})

	return db

}
