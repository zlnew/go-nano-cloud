APP_NAME=mini-s3
ENTRY=./cmd/api

build:
	go build -o bin/$(APP_NAME) $(ENTRY)

run:
	go run $(ENTRY)

clean:
	rm -rf bin/$(APP_NAME)

