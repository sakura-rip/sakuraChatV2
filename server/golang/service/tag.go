package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl TalkHandler) GetAllTags(ctx context.Context, in *TalkRPC.GetAllTagsRequest) (*TalkRPC.GetAllTagsResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	var user *database.User
	rs := userCol.FindOne(
		ctx,
		bson.M{"_id": uuid},
		options.FindOne().SetProjection(bson.D{{"tags", 1}}),
	)
	if rs.Decode(&user) != nil {
		return nil, status.New(codes.Internal, "db error").Err()
	}
	var result []*TalkRPC.Tag
	for key, value := range user.Tags {
		result = append(result, &TalkRPC.Tag{
			TagID: key,
			Name: value.Name,
			Description: value.Description,
			Color: value.Color,
			CreatorUUID: value.CreatorUUID,
			CreatedTime: value.CreatedTime,
		})
	}
	return &TalkRPC.GetAllTagsResponse{Tags: result}, nil
}

func (cl TalkHandler) CreateTag(ctx context.Context, in *TalkRPC.CreateTagRequest) (*TalkRPC.CreateTagResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	//TODO: TagIDのGenするかしないかの設定
	_, dbError := userCol.UpdateOne(
		ctx,
		bson.M{"_id": uuid},
		bson.D{
			{"$set", bson.D{{"tags." + in.Tag.TagID, database.Tag{
				TagID:       in.Tag.TagID,
				Name:        in.Tag.Name,
				Description: in.Tag.Description,
				Color:       in.Tag.Color,
				CreatorUUID: in.Tag.CreatorUUID,
				CreatedTime: in.Tag.CreatedTime,
			}}}},
		})
	if dbError != nil {
		return nil, status.New(codes.Internal, "db error").Err()
	}
	return &TalkRPC.CreateTagResponse{}, nil
}

func (cl TalkHandler) DeleteTag(ctx context.Context, in *TalkRPC.DeleteTagRequest) (*TalkRPC.DeleteTagResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	_, dbError := userCol.UpdateOne(
		ctx,
		bson.M{"_id": uuid},
		bson.M{"$unset": "tags." + in.TagID},
	)
	if dbError != nil {
		return nil, status.New(codes.Internal, "db error").Err()
	}
	return &TalkRPC.DeleteTagResponse{}, nil
}

func (cl TalkHandler) AddTagToContact(ctx context.Context, in *TalkRPC.AddTagToContactRequest) (*TalkRPC.AddTagToContactResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	_, dbErr := userCol.UpdateOne(
		ctx,
		bson.M{"_id": uuid},
		bson.M{"$push": bson.M{"contacts." + in.TargetUUID + ".tags": in.TagID}}
	)
	if dbErr != nil {
		return nil, status.New(codes.Internal, "db error").Err()
	}
	return &TalkRPC.AddTagToContactResponse{}, nil
}
