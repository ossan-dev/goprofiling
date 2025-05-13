# CVE-2025-30204

## Benchmark Vulnerable

```text
go mod tidy && go test -bench=. ./logic -memprofile memprofile_vulnerable.out
goos: linux
goarch: amd64
pkg: jwtparservulnerable/logic
cpu: 13th Gen Intel(R) Core(TM) i7-1355U
BenchmarkParseUnverified-12         6644            169807 ns/op          163906 B/op          3 allocs/op
PASS
ok      jwtparservulnerable/logic       1.952s
```

## Benchmark Fixed

```text
go mod tidy && go test -bench=. ./logic -memprofile memprofile.out
goos: linux
goarch: amd64
pkg: jwtparservulnerable/logic
cpu: 13th Gen Intel(R) Core(TM) i7-1355U
BenchmarkParseUnverified-12     13154956               112.8 ns/op           112 B/op          3 allocs/op
PASS
ok      jwtparservulnerable/logic       2.247s
```

## Run Vulnerable

```text
go mod tidy && time go run .

 === Before ===
Alloc = 0 MB
TotalAlloc = 0 MB
Sys = 6 MB
Mallocs = 399
NumGC = 0

failed to parse token token contains an invalid number of segments

 === After ===
Alloc = 18000 MB
TotalAlloc = 18000 MB
Sys = 18028 MB
Mallocs = 497
NumGC = 2

13.13user 18.24system 0:58.14elapsed 53%CPU (0avgtext+0avgdata 13731972maxresident)k
28797168inputs+24outputs (529403major+7489351minor)pagefaults 0swaps
```

## Run Fixed

```text
go mod tidy && time go run .

 === Before ===
Alloc = 0 MB
TotalAlloc = 0 MB
Sys = 6 MB
Mallocs = 398
NumGC = 0

failed to parse token token contains an invalid number of segments

 === After ===
Alloc = 2000 MB
TotalAlloc = 2000 MB
Sys = 2009 MB
Mallocs = 455
NumGC = 1

1.04user 1.08system 0:00.93elapsed 226%CPU (0avgtext+0avgdata 1959960maxresident)k
88224inputs+24outputs (623major+510982minor)pagefaults 0swaps
```
