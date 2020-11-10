package service

import (
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"net"
)

type TalkHandler struct{}

type AuthHandler struct{}

var db *database.DBClient

var userDB *mongo.Collection

func init() {
	db.Session = database.ConnectToMongoDB()
	userDB = db.Session.Database("sakuraChat").Collection("users")
}

func runServer() {
	listen, err := net.Listen("tcp", ":8806")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	talkService := TalkHandler{}
	TalkRPC.RegisterTalkServiceServer(server, talkService)
	server.Serve(listen)
}
