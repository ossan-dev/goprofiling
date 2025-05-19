package logic_test

import (
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

var (
	Token *jwt.Token
	Parts []string
	Err   error
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
	Token = token
	Parts = parts
	Err = err
}

func BenchmarkParseUnverified_1_24(b *testing.B) {
	parser := jwt.NewParser()
	rawToken := strings.Repeat("1.", 10_000)
	claims := jwt.MapClaims{}
	b.ReportAllocs()
	for b.Loop() {
		parser.ParseUnverified(rawToken, &claims)
	}
}
