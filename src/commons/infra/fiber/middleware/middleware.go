package middlewareInfra

import (
	configService "backend-skeleton-golang/commons/app/services/config-service"
	usersRepoMongo "backend-skeleton-golang/users/infra/mongodb/repo"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type HandlerMiddleware struct {
	repo usersRepoMongo.IUsers
}

func New(repo usersRepoMongo.IUsers) *HandlerMiddleware {

	return &HandlerMiddleware{repo: repo}
}

func (h HandlerMiddleware) Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(configService.GetJwtSecret()),
		ErrorHandler:   h.jwtError,
		SuccessHandler: h.successHandler,
	})
}

func (h HandlerMiddleware) jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func (h HandlerMiddleware) successHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := claims["id"].(string)
	userModel, _ := h.repo.FindById(id)

	if userModel.Id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON("invalid token")
	}

	return c.Next()
}
