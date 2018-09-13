test:
	go test -v .

run:
	go run ArrayStack.go ArrayDeque.go main.go
bench:
	go test -bench . -benchmem
