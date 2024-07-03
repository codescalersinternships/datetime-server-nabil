# Simple Datetime server Implemented in Go
Create a basic HTTP server that returns the current date and time. Implement this server using multiple web frameworks, add tests, and containerize the application using Docker and Docker Compose.

# Features

- Endpoint: GET /datetime Respond with Current date and time

# How to Setup

1- import package

```golang
go get -u golang.org/x/lint/golint
```

```golang
go get -u golang.org/x/lint/golint
```

2- create a new parser struct using NewParser()

```golang
parser := NewIniParser()
```

3- load from a file

```golang
_ = parser.LoadFromFile("./testdata/validini.ini")
```

4- load from a string

```golang
_ = parser.LoadFromString(validStringInput)
```

5- get a key value from a section

```golang
val, _ := parser.Get("Simple Values", "key")
fmt.Println(val)
// Output: value
```

6- set a value for a key in a section

```golang
_ = parser.Set("Simple Values", "key", "newvalue")
val, _ := parser.Get("Simple Values", "key")
fmt.Println(val)
// Output: newvalue
```

7- get section names

```golang
sectionsNames, _ := parser.GetSectionNames()
```

8- get parsed data

```golang
sections, _ := parser.GetSections()
```

9- convert data to string

```golang
str, _ := parser.String()
```

10- save data to file

```golang
_ = parser.SaveToFile("./testdata/output.ini")
```

## How to Test

- run go test ./... in root directory

```golang
go test ./...
```

- add the -v flag for more details about the specific tests that are running

```golang
go test -v ./...
```