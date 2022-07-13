include Makefile.vars

copy_vim_files:
	@if [ ! -d "./.vim" ]; then cp -r -L "$$HOME/.vim" ./.vim; fi
	@if [ ! -f "./.vimrc" ]; then cp -L "$$HOME/.vimrc" ./.vimrc; fi

remove_vim_files:
	@rm -rf ./.vim
	@rm -f ./.vimrc

build:
	@echo "==> Building runtime application container using Dockerfile..."
	@docker build -f Dockerfile -t alta3research-go-cert --build-arg APP_NAME=$(APP_NAME) .

run: build
	@echo "==> Entering runtime container & executing application..."
	@docker run -it alta3research-go-cert ./$(APP_NAME)

build_dev: copy_vim_files
	@echo "==> Building development environment container using Dockerfile-dev..."
	@docker build -f Dockerfile-dev -t alta3research-go-cert_dev .

run_dev: build_dev
	@echo "==> Entering development environment container - have fun..."
	@docker run -it -v "$$(pwd)":/app alta3research-go-cert_dev /bin/bash || true

clean: dist-clean remove_vim_files
	@echo "==> Destroying all container images..."
	-@docker rmi -f alta3research-go-cert > /dev/null 2>&1 || true
	-@docker rmi -f alta3research-go-cert_dev > /dev/null 2>&1 || true

test:
	PYTHONPATH=. pytest

dist: build
	@echo "==> Extracting built binaries from runtime application container into $(BUILD_DIR)"
	@mkdir -p $(BUILD_DIR)
	@docker run -d -v "$$(pwd)/$(BUILD_DIR)":/$(BUILD_DIR) alta3research-go-cert cp $(APP_NAME) /$(BUILD_DIR)

dist-clean:
	@echo "==> Destroying $(BUILD_DIR)"
	@rm -fr ./$(BUILD_DIR)

rebuild: clean run
