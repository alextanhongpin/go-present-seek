# Benchmark

This runs a benchmark for 10 seconds, using 10 threads, and keeping 100 HTTP connections open.

```bash 
$ wrk -t10 -c100 -d10s http://localhost:8080/
```

|---------------------------------|
| language | req/s  | avg latency |
|----------|--------|-------------|
| node     | 724.65 | 13.88ms     |
| express  | 541.46 | 18.60ms     |
| go       | 5.57k  | 2.24ms      |
|---------------------------------|

