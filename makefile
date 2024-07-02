server:
	go run ./cmd/main.go

test:
	go test -v -cover ./...

format:
	go fmt ./...

linter:
	golint ./...

build_binary:
	go build ./api/handler/datetime.go

Building the images:
	docker build -t imagename -f ./DockerFile .

Launching the containers:
	docker run --name myappname -d -p 8090:8090 imagename