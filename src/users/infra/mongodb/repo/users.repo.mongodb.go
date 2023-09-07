package usersRepoMongo

import (
	genericRepoMongodb "backend-skeleton-golang/commons/infra/mongodb/repo"

	usersDomain "backend-skeleton-golang/users/domain"
	usersModelMongodb "backend-skeleton-golang/users/infra/mongodb/models"

	"go.mongodb.org/mongo-driver/mongo"
)


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
