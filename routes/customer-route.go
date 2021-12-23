package routes

import (
	"github.com/beeerlian/motor-rental/controllers" // replace
	"github.com/gofiber/fiber/v2"
)

func CustomersRoute(route fiber.Router) {
	route.Post("/register/", controllers.CustomerRegistration)
	route.Post("/login/email/", controllers.LoginWithEmail)
	route.Delete("/:id", controllers.DeleteCustomer)
	route.Get("/", controllers.GetAllCustomer)
	route.Post("/participant/:motorId/:customerId", controllers.RentMotor)
	route.Get("/:id", controllers.GetCustomer)
}
