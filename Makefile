all: run

run:
	GOMAXPROCS=1 go test -v -bench=. -count=3 -run=none -benchmem | tee bench.txt
	# execute `go get -u golang.org/x/perf/cmd/benchstat` if benchstat failed
	benchstat bench.txt

clean:
	rm -rf bench.txt
