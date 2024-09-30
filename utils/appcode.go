package utils

import (
	"math/rand"
	"time"
)

const (
	appCodeLength = 6
	appCodeChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateAppCode creates a 6-digit alphanumeric code
func GenerateAppCode() string {
	code := make([]byte, appCodeLength)
	for i := range code {
		code[i] = appCodeChars[rand.Intn(len(appCodeChars))]
	}
	return string(code)
}
