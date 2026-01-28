APP_NAME=secret-cli
CMD_PATH=./cmd/secret-cli/main.go


run:clean build
	go run $(CMD_PATH)
build:clean
	go build -o bin/$(APP_NAME) $(CMD_PATH)
clean:
	rm -rf bin
