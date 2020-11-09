package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
	"github.com/ch31212y/sakuraChatV2/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (cl *AuthHandler) VerifyIDToken(ctx context.Context, in *TalkRPC.VerifyIDTokenRequest, opts ...grpc.CallOption) (*TalkRPC.VerifyIDTokenResponse, error) {
	jwt, err := auth.VerifyIDToken(ctx, in.JwtToken)
	if err != nil {
		return &TalkRPC.VerifyIDTokenResponse{}, status.New(codes.InvalidArgument, "bad token").Err()
	}
	claims := map[string]interface{}{
		"registered": false,
	}
	//premiumカラムをつけてToken 再発行
	token, err := auth.CustomTokenWithClaims(ctx, jwt.UID, claims)
	if err != nil {
		return &TalkRPC.VerifyIDTokenResponse{}, status.New(codes.Internal, "internal error").Err()
	}
	//クライアントのヘッダーを追加
	_ = grpc.SetHeader(ctx, metadata.Pairs("x-chat-token", token))
	return &TalkRPC.VerifyIDTokenResponse{}, nil
}

func (cl *AuthHandler) InitPrimaryAccount(ctx context.Context, in *TalkRPC.InitPrimaryAccountRequest, opts ...grpc.CallOption) (*TalkRPC.InitPrimaryAccountResponse, error) {
	token, ok := getHeader(ctx, "X-Chat-Access")
	if ok == false {
		return &TalkRPC.InitPrimaryAccountResponse{}, status.New(codes.Unauthenticated, "").Err()
	}
	jwt, err := auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		return &TalkRPC.InitPrimaryAccountResponse{}, status.New(codes.Unauthenticated, "").Err()
	}
	isInit, ok := jwt.Claims["registered"]
	//認証されていないかつカラムが存在する
	if isInit == false && ok == true {
		//User情報をDBに挿入
		docu := database.User{
			ID:              jwt.UID,
			Profile:         database.Profile{},
			Setting:         database.Setting{},
			JoinedGroupIds:  []string{},
			InvitedGroupIds: []string{},
			FriendRequests:  []database.FriendRequest{},
			FriendIds:       []string{},
			BlockedIds:      []string{},
			DeletedIds:      []string{},
		}
		_, err := userDB.InsertOne(ctx, docu)
		if err != nil {
			return &TalkRPC.InitPrimaryAccountResponse{}, status.New(codes.Internal, "").Err()
		}
		return &TalkRPC.InitPrimaryAccountResponse{}, nil
	}
	return &TalkRPC.InitPrimaryAccountResponse{}, status.New(codes.PermissionDenied, "").Err()
}
