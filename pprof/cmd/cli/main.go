package main

import (
	"flag"
	"fmt"
	"os"
	"pprof/logic"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
)

func getParams() (names []string, frequencies []int, err error) {
	pNames := flag.String("names", "john,suzy", "comma-separated list of names to add")
	pFrequency := flag.String("frequency", "100,100", "comma-separated list of the frequency we want to use for each name")
	flag.Parse()
	fmt.Println("names:", *pNames)
	fmt.Println("frequency:", *pFrequency)
	names = strings.Split(*pNames, ",")
	strFrequencies := strings.Split(*pFrequency, ",")
	frequencies = make([]int, len(strFrequencies))
	for k := range frequencies {
		frequency, err := strconv.Atoi(strFrequencies[k])
		if err != nil {
			return nil, nil, err
		}
		frequencies[k] = frequency
	}
	return
}

func main() {
	profile, isSet := os.LookupEnv("profile")
	if isSet && profile == "cpu" {
		// profile CPU
		file, err := os.Create("samples/cpu.out")
		if err != nil {
			fmt.Println("failed to create CPU profiling file:", err.Error())
			return
		}
		defer file.Close()

		if err := pprof.StartCPUProfile(file); err != nil {
			fmt.Println("failed to start CPU profile:", err.Error())
			return
		}
		defer pprof.StopCPUProfile()
	}

	if isSet && profile == "memory" {
		// profile memory
		file, err := os.Create("samples/memory.out")
		if err != nil {
			fmt.Println(fmt.Println("failed to create memory profiling file:", err.Error()))
			return
		}
		defer file.Close()

		defer func() {
			runtime.GC() // force a GC cycle to free memory up
			if err := pprof.WriteHeapProfile(file); err != nil {
				fmt.Println("failed to write heap profile:", err.Error())
				return
			}
		}()
	}

	fmt.Println("working with:")
	names, frequencies, err := getParams()
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range frequencies {
		for range v {
			_ = logic.Append(names[k])
		}
	}
	fmt.Println("savedNames:", logic.GetSavedNames(","))
}
