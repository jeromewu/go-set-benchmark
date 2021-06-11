Set benchmark in Golang
=======================

## Execution

```
$ go get -u golang.org/x/perf/cmd/benchstat
$ make
```

## Result

|          | map[]bool | map[]struct{}{} | map[]interface{} |
| -------- | ------------- | ------------------- | -------------------- |
| time/op (execution time) | 3.27s | 3.12s | 5.96s |
| alloc/op (memory consumption) | 884 MB | 802 MB | 1.98 GB |

Using `map[]struct{}` is around 5% faster in time and 10% less memory consumption.
