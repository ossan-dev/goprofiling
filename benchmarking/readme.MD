# Benchmarking

## perflock

Prepare env:

1. `sudo systemctl stop snapd.service`
2. `sudo systemctl stop unattended-upgrades.service`
3. `sudo systemctl stop thermald`
4. `sudo systemctl stop power-profiles-daemon`
5. `sudo cpupower frequency-set -g performance`
6. `echo 0 | sudo tee /sys/devices/system/cpu/intel_pstate/no_turbo`
7. `sudo cpupower idle-set -D 0`
8. `sudo perflock -governor none -daemon`

Restore env:

1. Kill the `perflock` process
2. `sudo cpupower frequency-set -g powersave`
3. `echo 1 | sudo tee /sys/devices/system/cpu/intel_pstate/no_turbo`
4. `sudo cpupower idle-set -E`
5. `sudo systemctl start thermald`
6. `sudo systemctl start power-profiles-daemon`
7. `sudo systemctl start snapd.service`
8. `sudo systemctl start unattended-upgrades.service`

## benchmark run

```shell
go test -c # save the test binary on FS
time ./concatenator.test -test.bench="." -test.count=2 -test.benchtime=1s | tee buffer.txt

# ... make the code changes ...

go test -c
time ./concatenator.test -test.bench="." -test.count=2 -test.benchtime=1s | tee join.txt
```

## benchcmp

```bash
benchcmp buffer.txt join.txt

benchmark               old ns/op     new ns/op     delta
BenchmarkStrings-12     77.6          45.4          -41.49%
BenchmarkStrings-12     74.7          45.4          -39.18%

benchmark               old allocs     new allocs     delta
BenchmarkStrings-12     2              1              -50.00%
BenchmarkStrings-12     2              1              -50.00%

benchmark               old bytes     new bytes     delta
BenchmarkStrings-12     80            16            -80.00%
BenchmarkStrings-12     80            16            -80.00%
```

## benchstat

```bash
benchstat buffer.txt join.txt

goos: linux
goarch: amd64
pkg: benchmarking/concatenator
cpu: 13th Gen Intel(R) Core(TM) i7-1355U
           │  buffer.txt  │            join.txt             │
           │    sec/op    │    sec/op     vs base           │
Strings-12   76.15n ± ∞ ¹   45.41n ± ∞ ¹  ~ (p=0.333 n=2) ²
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

           │ buffer.txt  │            join.txt            │
           │    B/op     │    B/op      vs base           │
Strings-12   80.00 ± ∞ ¹   16.00 ± ∞ ¹  ~ (p=0.333 n=2) ²
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

           │ buffer.txt  │            join.txt            │
           │  allocs/op  │  allocs/op   vs base           │
Strings-12   2.000 ± ∞ ¹   1.000 ± ∞ ¹  ~ (p=0.333 n=2) ²
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
```
