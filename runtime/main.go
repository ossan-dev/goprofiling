package main

import (
	"benchmarking/concatenator"
	"fmt"
	"runtime"
	"strings"
)

func BytesToMegaBytes(b uint64) uint64 {
	return b / 1000 / 1000
}

func PrintMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println()
	fmt.Printf("Alloc = %v MB\n", BytesToMegaBytes(m.Alloc))
	fmt.Printf("TotalAlloc = %v MB\n", BytesToMegaBytes(m.TotalAlloc))
	fmt.Printf("Sys = %v MB\n", BytesToMegaBytes(m.Sys))
	fmt.Printf("Mallocs = %v\n", m.Mallocs)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func main() {
	fmt.Println("=== Start ===")
	PrintMemStats() // track memory allocator statistics
	var result string
	rawResult := make([]string, 0)
	for range 1_000_000 {
		rawResult = append(rawResult, concatenator.Strings("Hello", ", World!"))
	}
	result = strings.Join(rawResult, "|")
	fmt.Println()
	fmt.Println("chars in string:", len(result))
	fmt.Println("\n=== After 1 milion ops ===")
	PrintMemStats()
}
