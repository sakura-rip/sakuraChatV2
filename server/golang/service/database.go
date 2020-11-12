package service

import (
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func findUserFromDB(id string, projection bson.D) (*database.User, error) {
	rs := userCol.FindOne(
		ctx,
		bson.D{{"_id", id}},
		options.FindOne().SetProjection(projection),
	)
	var user *database.User
	if rs.Decode(&user) != nil {
		return user, status.New(codes.NotFound, "user not found").Err()
	}
	return user, nil
}

func findProfileFromDB(uuid string) (*database.Profile, error) {
	rs := userCol.FindOne(
		ctx,
		bson.D{{"_id", uuid}},
		options.FindOne().SetProjection(bson.D{{"profile", 1}}),
	)
	var user *database.User
	if rs.Decode(&user) != nil {
		return nil, status.New(codes.NotFound, "user not found").Err()
	}
	return &user.Profile, nil
}

func findSettingFromDB(uuid string) (*database.Setting, error) {
	rs := userCol.FindOne(
		ctx,
		bson.D{{"_id", uuid}},
		options.FindOne().SetProjection(bson.D{{"setting", 1}}),
	)
	var user *database.User
	if rs.Decode(&user) != nil {
		return nil, status.New(codes.NotFound, "user not found").Err()
	}
	return &user.Setting, nil
}

func findContactFromDB(uuid, targetUUID string) (*database.Contact, error) {
	profile, err := findProfileFromDB(targetUUID)
	if err != nil {
		return nil, err
	}
	contact := database.Contact{
		UUID:            targetUUID,
		OverWrittenName: profile.Name,
		Status:          3,
		TagIds:          []string{},
	}
	rs := userCol.FindOne(
		ctx,
		bson.D{{"_id", uuid}},
		options.FindOne().SetProjection(bson.M{"contacts": bson.M{"$elemMatch": bson.M{"uuid": targetUUID}}}),
	)
	var user *database.User
	if rs.Decode(&user) != nil {
		return nil, status.New(codes.NotFound, "user not found").Err()
	}
	if len(user.Contacts) != 0 {
		contact.OverWrittenName = user.Contacts[0].OverWrittenName
		contact.Status = user.Contacts[0].Status
		contact.TagIds = user.Contacts[0].TagIds
	}
	return &contact, nil
}
