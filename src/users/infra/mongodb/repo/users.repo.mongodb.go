package usersRepoMongo

import (
	genericRepoMongodb "backend-skeleton-golang/commons/infra/mongodb/repo"

	usersDomain "backend-skeleton-golang/users/domain"
	usersModelMongodb "backend-skeleton-golang/users/infra/mongodb/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// type IUsers interface {
// 	Find(query interface{}) ([]usersDomain.User, error)
// 	FindById(id string) (*usersDomain.User, error)
// 	FindOne(query interface{}) (*usersDomain.User, error)
// 	FindPagination(query interface{}, limit int64, page int) (*serviceDomain.PaginationData[usersDomain.User], error)
// 	UpdateById(id string, data interface{}) (*usersDomain.User, error)
// 	Create(usersDomain.User) (usersDomain.User, error)
// 	FindWithNot(queryNot map[string]interface{}, query map[string]interface{}) (*usersDomain.User, error)
// 	DeleteById(id string) error
// }

type IUsers interface {
	genericRepoMongodb.IGeneric[usersModelMongodb.User, usersDomain.User]
}

type Users struct {
	genericRepoMongodb.Generic[usersModelMongodb.User, usersDomain.User]
}

func New(DB *mongo.Database) IUsers {
	repo := &Users{genericRepoMongodb.Generic[usersModelMongodb.User, usersDomain.User]{DB: DB.Collection("users")}}
	return repo
}
