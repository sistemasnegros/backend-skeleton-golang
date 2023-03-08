package authService

import (
	authDTO "backend-skeleton-golang/auth/app/dto"
	configService "backend-skeleton-golang/commons/app/services/config-service"
	resService "backend-skeleton-golang/commons/app/services/http-service"
	logService "backend-skeleton-golang/commons/app/services/log-service"
	smtpService "backend-skeleton-golang/commons/app/services/smtp-service"
	smtpDomain "backend-skeleton-golang/commons/domain/smtp"
	usersDomain "backend-skeleton-golang/users/domain"
	usersRepo "backend-skeleton-golang/users/infra/repo"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Register(body *authDTO.Register) (int, interface{})
	Login(body *authDTO.Login) (int, interface{})
}

type Service struct {
	repo *usersRepo.Users
	smtp smtpService.ISmtpService
}

func New(repo *usersRepo.Users, smtp smtpService.ISmtpService) *Service {
	return &Service{repo: repo, smtp: smtp}
}

func (s *Service) Register(body *authDTO.Register) (int, interface{}) {
	userIdFound, _ := s.repo.FindById(body.Id)

	if userIdFound.Id != "" {
		return resService.BadRequest("id already exists")
	}

	querySearchUser := map[string]interface{}{"email": body.Email}
	userEmailFound, err := s.repo.FindOne(querySearchUser)

	if userEmailFound.Id != "" {
		return resService.BadRequest("email already exists")
	}

	userDomain := usersDomain.User{}

	copier.Copy(&userDomain, &body)

	user, err := s.repo.Create(userDomain)

	if err != nil {
		logService.Error(err.Error())
		return resService.InternalServerError("err unknown")
	}

	return resService.Created(user)
}

func (s *Service) Login(body *authDTO.Login) (int, interface{}) {

	querySearchUser := map[string]interface{}{"email": body.Email}
	user, err := s.repo.FindOne(querySearchUser)

	if user.Id == "" {
		return resService.Unauthorized("user or password not found")
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if errPassword != nil {
		return resService.Unauthorized("user or password not found")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(configService.GetJwtExt())).Unix()

	tokenSingned, err := token.SignedString([]byte(configService.GetJwtSecret()))
	if err != nil {
		logService.Error(err.Error())
		return resService.InternalServerError("error singing token")
	}

	userRes := authDTO.LoginResUser{}
	copier.Copy(&userRes, &user)

	res := authDTO.LoginRes{
		Token: tokenSingned,
		User:  userRes,
	}

	return resService.Ok(res)
}

func (s *Service) ForgotPassword(body *authDTO.ForgotPassword) (int, interface{}) {

	querySearchUser := map[string]interface{}{"email": body.Email}
	user, err := s.repo.FindOne(querySearchUser)

	if user.Id == "" {
		return resService.Ok("password reset email was successfully sent")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(60)).Unix()

	tokenSingned, err := token.SignedString([]byte(configService.GetJwtSecret()))
	if err != nil {
		logService.Error(err.Error())
		return resService.InternalServerError("error singing token")
	}

	email := smtpDomain.SendArgs{
		To:       user.Email,
		Subject:  "Restore password!",
		Template: "notify.html",
		Data: smtpDomain.EmailTemplateDefault{
			FullName:      user.FirstName + " " + user.LastName,
			Message:       "You have received this email to set your password, click on the button to go to set it.",
			ButtonMessage: "Go to set password!",
			URL:           "http://localhost:3000/auth/restore-password/" + tokenSingned,
		},
	}

	logService.Info("token for password reset: " + tokenSingned)
	errSmtp := s.smtp.Send(email)

	if errSmtp != nil {
		resService.InternalServerError("error in smtp server")
	}

	return resService.Ok("password reset email was successfully sent")
}

func (s *Service) RestorePassword(tokenString string, body *authDTO.RestorePassword) (int, interface{}) {

	type TokenStruct struct {
		jwt.RegisteredClaims
		Id string
	}

	var tokenDecoded TokenStruct

	token, err := jwt.ParseWithClaims(tokenString, &tokenDecoded, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configService.GetJwtSecret()), nil
	})

	if err != nil {
		logService.Error(err.Error())
		return resService.BadRequest("invalid token")
	}

	if !token.Valid {
		return resService.BadRequest("invalid token")
	}

	user, err := s.repo.FindById(tokenDecoded.Id)

	if user.Id == "" {
		return resService.BadRequest("invalid token")
	}

	s.repo.UpdateById(tokenDecoded.Id, &usersDomain.User{Password: body.Password})

	return resService.Ok("password updated successfully")
}
