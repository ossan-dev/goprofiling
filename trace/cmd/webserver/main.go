package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"trace/logic"

	"golang.org/x/exp/trace"
)

func main() {
	// usage of exp FlightRecorder
	fr := trace.NewFlightRecorder()
	fr.SetSize(4096) // optionally set size (or duration) of the circular buffer
	if err := fr.Start(); err != nil {
		panic(err)
	}
	defer fr.Stop()

	// http
	http.HandleFunc("/factorial", func(w http.ResponseWriter, r *http.Request) {
		number, err := strconv.Atoi(r.URL.Query().Get("number"))
		if err != nil {
			number = 1
		}
		result := logic.GetFactorialsUpToN(number)
		if number > 20 { // is a relevant request to trace
			f, err := os.Create("trace_web.out")
			if err != nil {
				panic(err)
			}
			defer f.Close()
			if _, err := fr.WriteTo(f); err != nil {
				panic(err)
			}
		}
		data, _ := json.MarshalIndent(result, "", "\t")
		w.Write(data)
	})

	logic.Result = map[int]int{}
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, "could not start HTTP server,", err.Error())
	}
}
