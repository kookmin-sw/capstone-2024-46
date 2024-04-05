package controller

import (
	"google.golang.org/grpc/codes"

	"private-llm-backend/internal/api"
)

var (
	ErrInvalidArgument = api.NewAPIError(codes.InvalidArgument, "invalid information provided")
)
