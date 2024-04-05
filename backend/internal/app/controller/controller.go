package controller

import (
	"private-llm-backend/internal/api"
	"private-llm-backend/internal/app/application/chat"
)

// API implements the api.StrictServerInterface
var _ api.StrictServerInterface = (*API)(nil)

type API struct {
	chatService chat.Service
}

func NewAPI(chatService chat.Service) api.StrictServerInterface {
	return &API{
		chatService: chatService,
	}
}
