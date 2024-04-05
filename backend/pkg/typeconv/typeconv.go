package typeconv

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

	"private-llm-backend/pkg/errorutil"
	"private-llm-backend/pkg/pointerutil"
)

func AtoI64(s string) (int64, error) {
	if s == "" {
		return 0, errorutil.Error(errors.New("failed to convert empty string to int64"))
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errorutil.WithDetail(err, fmt.Errorf("failed to convert string to int64: %s", s))
	}
	return i, nil
}

func I64toA(i int64) string {
	return strconv.FormatInt(i, 10)
}

func PStringToInt64(s *string) (int64, error) {
	if s == nil {
		return 0, nil
	}
	i, err := AtoI64(*s)
	if err != nil {
		return 0, errorutil.Error(err)
	}
	return i, nil
}

func PStringToPInt64(s *string) (*int64, error) {
	if s == nil {
		return nil, nil
	}
	i, err := PStringToInt64(s)
	if err != nil {
		return nil, errorutil.Error(err)
	}
	return &i, nil
}

func ToDeleteTime(value *time.Time) *gorm.DeletedAt {
	if value == nil {
		return nil
	}
	return &gorm.DeletedAt{
		Time:  *value,
		Valid: true,
	}
}

// DeleteTime returns a pointer to time.Time from gorm.DeletedAt.
func DeleteTime(t *gorm.DeletedAt) *time.Time {
	if t == nil {
		return nil
	}
	return pointerutil.NullTime(sql.NullTime(*t))
}
