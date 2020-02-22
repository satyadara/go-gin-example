package main

import (
	"github.com/joho/godotenv"
	"satya-labs/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.Initialize()
}
