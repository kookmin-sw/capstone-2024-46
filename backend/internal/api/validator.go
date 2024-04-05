package api

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"

	"private-llm-backend/internal/database"
	"private-llm-backend/pkg/errorutil"
)

type keyIterator interface {
	Keys() map[string]struct{}
}

var nameConverter = database.NameConverter

func toColumnName(fieldName string) string {
	return nameConverter.ColumnName("", fieldName)
}

func ValidateUpdateMask(updateMask *string, body keyIterator) ([]string, error) {
	if updateMask == nil || *updateMask == "" {
		err := errorutil.WithDetail(
			ErrInvalidArgument,
			NewAPIError(codes.InvalidArgument, "an updateMask is needed to specify which fields to update. please include an updateMask and try again"),
		)
		return nil, err
	}
	keys := body.Keys()
	fields := strings.Split(*updateMask, ",")
	snakeCaseMask := make([]string, 0, len(fields))
	for _, fieldName := range fields {
		if _, ok := keys[fieldName]; !ok {
			err := errorutil.WithDetail(
				ErrInvalidArgument,
				NewAPIError(codes.InvalidArgument, fmt.Sprintf("the updateMask for '%s' is invalid. please check the field name and format of your updateMask and try again", fieldName)),
			)
			return nil, err
		}
		snakeCaseMask = append(snakeCaseMask, toColumnName(fieldName))
	}
	return snakeCaseMask, nil
}
