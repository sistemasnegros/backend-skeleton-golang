package usersController

import (
	logService "backend-skeleton-golang/commons/app/services/log-service"
	serviceDomain "backend-skeleton-golang/commons/domain/service"
	middlewareInfra "backend-skeleton-golang/commons/infra/fiber/middleware"
	utilsFiberInfra "backend-skeleton-golang/commons/infra/fiber/utils"
	usersDTO "backend-skeleton-golang/users/app/dto"
	usersService "backend-skeleton-golang/users/app/services"

	"github.com/gofiber/fiber/v2"
)

type IController interface {
	LoadRouter(fiber.Router)
	Create(*fiber.Ctx) error
	Find(*fiber.Ctx) error
	FindById(*fiber.Ctx) error
	DeleteById(*fiber.Ctx) error
	UpdateById(*fiber.Ctx) error
}

type Controller struct {
	service    usersService.IService
	middleware *middlewareInfra.HandlerMiddleware
}

func New(service usersService.IService, middleware *middlewareInfra.HandlerMiddleware) IController {
	return &Controller{service: service, middleware: middleware}
}

func (controller Controller) LoadRouter(router fiber.Router) {
	routerGroup := router.Group("/users")
	routerGroup.Get("/", controller.middleware.Protected(), controller.Find)
	routerGroup.Get("/:id", controller.middleware.Protected(), controller.FindById)
	routerGroup.Delete("/:id", controller.middleware.Protected(), controller.DeleteById)
	routerGroup.Post("/", controller.middleware.Protected(), controller.Create)
	routerGroup.Put("/:id", controller.middleware.Protected(), controller.UpdateById)

}

func (controller Controller) Find(c *fiber.Ctx) error {
	paginationOpts := &serviceDomain.PaginationOpts{
		Page: c.QueryInt("page", 1),
	}

	code, res := controller.service.Find(paginationOpts)
	return c.Status(code).JSON(res)
}

func (controller Controller) FindById(c *fiber.Ctx) error {
	code, res := controller.service.FindById(c.Params("id"))
	return c.Status(code).JSON(res)
}

func (controller Controller) DeleteById(c *fiber.Ctx) error {
	code, res := controller.service.DeleteById(c.Params("id"))
	return c.Status(code).JSON(res)
}

func (controller Controller) Create(c *fiber.Ctx) error {
	body := new(usersDTO.Create)

	if err := c.BodyParser(body); err != nil {
		logService.Error(err.Error())
		return err
	}

	errors := utilsFiberInfra.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	code, res := controller.service.Create(body)

	return c.Status(code).JSON(res)
}

func (controller Controller) UpdateById(c *fiber.Ctx) error {
	body := new(usersDTO.Update)

	if err := c.BodyParser(body); err != nil {
		logService.Error(err.Error())
		return err
	}

	errors := utilsFiberInfra.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	code, res := controller.service.UpdateById(c.Params("id"), body)

	return c.Status(code).JSON(res)
}
