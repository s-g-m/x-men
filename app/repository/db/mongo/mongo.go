package mongo

import (
	"context"
	"time"
	"x-men/util/constant"
	"x-men/util/logs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const timeOut = 30 * time.Second

type Mongo struct {
	db *mongo.Database
}

func NewMongo(uri string, dbName string) Mongo {
	client, err := connect(uri, dbName)

	if err != nil {
		logs.Fatal(constant.ErrorConectDB, err.Error())
	}

	logs.Info("connected mongo database")
	return Mongo{db: client}
}

func connect(uri string, dbName string) (db *mongo.Database, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetMaxPoolSize(10))

	if err != nil {
		return
	}

	if err = client.Connect(ctx); err != nil {
		return
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return
	}

	db = client.Database(dbName)

	return
}
