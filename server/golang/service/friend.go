package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (cl TalkHandler) SendFriendRequest(ctx context.Context, in *TalkRPC.SendFriendRequestRequest) (*TalkRPC.SendFriendRequestResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	field := bson.M{"FriendRequests." + in.UUID: database.FriendRequest{
		FromID:      uuid,
		ToID:        in.UUID,
		CreatedTime: time.Now().Unix(),
		Message:     in.Message,
	}}
	_, dbErr := userCol.UpdateOne(
		ctx,
		bson.M{"_id": uuid},
		bson.M{"$set": field},
	)
	if dbErr != nil {
		return nil, status.New(codes.Internal, "DB error").Err()
	}
	//受信したもの
	_, dbErr2 := userCol.UpdateOne(
		ctx,
		bson.M{"_id": in.UUID},
		bson.M{"$set": field},
	)
	if dbErr2 != nil {
		return nil, status.New(codes.Internal, "db error").Err()
	}
	return &TalkRPC.SendFriendRequestResponse{}, nil
}
func (cl TalkHandler) AcceptFriendRequest(ctx context.Context, in *TalkRPC.AcceptFriendRequestRequest) (*TalkRPC.AcceptFriendRequestResponse, error) {

}
func (cl TalkHandler) RejectFriendRequest(ctx context.Context, in *TalkRPC.RejectFriendRequestRequest) (*TalkRPC.RejectFriendRequestResponse, error) {

}
func (cl TalkHandler) GetAllFriendIds(ctx context.Context, in *TalkRPC.GetAllFriendIdsRequest) (*TalkRPC.GetAllFriendIdsResponse, error) {

}
func (cl TalkHandler) GetAllFriendRequestIds(ctx context.Context, in *TalkRPC.GetAllFriendRequestIdsRequest) (*TalkRPC.GetAllFriendRequestIdsResponse, error) {

}

func (cl TalkHandler) GetFriendRequestStatus(ctx context.Context, in *TalkRPC.GetFriendRequestStatusRequest) (*TalkRPC.GetFriendRequestStatusResponse, error) {

}
func (cl TalkHandler) GetAllBlockedIds(ctx context.Context, in *TalkRPC.GetAllBlockedIdsRequest) (*TalkRPC.GetAllBlockedIdsResponse, error) {

}
