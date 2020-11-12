package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl TalkHandler) UpdateProfile(ctx context.Context, in *TalkRPC.UpdateProfileRequest) (*TalkRPC.UpdateProfileResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	var attToUpdate []bson.E
	for _, key := range in.Keys {
		switch key {
		case TalkRPC.ProfileKey_NAME:
			attToUpdate = append(attToUpdate, bson.E{Key: "name", Value: in.Profile.Name})
			fallthrough
		case TalkRPC.ProfileKey_BIO:
			attToUpdate = append(attToUpdate, bson.E{Key: "bio", Value: in.Profile.Bio})
		}
	}
	_, dberr := userCol.UpdateOne(
		ctx,
		bson.M{"_id": uuid},
		bson.D{
			{"$set", attToUpdate},
		},
	)
	if dberr != nil {
		return nil, status.New(codes.Internal, "db error").Err()
	}
	return &TalkRPC.UpdateProfileResponse{}, nil
}

func (cl TalkHandler) GetProfile(ctx context.Context, _ *TalkRPC.GetProfileRequest) (*TalkRPC.GetProfileResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	//FIXME: ref: https://docs.mongodb.com/manual/tutorial/project-fields-from-query-results/#return-the-specified-fields-and-the-id-field-only

	user, err := findUserFromDB(uuid, bson.D{{"profile", 0}})
	if err != nil {
		return nil, err
	}
	profile := &TalkRPC.GetProfileResponse{
		Profile: &TalkRPC.Profile{
			UUID:        uuid,
			Name:        user.Profile.Name,
			Bio:         user.Profile.Bio,
			IconPath:    user.Profile.IconPath,
			CoverPath:   user.Profile.CoverPath,
			TwitterID:   user.Profile.TwitterID,
			InstagramID: user.Profile.InstagramID,
			GithubID:    user.Profile.GithubID,
		},
	}
	return profile, nil
}
