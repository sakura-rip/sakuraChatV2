package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl TalkHandler) UpdateSetting(ctx context.Context, in *TalkRPC.UpdateSettingRequest) (*TalkRPC.UpdateSettingResponse, error) {

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
