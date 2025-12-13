package middleware

import (
	"strconv"
)

func ParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}

