package main

import (
	"log"
	"os"

	"github.com/beeerlian/motor-rental/config"
	"github.com/beeerlian/motor-rental/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":     true,
			"message":     "You are at the root endpoint 😉",
			"github_repo": "https://github.com/beeerlian/motor-rental",
		})
	})
	app.Post("/api/customers/register/", controllers.CustomerRegistration)
	app.Post("/api/customers/login/email/", controllers.LoginWithEmail)
	app.Delete("/api/customers/:id", controllers.DeleteCustomer)
	app.Get("/api/customers/", controllers.GetAllCustomer)
	app.Post("/api/customers/rent/:motorId/:customerId", controllers.RentMotor)
	app.Get("/api/customers/:id", controllers.GetCustomer)

	app.Get("/api/motors/", controllers.GetAllMotors)
	app.Get("/api/motors/:id", controllers.GetMotor)
	app.Post("/api/motors/", controllers.AddMotor)
	app.Put("/api/motors/:id", controllers.UpdateMotor)
	app.Delete("/api/motors/:id", controllers.DeleteMotor)

	app.Get("/api/transactions/", controllers.GetAllTransactions)
	app.Post("/api/transactions/:motorId/:customerId", controllers.AddTransaction)
	app.Get("/api/transactions/:id", controllers.GetTransactionById)

}

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()

	setupRoutes(app)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
