package router

import (
	"fiber-gorm-restapi/handler"

	"github.com/gofiber/fiber/v2"
)

// setuproutes func
func Setuproutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	//routes
	v1.Get("/", handler.GetAllUsers)
	v1.Get("/:id", handler.GetSingleUser)
	v1.Post("/", handler.CreateUser)
	v1.Delete("/:id", handler.DeleteUser)
	v1.Put("/:id", handler.UpdateUser)
}
