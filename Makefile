build:
	go build -o bin/app

run: 
	docker-compose up -d && go run main.go

test:
	go test -v ./... -count=1