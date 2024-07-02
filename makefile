httpserver:
	go run ./datetimehttp/cmd/main.go

httpserver:
	go run ./datetimegin/cmd/main.go

test:
	go test -v -cover ./...

format:
	go fmt ./...

linter:
	golint ./...

http_build_binary:
	go build ./datetimehttp/api/handler/datetime.go

gin_build_binary:
	go build ./datetimegin/api/handler/datetime.go

http_Building the images:
	docker build -t httpimagename -f ./datetimehttp/DockerFile .

gin_Building the images:
	docker build -t ginimagename -f ./datetimegin/DockerFile .

Launching_the_containers:
	docker run --name myhttpappname -d -p 8090:8090 httpimagename

Launching_the_containers:
	docker run --name myginappname -d -p 8091:8090 ginimagename