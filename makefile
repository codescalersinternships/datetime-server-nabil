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
	go build ./api/datetimehttphandler/datetime.go

gin_build_binary:
	go build ./api/datetimeginhandler/datetime.go

http_Building the images:
	docker build -t datetimeapphttp -f ./DockerFile-http .

gin_Building the images:
	docker build -t datetimeappgin -f ./DockerFile-gin .

Launching_the_http_containers:
	docker run --name myhttpappname -d -p 8090:8090 datetimeapphttp

Launching_the_gin_containers:
	docker run --name myginappname -d -p 8091:8090 datetimeappgin