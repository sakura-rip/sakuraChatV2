package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl TalkHandler) UpdateSetting(ctx context.Context, in *TalkRPC.UpdateSettingRequest) (*TalkRPC.UpdateSettingResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	var attToUpdate []bson.E

	for _, key := range in.Keys {
		switch key {
		case TalkRPC.SettingKey_PRIVATE_USER_ID:
			attToUpdate = append(attToUpdate, bson.E{Key: "PUserID", Value: in.Setting.PrivateUserID})
			fallthrough
		case TalkRPC.SettingKey_ALLOW_SEARCH_BY_PRIVATE_USER_ID:
			attToUpdate = append(attToUpdate, bson.E{Key: "asbPUserID", Value: in.Setting.AllowSearchByPrivateUserID})
			fallthrough
		case TalkRPC.SettingKey_EMAIL:
			attToUpdate = append(attToUpdate, bson.E{Key: "Email", Value: in.Setting.Email})
			fallthrough
		case TalkRPC.SettingKey_ALLOW_SEARCH_BY_EMAIL:
			attToUpdate = append(attToUpdate, bson.E{Key: "asbEmail", Value: in.Setting.AllowSearchByEmail})
			fallthrough
		case TalkRPC.SettingKey_USER_TICKET:
			attToUpdate = append(attToUpdate, bson.E{Key: "UTicket", Value: in.Setting.UserTicket})
			fallthrough
		case TalkRPC.SettingKey_ALLOW_SEARCH_BY_USER_TICKET:
			attToUpdate = append(attToUpdate, bson.E{Key: "asbUTicket", Value: in.Setting.AllowSearchByUserTicket})
		}
	}
	_, dberr := userDB.UpdateOne(
		ctx,
		bson.M{"_id": uuid},
		bson.D{
			{"$set", attToUpdate},
		},
	)
	if dberr != nil {
		return nil, status.New(codes.Internal, "db error").Err()
	}
	return &TalkRPC.UpdateSettingResponse{}, nil
}

func (cl TalkHandler) GetSetting(ctx context.Context, in *TalkRPC.GetSettingRequest) (*TalkRPC.GetSettingResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}

	user, err := findUserFromDB(uuid, bson.D{{"setting", 0}})
	if err != nil {
		return nil, err
	}
	setting := &TalkRPC.GetSettingResponse{
		Setting: &TalkRPC.Setting{
			PrivateUserID:              user.Setting.PrivateUserID,
			AllowSearchByPrivateUserID: user.Setting.AllowSearchByPrivateUserID,
			Email:                      user.Setting.Email,
			AllowSearchByEmail:         user.Setting.AllowSearchByEmail,
			UserTicket:                 user.Setting.UserTicket,
			AllowSearchByUserTicket:    user.Setting.AllowSearchByUserTicket,
		},
	}
	return setting, nil
}
