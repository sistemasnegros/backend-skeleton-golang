package usersModelMongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	CreatedAt time.Time          `bson:"createAt"`
	UpdatedAt time.Time          `bson:"updateAt"`
}
