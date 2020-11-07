package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl *AuthHandler) VerifyIDToken(ctx context.Context, in *TalkRPC.VerifyIDTokenRequest, opts ...grpc.CallOption) (*TalkRPC.VerifyIDTokenResponse, error) {

}
