package pointerutil

import (
	"database/sql"
	"strconv"
	"time"
)

// Bool stores v in a new bool value and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int stores v in a new int value and returns a pointer to it.
func Int(v int) *int { return &v }

// Int32 stores v in a new int32 value and returns a pointer to it.
func Int32(v int32) *int32 { return &v }

// Int64 stores v in a new int64 value and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Uint32 stores v in a new uint32 value and returns a pointer to it.
func Uint32(v uint32) *uint32 { return &v }

// Uint64 stores v in a new uint64 value and returns a pointer to it.
func Uint64(v uint64) *uint64 { return &v }

// Float32 stores v in a new float32 value and returns a pointer to it.
func Float32(v float32) *float32 { return &v }

// Float64 stores v in a new float64 value and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// String stores v in a new string value and returns a pointer to it.
func String(v string) *string { return &v }

// Time stores t in a new time.Time value and returns a pointer to it.
func Time(t time.Time) *time.Time {
	return &t
}

// NullTime returns a pointer to the time value in the sql.NullTime struct if it is valid, otherwise it returns nil.
func NullTime(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

// BigIntString Convert i into a pointer of string.
// TODO: 가능하면 BigIntString 사용하는 대신 BigIntNullString 사용하도록 수정
func BigIntString(i int64) *string {
	s := strconv.FormatInt(i, 10)
	return &s
}

// BigIntNullString takes an int64 value and returns a pointer to a string representation of that value.
// If the value is 0, it returns nil.
func BigIntNullString(i int64) *string {
	if i == 0 {
		return nil
	}
	return BigIntString(i)
}

// IntString Convert i into a pointer of string.
func IntString(i int) *string {
	s := strconv.Itoa(i)
	return &s
}

func PointerBigIntString(i *int64) *string {
	if i == nil {
		return nil
	}
	return BigIntString(*i)
}

func GetTimeValue(value *time.Time, defaultValue time.Time) time.Time {
	if value == nil {
		return defaultValue
	}
	return *value
}

func GetStringValue(value *string, defaultValue string) string {
	if value == nil {
		return defaultValue
	}
	return *value
}

func GetBoolValue(value *bool, defaultValue bool) bool {
	if value == nil {
		return defaultValue
	}
	return *value
}
