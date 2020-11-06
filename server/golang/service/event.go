package service

import (
	"context"

	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl *ChatHandler) FetchEvents(ctx context.Context, in *TalkRPC.FetchEventsRequest, opts ...grpc.CallOption) (TalkRPC.TalkService_FetchEventsClient, error) {

}
