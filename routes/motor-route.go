package routes

import (
	"github.com/beeerlian/motor-rental/controllers" // replace
	"github.com/gofiber/fiber/v2"
)

func MotorsRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllMotors)
	route.Get("/:id", controllers.GetMotor)
	route.Post("/", controllers.AddMotor)
	route.Put("/:id", controllers.UpdateMotor)
	route.Delete("/:id", controllers.DeleteMotor)
}
