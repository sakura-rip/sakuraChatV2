package service

import (
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
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
		TagIds:          map[string]int64{},
	}
	rs := userCol.FindOne(
		ctx,
		bson.D{{"_id", uuid}},
		options.FindOne().SetProjection(bson.D{{"contacts", 1}}),
	)
	var user *database.User
	if rs.Decode(&user) != nil {
		return nil, status.New(codes.NotFound, "user not found").Err()
	}
	if len(user.Contacts) > 0 {
		value, ok := user.Contacts[targetUUID]
		if ok == true {
			contact.OverWrittenName = value.OverWrittenName
			contact.Status = value.Status
			contact.TagIds = value.TagIds
		}
	}
	return &contact, nil
}

func findRPCProfileFromDB(uuid string) (*TalkRPC.Profile, error) {
	dbProf, err := findProfileFromDB(uuid)
	if err != nil {
		return nil, err
	}
	profile := TalkRPC.Profile{
		UUID:        uuid,
		Name:        dbProf.Name,
		Bio:         dbProf.Bio,
		IconPath:    dbProf.IconPath,
		CoverPath:   dbProf.CoverPath,
		TwitterID:   dbProf.TwitterID,
		InstagramID: dbProf.InstagramID,
		GithubID:    dbProf.GithubID,
	}
	return &profile, nil
}

func findRPCSettingFromDB(uuid string) (*TalkRPC.Setting, error) {
	dbSet, err := findSettingFromDB(uuid)
	if err != nil {
		return nil, err
	}
	setting := TalkRPC.Setting{
		PrivateUserID:              dbSet.PrivateUserID,
		AllowSearchByPrivateUserID: dbSet.AllowSearchByPrivateUserID,
		Email:                      dbSet.Email,
		AllowSearchByEmail:         dbSet.AllowSearchByEmail,
		UserTicket:                 dbSet.UserTicket,
		AllowSearchByUserTicket:    dbSet.AllowSearchByUserTicket,
	}
	return &setting, nil
}

func findRPCContactFromDB(baseUUID, targetUUID string) (*TalkRPC.Contact, error) {
	profile, err := findRPCProfileFromDB(targetUUID)
	if err != nil {
		return nil, err
	}
	contact := TalkRPC.Contact{
		UUID:            targetUUID,
		Name:            profile.Name,
		OverWrittenName: profile.Name,
		Status:          3,
		TagIds:          []string{},
	}
	rs := userCol.FindOne(
		ctx,
		bson.D{{"_id", baseUUID}, {"contacts.", targetUUID}},
		options.FindOne().SetProjection(bson.D{{"contacts", 1}}),
	)
	var user *database.User
	if rs.Decode(&user) != nil {
		return nil, status.New(codes.NotFound, "user not found").Err()
	}
	if len(user.Contacts) == 0 {
		return nil, err
	}
	if _, ok := user.Contacts[targetUUID]; ok == true {
		contact.OverWrittenName = user.Contacts[targetUUID].OverWrittenName
		contact.TagIds = mapToSlice(user.Contacts[targetUUID].TagIds)
		switch user.Contacts[targetUUID].Status {
		case 0:
			contact.Status = TalkRPC.FriendStatus_friend
		case 1:
			contact.Status = TalkRPC.FriendStatus_block
		case 2:
			contact.Status = TalkRPC.FriendStatus_delete
		case 3:
			contact.Status = TalkRPC.FriendStatus_nothing
		}
	}
	return &contact, nil
}
