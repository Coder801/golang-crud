export GOPATH=$(PWD)

set:
	export GOPATH=$(PWD)

dev:
	go run $(GOPATH)/src/crud/main.go
