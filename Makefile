all: test bench

test:
	go test ./algo -v > test.log

bench:
	go test ./algo -bench '.' -benchmem > bench.log

clean:
	rm -f *.log

init: clean
	go mod download
