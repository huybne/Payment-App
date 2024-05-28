package Router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huybuine/Payment-App/Handlers"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/payment")
	// routes
	v1.Get("/", Handlers.GetPaymentDetails)
	v1.Get("/:id", Handlers.GetPaymentDetail)
	v1.Post("/", Handlers.CreatePaymentDetail)
	v1.Put("/:id", Handlers.UpdatePaymentDetail)
	v1.Delete("/:id", Handlers.DeletePaymentDetail)
}
