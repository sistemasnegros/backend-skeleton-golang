package httpInfra

func OkHttp(data interface{}) (int, interface{}) {
	return 200, data
}

func CreatedHttp(data interface{}) (int, interface{}) {
	return 201, data
}

func InternalServerErrorHttp(msg string) (int, string) {
	return 500, msg
}

func BadRequestHttp(msg string) (int, string) {
	return 400, msg
}

func UnauthorizedHttp(msg string) (int, string) {
	return 401, msg
}

func ForbiddenHttp(msg string) (int, string) {
	return 403, msg
}

func NotFoundHttp(msg string) (int, string) {
	return 404, msg
}
