package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl *ChatHandler) UpdateProfile(ctx context.Context, in *TalkRPC.UpdateProfileRequest, opts ...grpc.CallOption) (*TalkRPC.UpdateProfileResponse, error) {

}
func (cl *ChatHandler) GetProfile(ctx context.Context, in *TalkRPC.GetProfileRequest, opts ...grpc.CallOption) (*TalkRPC.GetProfileResponse, error) {

}
