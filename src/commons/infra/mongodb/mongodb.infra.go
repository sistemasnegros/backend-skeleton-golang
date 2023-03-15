package mongodbInfra

import (
	configService "backend-skeleton-golang/commons/app/services/config-service"
	logService "backend-skeleton-golang/commons/app/services/log-service"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New() *mongo.Database {
	ServerAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(configService.GetMongoDb()).SetServerAPIOptions(ServerAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic("Err connecting mongodb")
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	logService.Info("successfully connection mongodb!")

	return client.Database("skeleton")
}
