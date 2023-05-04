package persistence

import (
	"GoPetClinic/src/config"
	"GoPetClinic/src/system"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const DatabaseName = "petclinic"

var client *mongo.Client

func BootstrapMongoDB(ctx context.Context, appConfig *config.AppConfig) (system.ShutdownChannel, error) {
	logger := system.GetLogger("mongo")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?maxPoolSize=20&w=majority", appConfig.MongoUsername, appConfig.MongoPassword, appConfig.MongoHost, appConfig.MongoPort)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	var result bson.M
	if err = client.Database(DatabaseName).RunCommand(context.Background(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, fmt.Errorf("could not connect to mongodb: %w", err)
	}

	logger.Info("MongoDB connected")

	doneChan := make(chan error)

	go func() {
		<-ctx.Done()

		disconnectCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		err := client.Disconnect(disconnectCtx)
		doneChan <- err
	}()

	return doneChan, nil
}

type GetDB = func() *mongo.Database

func GetDatabase() *mongo.Database {
	if client == nil {
		panic("MongoDB is not bootstrapped")
	}

	return client.Database(DatabaseName)
}
