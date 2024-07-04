# Simple Datetime server Implemented in Go
Create a basic HTTP server that returns the current date and time. Implement this server using multiple web frameworks, add tests, and containerize the application using Docker and Docker Compose.

# Features

- Endpoint: GET /datetime Respond with Current date and time

# How to Setup

1- import package

```golang
go get -u golang.org/x/lint/golint
```

if using server implemented using gin
```golang
go get -u github.com/gin-gonic/gin
go get -u github.com/fvbock/endless
```

# How to Run
To run Server implemented with http standard library
```golang
make httpserver
```
To run Server implemented with gin
```golang
make ginserver
```

and it will run on localhost:8090/datetime

You can also run as a docker container but it will run both servers
but insure that you have docker installed on your machine [Guidence](https://docs.docker.com/)

Using Docker compose
```golang
make Run_dockerCompose
```
when you finnish
```golang
make down_dockerCompose
```
and you can listen for httpserver on localhost:8090/datetime

and ginserver on localhost:8091/datetime

if you want to run only http server as docker image 
```golang
make http_Building_images
make Launching_the_http_containers
```

if you want to run only gin server as docker image 
```golang
make gin_Building_images
make Launching_the_gin_containers
```


when you finnish up delete the images and containers if you want

you can use these to stop http cantainer using this command
```golang
make stop_the_htpp_containers
```
you can use these to stop gin cantainer using this command
```golang
make stop_the_gin_containers
```

you can use these to delete http images using this command
```golang
make delete_the_http_image
```
you can use these to delete gin cantainer using this command
```golang
make delete_the_gin_image
```

## How to Test

```golang
make test
```