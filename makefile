APP_NAME=chatgpt-turbo-sample-code

all: clean build

build:
	go build -o $(APP_NAME)

clean:
	rm -f $(APP_NAME)

run:
	./$(APP_NAME)

.PHONY: all build clean run
