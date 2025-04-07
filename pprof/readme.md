# pprof

## Instrumentation

Done with the code both for CLI app and for the web server app.
To confirm correct profiling: `http://localhost:8080/debug/pprof/`

## Profile Generation

### CLI

1. `time profile=memory ./cli -names=john,ivan -frequency=10000,10000`: it generates the memory profile

### Web

1. run the application with `go run main.go`
2. issue some requests `curl --location 'http://localhost:8080/append?name=john&times=10000'`
3. generate the profile `curl -v <http://localhost:8080/debug/pprof/allocs> > memory_web.out`

### Benchmarks

1. generate a memory profile with `go test -bench=. -memprofile memory_profile.out`

## Reading Results

## CLI

```text
go tool pprof -alloc_space memory.out

File: cli
Type: alloc_space
Time: Apr 14, 2025 at 3:22pm (CEST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 988.85MB, 100% of 988.85MB total
      flat  flat%   sum%        cum   cum%
  987.64MB 99.88% 99.88%   987.64MB 99.88%  strings.(*Builder).grow
    1.21MB  0.12%   100%   988.85MB   100%  pprof/logic.Append (inline)
         0     0%   100%   988.85MB   100%  main.main
         0     0%   100%   988.85MB   100%  runtime.main
         0     0%   100%   987.64MB 99.88%  strings.(*Builder).Grow
         0     0%   
```

## Graphical Report

```bash
# Prerequisite => "sudo apt install graphviz"

go tool pprof -svg -output memory.svg -alloc_space memory.out

# Generating report in memory.svg
```

## Web-UI

```bash
go tool pprof -http=:8083 -alloc_space memory.out
```
