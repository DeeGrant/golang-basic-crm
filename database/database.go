package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DbConn *gorm.DB
)

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load environment variables.")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, dbName, port)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to database.")
	}
	fmt.Println("Connection opened to database")
	DbConn = conn
}

func GetDb() *gorm.DB {
	return DbConn
}
