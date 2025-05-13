# location of the go.mod file
GOMOD = go.mod

.PHONY: create_go_mod
# create_go_mod prepares the go.mod file
create_go_mod:
	echo "module jwtparservulnerable" > ${GOMOD}
	echo "\ngo 1.24.2" >> ${GOMOD}
	echo "\nrequire github.com/golang-jwt/jwt/v4 v4.5.2" >> ${GOMOD}

.PHONY: use_vulnerable
# use_vulnerable appends the replace directive to use the vulnerable version
use_vulnerable: create_go_mod
	echo "\nreplace github.com/golang-jwt/jwt/v4 => github.com/golang-jwt/jwt/v4 v4.5.1" >> ${GOMOD}

.PHONY: benchmark
# benchmark runs benchmarks with secure code
benchmark: create_go_mod
	go mod tidy && go test -bench=. ./logic -memprofile memprofile.out
	
.PHONY: benchmark_vulnerable
# benchmark_vulnerable runs benchmark with vulnerable code
benchmark_vulnerable: use_vulnerable
	go mod tidy && go test -bench=. ./logic -memprofile memprofile_vulnerable.out

.PHONY: run
# run the Go program
run: create_go_mod
	go mod tidy && time go run .

.PHONY: run_vulnerable
# run the Go program
run_vulnerable: use_vulnerable
	go mod tidy && time go run .

