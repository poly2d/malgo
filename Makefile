run:
	go run mal.go
dep:
	go get github.com/stretchr/testify/assert
fmt:
	go fmt ./...
test:
	go test ./read
	go test ./eval
