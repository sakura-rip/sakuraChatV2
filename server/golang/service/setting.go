package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl *ChatHandler) UpdateSetting(ctx context.Context, in *TalkRPC.UpdateSettingRequest, opts ...grpc.CallOption) (*TalkRPC.UpdateSettingResponse, error) {

}
func (cl *ChatHandler) GetSetting(ctx context.Context, in *TalkRPC.GetSettingRequest, opts ...grpc.CallOption) (*TalkRPC.GetSettingResponse, error) {

}
