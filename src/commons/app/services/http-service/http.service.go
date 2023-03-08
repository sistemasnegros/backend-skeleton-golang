package resService

import (
	httpInfra "backend-skeleton-golang/commons/infra/http"
)

func Ok(msg interface{}) (int, interface{}) {
	return httpInfra.OkHttp(msg)
}

func Created(res interface{}) (int, interface{}) {
	return httpInfra.CreatedHttp(res)
}

func InternalServerError(msg string) (int, string) {
	return httpInfra.InternalServerErrorHttp(msg)
}

func BadRequest(msg string) (int, string) {
	return httpInfra.BadRequestHttp(msg)
}

func Unauthorized(msg string) (int, string) {
	return httpInfra.UnauthorizedHttp(msg)
}

func Forbidden(msg string) (int, string) {
	return httpInfra.ForbiddenHttp(msg)
}

func NotFound(msg string) (int, string) {
	return httpInfra.NotFoundHttp(msg)
}
