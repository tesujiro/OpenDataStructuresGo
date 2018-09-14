test:
	go test -v .

run:
	go run main.go
bench:
	go test -bench . -benchmem
