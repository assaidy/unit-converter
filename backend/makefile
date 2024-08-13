BIN_DIR = bin
CMD_DIR = cmd

all: run

run: build
	@./$(BIN_DIR)/api

build:
	@go build -o $(BIN_DIR)/api $(CMD_DIR)/main.go

clean:
	@rm -rf $(BIN_DIR)

