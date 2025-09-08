run: 
	go run cmd/main.go

dev: 
	go install github.com/air-verse/air@latest
	$GOPATH/bin/air