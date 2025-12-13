package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

func StringSliceToJSON(slice []string) string {
	if len(slice) == 0 {
		return "[]"
	}
	data, _ := json.Marshal(slice)
	return string(data)
}

func JSONToStringSlice(s string) []string {
	var slice []string
	if s == "" || s == "[]" {
		return slice
	}
	json.Unmarshal([]byte(s), &slice)
	return slice
}

func NullStringToStringPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func StringPtrToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

func NullFloat64ToFloat64Ptr(nf sql.NullFloat64) *float64 {
	if !nf.Valid {
		return nil
	}
	return &nf.Float64
}

func Float64PtrToNullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{Valid: false}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

func ParseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

