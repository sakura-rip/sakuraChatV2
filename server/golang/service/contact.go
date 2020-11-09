package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl *TalkHandler) GetContact(ctx context.Context, in *TalkRPC.GetContactRequest, opts ...grpc.CallOption) (*TalkRPC.GetContactResponse, error) {

}
func (cl *TalkHandler) GetContacts(ctx context.Context, in *TalkRPC.GetContactsRequest, opts ...grpc.CallOption) (*TalkRPC.GetContactsResponse, error) {

}
func (cl *TalkHandler) UpdateContact(ctx context.Context, in *TalkRPC.UpdateContactRequest, opts ...grpc.CallOption) (*TalkRPC.UpdateContactResponse, error) {

}
func (cl *TalkHandler) BlockContact(ctx context.Context, in *TalkRPC.BlockContactRequest, opts ...grpc.CallOption) (*TalkRPC.BlockContactResponse, error) {

}
func (cl *TalkHandler) ReportContact(ctx context.Context, in *TalkRPC.ReportContactRequest, opts ...grpc.CallOption) (*TalkRPC.ReportContactResponse, error) {

}
