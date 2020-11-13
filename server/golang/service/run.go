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

var userCol *mongo.Collection

func init() {
	db.Session = database.ConnectToMongoDB()
	userCol = db.Session.Database("sakuraChat").Collection("users")
}

func RunServer() {
	listen, err := net.Listen("tcp", ":8806")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	talkService := TalkHandler{}
	authService := AuthHandler{}
	TalkRPC.RegisterTalkServiceServer(server, talkService)
	TalkRPC.RegisterAuthServiceServer(server, authService)
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
