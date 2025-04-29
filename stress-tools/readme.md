# Stress-test Tools

## Apache Benchmark - AB

```bash
ab -c 10 -n 1000 -l \
"http://localhost:8080/append?name=john&times=10"
# '-c' means the no of requests per time. Default is one
# '-n' means the no of requests per session. Default is one
# '-r' means to not exit on socket receive errors
# '-l' means accept variable document length in responses
```

```text
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>

Copyright 1996 Adam Twiss, Zeus Technology Ltd, <http://www.zeustech.net/>
Licensed to The Apache Software Foundation, <http://www.apache.org/>

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests

Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /append?name=john&times=10
Document Length:        Variable

Concurrency Level:      10
Time taken for tests:   0.305 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      24984404 bytes
HTML transferred:       24886545 bytes
Requests per second:    3281.73 [#/sec] (mean)
Time per request:       3.047 [ms] (mean)
Time per request:       0.305 [ms] (mean, across all concurrent requests)
Transfer rate:          80070.47 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       1
Processing:     0    3   1.1      3       7
Waiting:        0    2   1.0      2       6
Total:          1    3   1.1      3       7

Percentage of the requests served within a certain time (ms)
  50%      3
  66%      3
  75%      4
  80%      4
  90%      4
  95%      5
  98%      5
  99%      6
 100%      7 (longest request)
```

## hey

```bash
hey -c 10 -n 1000 -cpus=12 -h2 \
"http://localhost:8080/append?name=john&times=10"
# '-c', '-n' same meaning as before
# '-cpus' number of used CPU cores
# '-h2' enables HTTP/2
```

```text
Summary:
  Total:        0.2685 secs
  Slowest:      0.0097 secs
  Fastest:      0.0003 secs
  Average:      0.0026 secs
  Requests/sec: 3724.8691
  
  Total data:   44167 bytes
  Size/request: 44 bytes

Response time histogram:
  0.000 [1]     |
  0.001 [110]   |■■■■■■■■■■■■
  0.002 [229]   |■■■■■■■■■■■■■■■■■■■■■■■■■
  0.003 [364]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [205]   |■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [66]    |■■■■■■■
  0.006 [14]    |■■
  0.007 [3]     |
  0.008 [5]     |■
  0.009 [2]     |
  0.010 [1]     |


Latency distribution:
  10% in 0.0011 secs
  25% in 0.0019 secs
  50% in 0.0026 secs
  75% in 0.0033 secs
  90% in 0.0040 secs
  95% in 0.0045 secs
  99% in 0.0065 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0003 secs, 0.0097 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0008 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0024 secs
  resp wait:    0.0024 secs, 0.0002 secs, 0.0075 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0052 secs

Status code distribution:
  [200] 1000 responses
```
