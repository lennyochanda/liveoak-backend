package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/lennyochanda/LiveOak/user"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not find env variables!\n")
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	defer db.Close()

	fmt.Print("Connected to the database!\n")

	userRepo := user.NewMySQLUserRepository(db)
	userService := user.NewUserService(userRepo)
	user.SetUpUserRoutes(userService)
}
