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
			attToUpdate = append(attToUpdate, bson.E{Key: "profile.name", Value: in.Profile.Name})
			fallthrough
		case TalkRPC.ProfileKey_BIO:
			attToUpdate = append(attToUpdate, bson.E{Key: "profile.bio", Value: in.Profile.Bio})
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

	profile, err := findRPCProfileFromDB(uuid)
	return &TalkRPC.GetProfileResponse{
		Profile: profile,
	}, err
}
