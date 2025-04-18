package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/trace"
	"trace/logic"
)

var p_number *int

func init() {
	p_number = flag.Int("number", 10, "the last number you want the factorial. Maximum value allowed is 20")
	flag.Parse()
	logic.Result = map[int]int{}
}

func main() {
	// trace your program
	f, err := os.Create("trace_cli.out")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer trace.Stop()

	// program to be traced
	results := logic.GetFactorialsUpToN(*p_number)
	for k, v := range results {
		fmt.Println("k:", k+1, "v:", v)
	}
}
