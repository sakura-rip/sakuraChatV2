package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ch31212y/sakuraChatV2/TalkRPC"
)

func (cl TalkHandler) FindUserByUserID(ctx context.Context, in *TalkRPC.FindUserByUserIDRequest) (*TalkRPC.FindUserByUserIDResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	rs := userCol.FindOne(
		ctx,
		bson.D{{"setting.PUserID", in.PrivateUserID}, {"setting.asbPUserID", true}},
		options.FindOne().SetProjection(bson.M{"_id": 1}),
	)
	var user *database.User
	if rs.Decode(&user) != nil {
		return nil, status.New(codes.NotFound, "No user found").Err()
	}
	contact, err := findRPCContactFromDB(uuid, user.ID)
	if err != nil {
		return nil, status.New(codes.NotFound, "No user found").Err()
	}
	return &TalkRPC.FindUserByUserIDResponse{
		Contact: contact,
	}, nil
}

func (cl TalkHandler) FindUserByTicket(ctx context.Context, in *TalkRPC.FindUserByTicketRequest) (*TalkRPC.FindUserByTicketResponse, error) {

}
func (cl TalkHandler) FindUserByEmail(ctx context.Context, in *TalkRPC.FindUserByEmailRequest) (*TalkRPC.FindUserByEmailResponse, error) {

}
