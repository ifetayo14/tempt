package main

import (
	"agit/database"
	"agit/router"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")

	database.StartDB()

	r := router.StartApp()
	err := r.Run(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}

}
