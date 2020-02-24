package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"satya-labs/api"
	"satya-labs/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, _ := database.Initialize()

	port := os.Getenv("PORT")

	app := gin.Default()
	app.Use(database.Inject(db))
	api.ApplyRoutes(app)
	app.Run(":" + port)
}
