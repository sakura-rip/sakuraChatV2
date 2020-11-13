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
}
func (cl TalkHandler) ReportContact(ctx context.Context, in *TalkRPC.ReportContactRequest) (*TalkRPC.ReportContactResponse, error) {
}
