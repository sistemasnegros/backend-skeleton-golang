package authControllers

import (
	authDTO "backend-skeleton-golang/auth/app/dto"
	authService "backend-skeleton-golang/auth/app/services"
	logService "backend-skeleton-golang/commons/app/services/log-service"
	utilsFiberInfra "backend-skeleton-golang/commons/infra/fiber/utils"

	"github.com/gofiber/fiber/v2"
)

type IController interface {
	LoadRouter(fiber.Router)
	Register(*fiber.Ctx) error
	Login(*fiber.Ctx) error
}

type Controller struct {
	service *authService.Service
}

func New(service *authService.Service) IController {
	return &Controller{service: service}
}

func (controller Controller) LoadRouter(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Post("/forgot-password", controller.ForgotPassword)
	auth.Put("/restore-password/:token", controller.RestorePassword)
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
