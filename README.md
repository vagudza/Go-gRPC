# fibonacci-grpc
The gRPC API service on Golang, which has one method-which returns a slice from the Fibonacci sequence from the indices x to y.

# Get Fibonachi Slice
The REST API service on Golang, which has one method-which returns a slice from the Fibonacci sequence from the indices `x` to `y`. 
Code have tests.

## Input parameters
`x`, `y` - the integer ordinal indices of the Fibonacci sequence

## Output
The output should list all the numbers, the Fibonacci sequences with ordinal numbers from `x` to `y`    
`{    
  "result": "*The slice of Fibonacci sequence*"    
}`

## Limitations
Conditions under which the service will give the correct answer is:    
`y >= x` and `x, y > 0` and `x, y <= 92`    
The restriction of 92 is due to the overflow of a variable of type int

## Requirements
+ Protocol Buffers
+ gRPC package for Golang
+ `protoc` plugin for Golang

## Examples
For testing the service you need a gRPC client. I suggest using the Evans utility https://github.com/ktr0731/evans (don't forget add path to evans file to system `PATH`. It needs to run in terminal). You can interact with gRPC srever by steps:
+ Run service (see `Quick start`)
+ Run `evans` in terminal `evans api/proto/fibonachi.proto -p 8080`
+ Call Fibonacci procedure `call Fibonachi`
+ Input `x` and `y` correspond to the "Limitations" section
+ Ready! Check response from the service

For example, for `x (TYPE_INT32) => 0` and `y (TYPE_INT32) => 5` service will respond `{    
  "result": "0, 1, 1, 2, 3, 5"    
}`     


## Quick start
+ Install Go (if you haven't already): https://golang.org/doc/tutorial/getting-started
+ Install Protocol Buffers to compile code from a proto-file https://developers.google.com/protocol-buffers?hl=ru
+ Install gRPC package by run in the terminal `go get -u google.golang.org/grpc`
+ (Optional) Install `protoc` plagin for Golang by run in the terminal `go get -u github.com/golang/protobuf/protoc-gen-go`. You can generate your own code of gRPC service from scheme of the service with command (from the project directory) `protoc -I api/proto --go_out=plugins=grpc:pkg/api api/proto/fibonachi.proto`. Code will be generated in `/pkg/api/fibonachi.pb.go`
+ In the project directory run in the terminal: `go run cmd/server/main.go`
+ Ready! Service is running, you can test it (see `Examples`)


