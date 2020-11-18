package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl TalkHandler) GetAllTags(ctx context.Context, in *TalkRPC.GetAllTagsRequest) (*TalkRPC.GetAllTagsResponse, error) {

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

func (cl TalkHandler) AddTagToContact(context.Context, *TalkRPC.AddTagToContactRequest) (*TalkRPC.AddTagToContactResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
}
