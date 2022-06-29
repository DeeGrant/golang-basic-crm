package main

import (
	"fmt"
	"github.com/DeeGrant/golang-basic-crm/database"
	"github.com/DeeGrant/golang-basic-crm/routes"
	"github.com/gofiber/fiber"
)

func main() {
	fmt.Println("Initializing App ...")
	app := fiber.New()
	routes.RegisterRoutes(app)

	database.Connect()

	err := app.Listen(3000)
	if err != nil {
		panic("App failed to initialize.")
	}
	fmt.Println("App started on port 3000")
}
