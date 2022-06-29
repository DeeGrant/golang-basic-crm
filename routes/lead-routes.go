package routes

import (
	"github.com/DeeGrant/golang-basic-crm/controllers"
	"github.com/gofiber/fiber"
)

var RegisterRoutes = func(app *fiber.App) {
	app.Get("/leads", controllers.GetLeads)
	app.Get("/lead/:id", controllers.GetLead)
	app.Post("/lead", controllers.CreateLead)
	app.Put("/lead/:id", controllers.UpdateLead)
	app.Delete("/lead/:id", controllers.DeleteLead)
}
