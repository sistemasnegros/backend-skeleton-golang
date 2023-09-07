package mongodbInfra

import (
	configService "backend-skeleton-golang/commons/app/services/config-service"
	logService "backend-skeleton-golang/commons/app/services/log-service"
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New() (*mongo.Database, *mongo.Client) {
	ServerAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(configService.GetMongoDb()).SetServerAPIOptions(ServerAPI)

	// opts.SetConnectTimeout(time.Second * 1)
  opts.SetTimeout(time.Second * 1)
  // opts.SetSocketTimeout(time.Second * 1)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic("err connecting mongodb")
	}

	logService.Info("successfully connection mongodb!")

	databaseCuted := strings.Split(configService.GetMongoDb(), "/")
	databaseCuted = strings.Split(databaseCuted[len(databaseCuted)-1], "?")
	database := databaseCuted[0]

	return client.Database(database), client
}
