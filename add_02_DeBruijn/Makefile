all:
	go build .

test:
	go test -v .

bench:
	go test -bench . -benchmem

prof:
	go test -bench . -benchmem -cpuprofile cpu.out
	go tool pprof --svg add_02_DeBruijn.test cpu.out > cpu.svg
