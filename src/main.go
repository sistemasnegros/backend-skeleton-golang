package main

import (
	authService "backend-skeleton-golang/auth/app/services"
	authControllers "backend-skeleton-golang/auth/infra/controllers"
	smtpService "backend-skeleton-golang/commons/app/services/smtp-service"
	msgDomain "backend-skeleton-golang/commons/domain/msg"
	fiberInfra "backend-skeleton-golang/commons/infra/fiber/fiber"
	middlewareInfra "backend-skeleton-golang/commons/infra/fiber/middleware"
	godotenvInfra "backend-skeleton-golang/commons/infra/godotenv"
	gomailInfra "backend-skeleton-golang/commons/infra/gomail"
	logrusInfra "backend-skeleton-golang/commons/infra/logrus"
	mongodbInfra "backend-skeleton-golang/commons/infra/mongodb"
	usersService "backend-skeleton-golang/users/app/services"
	usersController "backend-skeleton-golang/users/infra/controllers"
	usersRepoMongo "backend-skeleton-golang/users/infra/mongodb/repo"
	"context"

	"go.uber.org/fx"
)

func main() {

	godotenvInfra.Load()
	logrusInfra.Init()
	msgDomain.New()

	app := fx.New(fx.Provide(

		mongodbInfra.New,

		usersRepoMongo.New,

		gomailInfra.New,
		smtpService.New,

		middlewareInfra.New,

		authService.New,
		authControllers.New,

		usersController.New,
		usersService.New,

		fiberInfra.New,
	),
		fx.Invoke(
			start,
		),
	)

	app.Run()
}

func start(lc fx.Lifecycle, api *fiberInfra.API) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			go api.Start()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
