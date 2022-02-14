package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
}

func AdminontrolRoute(app *fiber.App) {
	app.Post("/AdminControl", controllers.CreateAdmincontrol)
	app.Get("/admin/:adminId", controllers.GetAdmincontrol)
	app.Put("/admin/:adminId", controllers.EditAdmincontrol)
	app.Delete("/admin/:adminId", controllers.DeleteAdmincontrol)
	app.Get("/admindata", controllers.GetAlladmindata)
}
