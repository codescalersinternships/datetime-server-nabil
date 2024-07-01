server:
	go run main.go

test:
	go test -v -cover ./...

format:
	go fmt ./...

linter:
	golint ./...

build_binary:
	go build ./api/handler/datetime.go

