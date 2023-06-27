
NAME=app

all: build run

run:
	./$(NAME)
clean:
	rm -rf $(NAME)

build:
	go build -o $(NAME) ./cmd

test:
	go test -v -count=1 ./...

race:
	go test -v -race -count=1 ./...

.PHONY: cover
cover:
	go test ./... -short -coverprofile=coverage.out
	go tool cover -html=covarage.out
