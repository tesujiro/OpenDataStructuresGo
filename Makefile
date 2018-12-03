test:
	go vet
	go test -v . -coverpkg ./...

cover:
	go test -v . --coverpkg ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

run:
	go run main.go
bench:
	go test -bench . -benchmem

profile:
	go test -bench . -benchmem -cpuprofile cpu.out
	go tool pprof --svg OpenDataStructuresGo.test cpu.out > cpu.svg
