package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DBClient struct {
	Session *mongo.Client
}

var database *DBClient

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI("mangodb://localhost:27017"))
	if err != nil {
		panic(fmt.Errorf("error on init mongo db: %v", err))
	}
	database = &DBClient{Session: db}
}
