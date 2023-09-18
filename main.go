package main

import (
	"fiber-gorm-restapi/database"
	"fiber-gorm-restapi/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New())

	router.Setuproutes(app)

	app.Listen(":3000")
}
