package main

import (
	"et-practice/config"
	"et-practice/router"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	app := fiber.New()

	//run database
	config.ConnectDB()

	// router
	router.UserRoute(app)

	apiPort := os.Getenv("API_PORT")
	app.Listen(":" + apiPort)

}
