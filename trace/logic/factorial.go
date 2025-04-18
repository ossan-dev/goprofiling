package logic

import (
	"time"
)

const N_MAX = 50

var Result map[int]int

func GetFactorialsUpToN(n int) map[int]int {
	if n < 1 {
		n = 1
	}
	if n > N_MAX {
		n = N_MAX
	}
	for i := range n {
		Result[i+1] = Factorial(i + 1)
	}
	return Result
}

func Factorial(n int) int {
	result := 1
	for i := range n {
		result *= i + 1
	}
	time.Sleep(time.Millisecond * 100)
	return result
}
