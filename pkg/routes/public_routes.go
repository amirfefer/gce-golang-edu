package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/computes", controllers.GetComputes)   // get list of all computes
	route.Get("/compute/:id/list", controllers.GetInstancesList) // get all compute instances
    route.Get("/zones/project/:id", controllers.GetZones)

	// Routes for POST method:
	route.Post("/compute", controllers.CreateCompute) // create new compute
	route.Post("/instance", controllers.CreateInstance) // create new instance
	route.Post("/compute/:id/instance/:name", controllers.Operate) // create new instance
	route.Post("/user/sign/up", controllers.UserSignUp) // register a new user
	route.Post("/user/sign/in", controllers.UserSignIn) // auth, return Access & Refresh tokens
}
