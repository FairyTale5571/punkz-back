package db

import (
	"context"
	"time"

	"github.com/fairytale5571/punkz/internal/site"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Provider interface {
	site.DBProvider
}

type database struct {
	client *mongo.Client
	mainDB *mongo.Database
	ctx    context.Context
}

const (
	connectionTimeoutInSecond = 200
	databaseName              = "punkz"
)

func NewProvider(mongoUri string) (Provider, error) {
	var err error
	dbConnection := database{
		ctx: context.Background(),
	}
	if dbConnection.client, err = connectMongoDB(mongoUri); err != nil {
		return nil, err
	}
	dbConnection.mainDB = dbConnection.client.Database(databaseName)
	if err := dbConnection.client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}

	return &dbConnection, nil
}

func connectMongoDB(connectionString string) (*mongo.Client, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), connectionTimeoutInSecond*time.Second)
	defer cancelFunc()
	return mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
}
