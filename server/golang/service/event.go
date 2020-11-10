package service

import (
	"context"

	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl TalkHandler) FetchEvents(*TalkRPC.FetchEventsRequest, TalkRPC.TalkService_FetchEventsServer) error {

}
