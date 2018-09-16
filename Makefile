test:
	go test -v .

run:
	go run main.go
bench:
	go test -bench . -benchmem

profile:
	go test -bench . -benchmem -cpuprofile cpu.out
	go tool pprof --svg OpenDataStructuresGo.test cpu.out > cpu.svg

