package main

import (
	"fmt"
	"net/http"
	"pprof/logic"
	"strconv"

	_ "net/http/pprof"
)

func appendName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "john"
	}
	times, err := strconv.Atoi(r.URL.Query().Get("times"))
	if err != nil {
		times = 1000
	}
	for range times {
		_ = logic.Append(name)
	}
	fmt.Fprintln(w, "savedNames:", logic.GetSavedNames(","))
}

func main() {
	http.HandleFunc("/append", appendName)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
