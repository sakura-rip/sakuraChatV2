package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl TalkHandler) SendFriendRequest(ctx context.Context, in *TalkRPC.SendFriendRequestRequest) (*TalkRPC.SendFriendRequestResponse, error) {
	uuid, ok, _ := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return nil, status.New(codes.Unauthenticated, "Invalid Token").Err()
	}

}
func (cl TalkHandler) AcceptFriendRequest(ctx context.Context, in *TalkRPC.AcceptFriendRequestRequest) (*TalkRPC.AcceptFriendRequestResponse, error) {

}
func (cl TalkHandler) RejectFriendRequest(ctx context.Context, in *TalkRPC.RejectFriendRequestRequest) (*TalkRPC.RejectFriendRequestResponse, error) {

}
func (cl TalkHandler) GetAllFriendIds(ctx context.Context, in *TalkRPC.GetAllFriendIdsRequest) (*TalkRPC.GetAllFriendIdsResponse, error) {

}
func (cl TalkHandler) GetAllFriendRequestIds(ctx context.Context, in *TalkRPC.GetAllFriendRequestIdsRequest) (*TalkRPC.GetAllFriendRequestIdsResponse, error) {

}

func (cl TalkHandler) GetFriendRequestStatus(ctx context.Context, in *TalkRPC.GetFriendRequestStatusRequest) (*TalkRPC.GetFriendRequestStatusResponse, error) {

}
func (cl TalkHandler) GetAllBlockedIds(ctx context.Context, in *TalkRPC.GetAllBlockedIdsRequest) (*TalkRPC.GetAllBlockedIdsResponse, error) {

}
