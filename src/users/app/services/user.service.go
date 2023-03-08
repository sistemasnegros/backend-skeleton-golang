package usersService

import (
	resService "backend-skeleton-golang/commons/app/services/http-service"
	logService "backend-skeleton-golang/commons/app/services/log-service"
	msgDomain "backend-skeleton-golang/commons/domain/msg"
	usersDTO "backend-skeleton-golang/users/app/dto"
	usersDomain "backend-skeleton-golang/users/domain"
	usersRepo "backend-skeleton-golang/users/infra/repo"

	"github.com/jinzhu/copier"
)

type IService interface {
	Create(body *usersDTO.Create) (int, interface{})
	Find() (int, interface{})
	DeleteById(id string) (int, interface{})
	FindById(id string) (int, interface{})
	UpdateById(id string, body *usersDTO.Update) (int, interface{})
}

type Service struct {
	repo *usersRepo.Users
}

func New(repo *usersRepo.Users) IService {
	return &Service{repo: repo}
}

func (s *Service) Find() (int, interface{}) {
	query := map[string]interface{}{}
	users, err := s.repo.Find(query)

	if err != nil {
		return resService.InternalServerError(msgDomain.Msg.ERR_SAVING_IN_DATABASE)
	}

	return resService.Ok(users)
}

func (s *Service) Create(body *usersDTO.Create) (int, interface{}) {
	userIdFound, _ := s.repo.FindById(body.Id)

	if userIdFound.Id != "" {
		return resService.BadRequest(msgDomain.Msg.ERR_ID_ALREADY_EXISTS)
	}

	querySearchUser := map[string]interface{}{"email": body.Email}
	userEmailFound, err := s.repo.FindOne(querySearchUser)

	if userEmailFound.Id != "" {
		return resService.BadRequest(msgDomain.Msg.ERR_EMAIL_ALREADY_EXISTS)
	}

	userDomain := usersDomain.User{}

	copier.Copy(&userDomain, &body)

	user, err := s.repo.Create(userDomain)

	if err != nil {
		logService.Error(err.Error())
		return resService.InternalServerError(msgDomain.Msg.ERR_SAVING_IN_DATABASE)
	}

	return resService.Created(user)
}

func (s *Service) UpdateById(id string, body *usersDTO.Update) (int, interface{}) {
	userIdFound, _ := s.repo.FindById(id)

	if userIdFound.Id == "" {
		return resService.BadRequest(msgDomain.Msg.ERR_NOT_FOUND)
	}

	queryNotId := map[string]interface{}{"id": id}
	queryEmail := map[string]interface{}{"email": body.Email}

	userEmailFound, _ := s.repo.FindWithNot(queryNotId, queryEmail)

	if userEmailFound.Email != "" {
		return resService.BadRequest(msgDomain.Msg.ERR_EMAIL_ALREADY_EXISTS)
	}

	userDomain := usersDomain.User{}

	copier.Copy(&userDomain, &body)

	user, err := s.repo.UpdateById(id, userDomain)

	if err != nil {
		logService.Error(err.Error())
		return resService.InternalServerError(msgDomain.Msg.ERR_SAVING_IN_DATABASE)
	}

	return resService.Ok(user)
}

func (s *Service) DeleteById(id string) (int, interface{}) {
	user, err := s.repo.FindById(id)

	if err != nil {
		return resService.InternalServerError(msgDomain.Msg.ERR_SAVING_IN_DATABASE)
	}

	if user.Id == "" {
		return resService.NotFound(msgDomain.Msg.ERR_NOT_FOUND)
	}

	s.repo.DeleteById(id)

	userRes := usersDTO.UserRes{}
	copier.Copy(&userRes, &user)

	return resService.Ok(userRes)
}

func (s *Service) FindById(id string) (int, interface{}) {
	user, err := s.repo.FindById(id)

	if err != nil {
		return resService.InternalServerError(msgDomain.Msg.ERR_SAVING_IN_DATABASE)
	}

	if user.Id == "" {
		return resService.NotFound(msgDomain.Msg.ERR_NOT_FOUND)
	}

	userRes := usersDTO.UserRes{}
	copier.Copy(&userRes, &user)

	return resService.Ok(userRes)
}
