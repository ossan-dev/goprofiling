# Trace

## Generation

1. `curl --location http://localhost:8080/debug/pprof/trace?seconds=20 > trace_20_secs.out` => it generates a 20-second long trace. The web server must be instrumented with the import for side effects `_ "/net/http/pprof"`

## Reading

```shell
# it launches a Chrome tab on a random port
go tool trace trace_cli.out
```

1. `go tool trace trace_cli.out` => it will open a Chrome page
2. `View trace by proc` => it will lead you to the graph to look at
