package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl *TalkHandler) UpdateProfile(ctx context.Context, in *TalkRPC.UpdateProfileRequest, opts ...grpc.CallOption) (*TalkRPC.UpdateProfileResponse, error) {

}
func (cl *TalkHandler) GetProfile(ctx context.Context, _ *TalkRPC.GetProfileRequest, _ ...grpc.CallOption) (*TalkRPC.GetProfileResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	//FIXME: ref: https://docs.mongodb.com/manual/tutorial/project-fields-from-query-results/#return-the-specified-fields-and-the-id-field-only
	var DBProfile database.Profile
	filter := bson.M{"_id": uuid}
	if err := userDB.FindOne(ctx, filter).Decode(&DBProfile); err != nil {
		return nil, err
	}
	profile := &TalkRPC.GetProfileResponse{
		Profile: &TalkRPC.Profile{
			Uuid:        uuid,
			Name:        DBProfile.Name,
			Bio:         DBProfile.Bio,
			IconPath:    DBProfile.IconPath,
			CoverPath:   DBProfile.CoverPath,
			TwitterID:   DBProfile.TwitterID,
			InstagramID: DBProfile.InstagramID,
			GithubID:    DBProfile.GithubID,
		},
	}
	return profile, nil
}
