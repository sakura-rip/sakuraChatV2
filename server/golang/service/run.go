package service

import "github.com/ch31212y/sakuraChatV2/database"

type TalkHandler struct{}

type AuthHandler struct{}

var db *database.DBClient

func init() {
	db.Session = database.ConnectToMongoDB()
}
