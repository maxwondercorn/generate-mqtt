BINARY_NAME=generate_mqtt
 
all: build test
 
build:
	go build -o ../${BINARY_NAME} main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ../${BINARY_NAME} main.go
 
test:
	go test -v main.go
 
run:
	go build -o ${BINARY_NAME} ../main.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}