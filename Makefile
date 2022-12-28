run: build
	./gogo-db

format:
	go fmt

build: format
	go build
