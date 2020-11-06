package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl *ChatHandler) FindUserByUserID(ctx context.Context, in *TalkRPC.FindUserByUserIDRequest, opts ...grpc.CallOption) (*TalkRPC.FindUserByUserIDResponse, error) {

}
func (cl *ChatHandler) FindUserByTicket(ctx context.Context, in *TalkRPC.FinduserByTicketRequest, opts ...grpc.CallOption) (*TalkRPC.FinduserByTicketResponse, error) {

}
func (cl *ChatHandler) FindUserByEmail(ctx context.Context, in *TalkRPC.FindUserByEmailRequest, opts ...grpc.CallOption) (*TalkRPC.FindUserByEmailResponse, error) {

}
