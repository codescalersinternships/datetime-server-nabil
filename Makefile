httpserver:
	go run ./cmd/datetimehttp/main.go

ginserver:
	go run ./cmd/datetimegin/main.go

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

http_Building_images:
	docker build -t datetimeapphttp -f ./DockerFile-http .

Launching_the_http_containers:
	docker run --name myhttpappname -d -p 8090:8090 datetimeapphttp

stop_the_htpp_containers:
	docker stop myhttpappname

delete_the_http_image:
	docker image rm datetimeapphttp -f

gin_Building_images:
	docker build -t datetimeappgin -f ./DockerFile-gin .

Launching_the_gin_containers:
	docker run --name myginappname -d -p 8091:8090 datetimeappgin

stop_the_gin_containers:
	docker stop myginappname

delete_the_gin_image:
	docker image rm datetimeappgin -f

Run_dockerCompose:
	docker-compose up -d

down_dockerCompose:
	docker-compose down