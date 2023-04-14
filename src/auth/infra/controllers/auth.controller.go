package authControllers

import (
	authDTO "backend-skeleton-golang/auth/app/dto"
	authService "backend-skeleton-golang/auth/app/services"
	logService "backend-skeleton-golang/commons/app/services/log-service"
	middlewareInfra "backend-skeleton-golang/commons/infra/fiber/middleware"
	utilsFiberInfra "backend-skeleton-golang/commons/infra/fiber/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type IController interface {
	LoadRouter(fiber.Router)
	Register(*fiber.Ctx) error
	Login(*fiber.Ctx) error
}

type Controller struct {
	service    authService.IService
	middleware *middlewareInfra.HandlerMiddleware
}

func New(service authService.IService, middleware *middlewareInfra.HandlerMiddleware) IController {
	return &Controller{service: service, middleware: middleware}
}

func (controller Controller) LoadRouter(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Post("/forgot-password", controller.ForgotPassword)
	auth.Put("/restore-password/:token", controller.RestorePassword)

	auth.Get("/me", controller.middleware.Protected(), controller.GetMe)
	auth.Patch("/me", controller.middleware.Protected(), controller.UpdateMe)

}

func (controller Controller) Register(c *fiber.Ctx) error {
	body := new(authDTO.Register)

	if err := c.BodyParser(body); err != nil {
		logService.Error(err.Error())
		return err
	}

	errors := utilsFiberInfra.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	code, res := controller.service.Register(body)

	return c.Status(code).JSON(res)
}

func (controller Controller) Login(c *fiber.Ctx) error {
	body := new(authDTO.Login)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	errors := utilsFiberInfra.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	code, res := controller.service.Login(body)

	return c.Status(code).JSON(res)
}

func (controller Controller) ForgotPassword(c *fiber.Ctx) error {
	body := new(authDTO.ForgotPassword)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	errors := utilsFiberInfra.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	code, res := controller.service.ForgotPassword(body)

	return c.Status(code).JSON(res)
}

func (controller Controller) RestorePassword(c *fiber.Ctx) error {
	body := new(authDTO.RestorePassword)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	errors := utilsFiberInfra.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	code, res := controller.service.RestorePassword(c.Params("token"), body)

	return c.Status(code).JSON(res)
}

func (controller Controller) GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	code, res := controller.service.GetMe(claims["id"].(string))

	return c.Status(code).JSON(res)
}

func (controller Controller) UpdateMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	body := new(authDTO.UpdateMe)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	errors := utilsFiberInfra.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	code, res := controller.service.UpdateMe(claims["id"].(string), body)

	return c.Status(code).JSON(res)
}
