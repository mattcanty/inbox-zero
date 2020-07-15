build:
	go test -v ./...
	go build
install: build
	go install
run:
	go build
	./inbox-zero