package chat

import (
	"context"
	"errors"
	"fmt"
	"time"

	"private-llm-backend/internal/api"
	"private-llm-backend/pkg/client/openai"
	"private-llm-backend/pkg/errorutil"
)

var _ Service = (*chatService)(nil)

type Service interface {
	CreateThread(ctx context.Context) (*Thread, error)
	Conversation(ctx context.Context, threadID string, messages string) (*api.CreateChatConversationResponse, error)
}

type chatService struct {
	assistantID string
	client      openai.ClientWithResponsesInterface
}

func (c *chatService) retrieveRunMessage(ctx context.Context, threadID string, runID string) (*openai.Run, error) {
	resp, err := c.client.ThreadServiceRetrieveRunWithResponse(
		ctx,
		threadID,
		runID,
		&openai.ThreadServiceRetrieveRunParams{
			OpenAIBeta: openai.AssistantsBetaV1,
		},
	)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to get response"))
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}
	if resp.JSON200 == nil {
		return nil, errors.New("unexpected response body")
	}
	return resp.JSON200, nil
}

func (c *chatService) listThreadMessage(ctx context.Context, threadID string) ([]openai.ThreadMessage, error) {
	resp, err := c.client.ThreadServiceListMessagesWithResponse(ctx, threadID, &openai.ThreadServiceListMessagesParams{
		OpenAIBeta: openai.AssistantsBetaV1,
	})
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to get response"))
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}
	if resp.JSON200 == nil {
		return nil, errors.New("unexpected response body")
	}
	return *resp.JSON200.Data, nil
}

func (c *chatService) createChat(ctx context.Context, threadID string, messages string) error {
	resp, err := c.client.ThreadServiceCreateMessageWithResponse(
		ctx,
		threadID,
		&openai.ThreadServiceCreateMessageParams{
			OpenAIBeta: openai.AssistantsBetaV1,
		},
		openai.CreateMessageRequest{
			Content: messages,
			Role:    openai.CreateMessageRequestRoleUser,
		},
	)
	if err != nil {
		return errorutil.WithDetail(err, errors.New("failed to get response"))
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}
	return nil
}

func (c *chatService) createResponse(ctx context.Context, threadID string) (*openai.Run, error) {
	resp, err := c.client.ThreadServiceCreateRunWithResponse(
		ctx,
		threadID,
		&openai.ThreadServiceCreateRunParams{
			OpenAIBeta: openai.AssistantsBetaV1,
		},
		openai.ThreadServiceCreateRunJSONRequestBody{
			AssistantId: c.assistantID,
		},
	)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to get response"))
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}
	if resp.JSON200 == nil {
		return nil, errors.New("unexpected response body")
	}
	return resp.JSON200, nil
}

func (c *chatService) Conversation(ctx context.Context, threadID string, message string) (*api.CreateChatConversationResponse, error) {
	err := c.createChat(ctx, threadID, message)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to create chat"))
	}
	runResult, err := c.createResponse(ctx, threadID)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to create response"))
	}
	runID := runResult.Id

	for i := 0; i < 100; i++ {
		time.Sleep(200 * time.Millisecond)
		runResult, err = c.retrieveRunMessage(ctx, threadID, runID)
		if err != nil {
			return nil, errorutil.WithDetail(err, errors.New("failed to retrieve run message"))
		}
		if runResult.Status == openai.Completed {
			break
		}
	}
	if runResult.Status != openai.Completed {
		return nil, errorutil.WithDetail(err, errors.New("maximum retries exceeded"))
	}

	messages, err := c.listThreadMessage(ctx, threadID)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to list thread message"))
	}
	var runMessage *openai.ThreadMessage
	for _, threadMessage := range messages {
		if threadMessage.RunId == nil {
			continue
		}
		if *threadMessage.RunId == runID {
			runMessage = &threadMessage
			break
		}
	}
	if runMessage == nil {
		return nil, errorutil.WithDetail(err, errors.New("run message not found"))
	}

	if runMessage.Content == nil {
		return nil, errorutil.WithDetail(err, errors.New("nil message content"))
	}
	contents := *runMessage.Content
	if len(contents) == 0 {
		return nil, errorutil.WithDetail(err, errors.New("empty message content"))
	}

	text := ""
	if contents[0].Text != nil {
		text = contents[0].Text.Value
	}

	return &api.CreateChatConversationResponse{
		CreateTime: time.Unix(int64(runMessage.CreatedAt), 0),
		Id:         runMessage.Id,
		Message: api.ConversationMessage{
			Content: text,
			Role:    api.ConversationMessageRole(runMessage.Role),
		},
		ObjectType:        api.CreateChatConversationResponseObjectTypeChatCompletion,
		SystemFingerprint: nil,
	}, nil
}

func (c *chatService) CreateThread(ctx context.Context) (*Thread, error) {
	resp, err := c.client.ThreadServiceCreateThreadWithResponse(ctx, &openai.ThreadServiceCreateThreadParams{
		OpenAIBeta: openai.AssistantsBetaV1,
	})
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to get response"))
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}
	if resp.JSON200 == nil {
		return nil, errors.New("unexpected response body")
	}
	return &Thread{
		ThreadID: resp.JSON200.Id,
	}, nil
}

func NewChatService(assistantID string, openaiClient openai.ClientWithResponsesInterface) Service {
	return &chatService{
		assistantID: assistantID,
		client:      openaiClient,
	}
}
