package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc"
)

func (cl *ChatHandler) SendFriendRequest(ctx context.Context, in *TalkRPC.SendFriendRequestRequest, opts ...grpc.CallOption) (*TalkRPC.SendFriendRequestResponse, error) {

}
func (cl *ChatHandler) AcceptFriendRequest(ctx context.Context, in *TalkRPC.AcceptFriendRequestRequest, opts ...grpc.CallOption) (*TalkRPC.AcceptFriendRequestResponse, error) {

}
func (cl *ChatHandler) RejectFriendRequest(ctx context.Context, in *TalkRPC.RejectFriendRequestRequest, opts ...grpc.CallOption) (*TalkRPC.RejectFriendRequestResponse, error) {

}
func (cl *ChatHandler) GetAllFriendIds(ctx context.Context, in *TalkRPC.GetAllFriendIdsRequest, opts ...grpc.CallOption) (*TalkRPC.GetAllFriendIdsResponse, error) {

}
func (cl *ChatHandler) GetAllFriendRequestIds(ctx context.Context, in *TalkRPC.GetAllFriendRequestIdsRequest, opts ...grpc.CallOption) (*TalkRPC.GetAllFriendRequestIdsResponse, error) {

}
func (cl *ChatHandler) GetALlFriendRequestedIds(ctx context.Context, in *TalkRPC.GetALlFriendRequestedIdsRequest, opts ...grpc.CallOption) (*TalkRPC.GetALlFriendRequestedIdsResponse, error) {

}
func (cl *ChatHandler) GetFriendRequestStatus(ctx context.Context, in *TalkRPC.GetFriendRequestStatusRequest, opts ...grpc.CallOption) (*TalkRPC.GetFriendRequestStatusResponse, error) {

}
func (cl *ChatHandler) GetAllBlockedIds(ctx context.Context, in *TalkRPC.GetAllBlockedIdsRequest, opts ...grpc.CallOption) (*TalkRPC.GetAllBlockedIdsResponse, error) {

}
