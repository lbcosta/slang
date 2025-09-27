BIN_DIR=bin
BIN_NAME=slang
SRC=./src/main.go

$(BIN_DIR)/$(BIN_NAME): $(SRC)
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) $(SRC)

build: $(BIN_DIR)/$(BIN_NAME)

clean:
	rm -rf $(BIN_DIR)
