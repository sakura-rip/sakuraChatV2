package service

import (
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type TalkHandler struct{}

type AuthHandler struct{}

var db *database.DBClient

var userDB *mongo.Collection

func init() {
	db.Session = database.ConnectToMongoDB()
	userDB = db.Session.Database("sakuraChat").Collection("users")
}
