package routes

import (
	"spotifyAPI/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	app := fiber.New()
	app.Post("/api/v1/login", controllers.Login)
	app.Post("/api/v1/register", controllers.Register)
	app.Get("/api/v1/user", controllers.GetUser)
	app.Put("/api/v1/user", controllers.UpdateUser)
	app.Delete("/api/v1/user", controllers.DeleteUser)
	app.Get("/api/v1/admin/user", controllers.GetAllUsers)
	app.Get("/api/v1/admin/user/:id", controllers.GetUserById)
	app.Put("/api/v1/admin/user/:id", controllers.UpdateUserById)
	app.Delete("/api/v1/admin/user/:id", controllers.DeleteUserById)
	app.Get("/api/v1/getPremium", controllers.GetPremium)
	app.Listen(":8080")
}
