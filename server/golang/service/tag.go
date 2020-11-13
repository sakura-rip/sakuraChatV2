package service

import (
	"context"
	"github.com/ch31212y/sakuraChatV2/TalkRPC"
)

func (cl TalkHandler) GetAllTags(ctx context.Context, in *TalkRPC.GetAllTagsRequest) (*TalkRPC.GetAllTagsResponse, error) {

}

func (cl TalkHandler) CreateTag(ctx context.Context, in *TalkRPC.CreateTagRequest) (*TalkRPC.CreateTagResponse, error) {

}

func (cl TalkHandler) DeleteTag(ctx context.Context, in *TalkRPC.DeleteTagRequest) (*TalkRPC.DeleteTagResponse, error) {

}

func (cl TalkHandler) AddTagToContact(context.Context, *TalkRPC.AddTagToContactRequest) (*TalkRPC.AddTagToContactResponse, error) {

}
