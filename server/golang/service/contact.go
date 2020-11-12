package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl TalkHandler) GetContact(ctx context.Context, in *TalkRPC.GetContactRequest) (*TalkRPC.GetContactResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}
	dbContact, err := findContactFromDB(uuid, in.UUID)
	if err != nil {
		return nil, status.New(codes.NotFound, "target not found").Err()
	}
	contact := TalkRPC.GetContactResponse{
		Contact: &TalkRPC.Contact{},
	}
	return &contact, nil

}
func (cl TalkHandler) GetContacts(ctx context.Context, in *TalkRPC.GetContactsRequest) (*TalkRPC.GetContactsResponse, error) {

}
func (cl TalkHandler) UpdateContact(ctx context.Context, in *TalkRPC.UpdateContactRequest) (*TalkRPC.UpdateContactResponse, error) {

}
func (cl TalkHandler) BlockContact(ctx context.Context, in *TalkRPC.BlockContactRequest) (*TalkRPC.BlockContactResponse, error) {

}
func (cl TalkHandler) ReportContact(ctx context.Context, in *TalkRPC.ReportContactRequest) (*TalkRPC.ReportContactResponse, error) {

}
