package fiberInfra

import (
	authControllers "backend-skeleton-golang/auth/infra/controllers"
	usersControllers "backend-skeleton-golang/users/infra/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type API struct {
	app             *fiber.App
	authController  authControllers.IController
	usersController usersControllers.IController
}

func New(authController authControllers.IController, usersController usersControllers.IController) *API {
	return &API{
		authController:  authController,
		usersController: usersController,
	}
}

func HandlerRoot(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func (a *API) Start() error {
	app := fiber.New(fiber.Config{
		StrictRouting: false,
	})

	app.Use(cors.New())

	a.app = app

	a.LoadRouter()

	return app.Listen(":3000")

}

func (a *API) LoadRouter() {

	a.app.Get("/", HandlerRoot)

	router := a.app.Group("/api")
	a.authController.LoadRouter(router)
	a.usersController.LoadRouter(router)

}
