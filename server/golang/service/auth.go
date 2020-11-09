package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
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
		"premium": false,
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
	uuid, ok := VerifyTokenAndGetUUID(ctx)
	if ok == false {
		return &TalkRPC.InitPrimaryAccountResponse{}, status.New(codes.Unauthenticated, "").Err()
	}
	//	TODO: DB操作
}
