package controller

import (
	"context"
	"errors"

	"private-llm-backend/internal/api"
	"private-llm-backend/pkg/errorutil"
)

func (a *API) PlaygroundChatServiceCreateChatConversation(ctx context.Context, request api.PlaygroundChatServiceCreateChatConversationRequestObject) (api.PlaygroundChatServiceCreateChatConversationResponseObject, error) {
	if len(request.Body.Messages) == 0 || request.Body.ThreadId == "" {
		return &api.PlaygroundChatServiceCreateChatConversationdefaultJSONResponse{
			Body:       api.NewErrorBody(ErrInvalidArgument),
			StatusCode: 400,
		}, nil
	}
	lastMessage := request.Body.Messages[len(request.Body.Messages)-1]
	resp, err := a.chatService.Conversation(ctx, request.Body.ThreadId, lastMessage.Content)
	if err != nil {
		err = errorutil.WithDetail(err, errors.New("failed to create chat conversation"))
		return nil, err
	}
	return api.PlaygroundChatServiceCreateChatConversation200JSONResponse(*resp), nil
}

func (a *API) PlaygroundChatServiceCreateThread(ctx context.Context, request api.PlaygroundChatServiceCreateThreadRequestObject) (api.PlaygroundChatServiceCreateThreadResponseObject, error) {
	newThread, err := a.chatService.CreateThread(ctx)
	if err != nil {
		err = errorutil.WithDetail(err, errors.New("failed to create thread"))
		return nil, err
	}
	return &api.PlaygroundChatServiceCreateThread200JSONResponse{
		ThreadId: newThread.ThreadID,
	}, nil
}
