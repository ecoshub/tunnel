NAME="tunnel"

build: build-mac build-linux

build-mac:
	@echo "building..."
	@go mod tidy
	@go build -o $(NAME) .
	@echo "done."
	@echo "executable created ./$(NAME)"

build-linux:
	@echo "building for linux..."
	@go mod tidy
	@GOOS=linux go build -o $(NAME)-linux .
	@echo "done."
	@echo "executable created ./$(NAME)-linux"

help:
	@echo '-----------------------'
	@echo '|  welcome to tunnel  |'
	@echo '-----------------------'
	@echo ''
	@echo 'example usage:'
	@echo '  $$ make build'
	@echo ''
	@echo 'options:'
	@echo '  - help .........: Print this help dialog.'
	@echo '  - build ........: Build the package. create executable for "mac" and "linux" in the root dir.'
	@echo '  - build-mac ....: Build the package for "macos". create executable in the root dir.'
	@echo '  - build-linux ..: Build the package for "linux". create executable in the root dir.'


.DEFAULT_GOAL := help
