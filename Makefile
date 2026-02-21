ROOT_DIR:=$(shell pwd)
BIN_DIR:=$(ROOT_DIR)/bin
CMD_DIR=$(ROOT_DIR)/cmd
.PHONY: clean build up run api cli


run-%:clean-% build-%
	@echo "Running: $*"
	$(BIN_DIR)/$*
build-%:
	@echo "Bulding :$*"
	@go build -v -o ${BIN_DIR}/$* $(CMD_DIR)/$*/main.go
clean-%:
	rm -rf $(BIN_DIR)/$*

