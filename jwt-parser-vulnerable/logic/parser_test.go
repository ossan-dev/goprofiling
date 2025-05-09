package logic_test

import (
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

var (
	g_token *jwt.Token
	g_parts []string
	g_err   error
)

func BenchmarkParseUnverified(b *testing.B) {
	// setup
	var token *jwt.Token
	var parts []string
	var err error
	parser := jwt.NewParser()
	rawToken := strings.Repeat("1.", 10_000)
	claims := jwt.MapClaims{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		token, parts, err = parser.ParseUnverified(rawToken, &claims)
	}
	g_token = token
	g_parts = parts
	g_err = err
}
