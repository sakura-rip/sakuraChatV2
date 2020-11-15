package service

import (
	"context"
	"fmt"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"github.com/ch31212y/sakuraChatV2/database"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl TalkHandler) GetContact(ctx context.Context, in *TalkRPC.GetContactRequest) (*TalkRPC.GetContactResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	contact, err := findRPCContactFromDB(uuid, in.UUID)
	return &TalkRPC.GetContactResponse{
		Contact: contact,
	}, err
}

func (cl TalkHandler) GetContacts(ctx context.Context, in *TalkRPC.GetContactsRequest) (*TalkRPC.GetContactsResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	response := &TalkRPC.GetContactsResponse{
		Contacts: []*TalkRPC.Contact{},
	}
	for _, cuuid := range in.UUIDs {
		con, err := findRPCContactFromDB(uuid, cuuid)
		if err != nil {
			response.Contacts = append(response.Contacts, con)
		}
	}
	return response, nil
}
func (cl TalkHandler) UpdateContact(ctx context.Context, in *TalkRPC.UpdateContactRequest) (*TalkRPC.UpdateContactResponse, error) {

}

func (cl TalkHandler) BlockContact(ctx context.Context, in *TalkRPC.BlockContactRequest) (*TalkRPC.BlockContactResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	contact, err := findRPCContactFromDB(uuid, in.UUID)
	if err != nil {
		return nil, err
	}
	switch contact.Status {
	case TalkRPC.FriendStatus_friend:
		_, _ := userCol.UpdateOne(
			ctx,
			bson.M{"_id": uuid},
			"", //https://docs.mongodb.com/manual/reference/operator/update/pull/
		)
	case TalkRPC.FriendStatus_delete:
	case TalkRPC.FriendStatus_block:
	case TalkRPC.FriendStatus_nothing:
	}
	updateAttr := bson.M{"$set": bson.M{
		fmt.Sprintf("contacts.%d", in.UUID): database.Contact{
			UUID:            in.UUID,
			OverWrittenName: contact.OverWrittenName,
			Status:          2,
			TagIds:          contact.TagIds,
		}},
	}
	if _, err := userCol.UpdateOne(
		ctx,
		bson.M{"_id": uuid},
		updateAttr,
	); err != nil {
		return nil, status.New(codes.NotFound, "no contact found").Err()
	}
}

func (cl TalkHandler) ReportContact(ctx context.Context, in *TalkRPC.ReportContactRequest) (*TalkRPC.ReportContactResponse, error) {
}
