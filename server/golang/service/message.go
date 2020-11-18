package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
)

func (cl TalkHandler) SendMessage(ctx context.Context, in *TalkRPC.SendMessageRequest) (*TalkRPC.SendMessageResponse, error) {

}
func (cl TalkHandler) EditMessage(ctx context.Context, in *TalkRPC.EditMessageRequest) (*TalkRPC.EditMessageResponse, error) {

}
func (cl TalkHandler) UnsendMessage(ctx context.Context, in *TalkRPC.UnsendMessageRequest) (*TalkRPC.UnsendMessageResponse, error) {

}
func (cl TalkHandler) HideMessage(ctx context.Context, in *TalkRPC.HideMessageRequest) (*TalkRPC.HideMessageResponse, error) {

}
func (cl TalkHandler) ReportMessage(ctx context.Context, in *TalkRPC.ReportMessageRequest) (*TalkRPC.ReportMessageResponse, error) {
	return &TalkRPC.ReportMessageResponse{}, nil
}
