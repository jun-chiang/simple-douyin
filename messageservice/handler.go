package main

import (
	"context"

	"github.com/jun-chiang/simple-douyin/common/consts"
	api "github.com/jun-chiang/simple-douyin/messageservice/kitex_gen/api"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// SendMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendMessage(ctx context.Context, req *api.SendMessageRequest) (resp *api.SendMessageResponse, err error) {
	// TODO: Your code here...

	resp = &api.SendMessageResponse{
		StatusCode: int64(consts.StatusOk),
		StatusMsg:  consts.StatusMsg[consts.StatusOk],
	}
	return
}

// GetMessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageList(ctx context.Context, req *api.GetMessageListRequest) (resp *api.GetMessageListResponse, err error) {
	// TODO: Your code here...
	message_list := make([]*api.Message, 0, 10)
	resp = &api.GetMessageListResponse{
		StatusCode:  int64(consts.StatusOk),
		StatusMsg:   consts.StatusMsg[consts.StatusOk],
		MessageList: message_list,
	}
	return
}
