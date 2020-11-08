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

var ctx context.Context
var DB *DBClient

func init() {
	ctx = context.Background()
	mctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	db, err := mongo.Connect(mctx, options.Client().ApplyURI("mangodb://localhost:27017"))
	if err != nil {
		panic(fmt.Errorf("error on init mongo db: %v", err))
	}
	DB = &DBClient{Session: db}
}
