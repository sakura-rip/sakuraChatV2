package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl *TalkHandler) SendMessage(ctx context.Context, in *TalkRPC.SendMessageRequest, opts ...grpc.CallOption) (*TalkRPC.SendMessageResponse, error) {

}
func (cl *TalkHandler) EditMessage(ctx context.Context, in *TalkRPC.EditMessageRequest, opts ...grpc.CallOption) (*TalkRPC.EditMessageResponse, error) {

}
func (cl *TalkHandler) UnsendMessage(ctx context.Context, in *TalkRPC.UnsendMessageRequest, opts ...grpc.CallOption) (*TalkRPC.UnsendMessageResponse, error) {

}
func (cl *TalkHandler) HideMessage(ctx context.Context, in *TalkRPC.HideMessageRequest, opts ...grpc.CallOption) (*TalkRPC.HideMessageResponse, error) {

}
func (cl *TalkHandler) ReportMessage(ctx context.Context, in *TalkRPC.ReportMessageRequest, opts ...grpc.CallOption) (*TalkRPC.ReportMessageResponse, error) {

}
