package service

import (
	"context"

	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl *TalkHandler) FindUserByUserID(ctx context.Context, in *TalkRPC.FindUserByUserIDRequest, opts ...grpc.CallOption) (*TalkRPC.FindUserByUserIDResponse, error) {

}
func (cl *TalkHandler) FindUserByTicket(ctx context.Context, in *TalkRPC.FinduUserByTicketRequest, opts ...grpc.CallOption) (*TalkRPC.FindUserByTicketResponse, error) {

}
func (cl *TalkHandler) FindUserByEmail(ctx context.Context, in *TalkRPC.FindUserByEmailRequest, opts ...grpc.CallOption) (*TalkRPC.FindUserByEmailResponse, error) {

}
