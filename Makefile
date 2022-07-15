include Makefile.vars

build:
	@echo "==> Building runtime application container using Dockerfile..."
	@docker build -f Dockerfile -t alta3research-go-cert $(QUIET_DOCKER_BUILD) --build-arg APP_NAME=$(APP_NAME) .

run: build
	@echo "==> Entering runtime container & executing application..."
	@docker run -it alta3research-go-cert ./$(APP_NAME)

build_dev: build
	@echo "==> Building development environment container using Dockerfile-dev..."
	@# _dev container is layered ontop of runtime container via Dockerfile-dev FROM
	@docker build -f Dockerfile-dev -t alta3research-go-cert_dev $(QUIET_DOCKER_BUILD) .

run_dev: build_dev
	@echo "==> Entering development environment container - have fun..."
	@# Mounting your users home area inside the container ensures all the conveniences of a configured git, ssh, vim
	@# and other environmental factors are retained
	@# Mounting CWD (i.e. the src tree) into the container ensures container has the latest source code on disk
	@# Using the "-dev" container which is layered on the runtime (non-dev) container means we have all the
	@# dependencies and environment as per the runtime, just with additional stuff.  Same same... just with more.
	@docker run -it \
		-u "$$(id -u ${USER})":"$$(id -g ${USER})" \
		-v "$$(echo ~)":"/home/${USER}" \
		-v "$$(pwd)":/app \
		-v /etc/localtime:/etc/localtime:ro \
		-v /etc/passwd:/etc/passwd:ro \
		-v /etc/group:/etc/group:ro \
		alta3research-go-cert_dev /bin/bash || true

clean: dist-clean test-clean 
	@echo "==> Destroying all container images..."
	-@docker rmi -f alta3research-go-cert > /dev/null 2>&1 || true
	-@docker rmi -f alta3research-go-cert_dev > /dev/null 2>&1 || true

test: build_dev
	@echo "==> Running test suite..."
	@# Mounting CWD (i.e. the src tree) into the container ensures container has the latest source code on disk
	@# ensuring tests are run against current source tree - not a potentially stale container
	@docker run -it -v "$$(pwd)":/app alta3research-go-cert_dev /app/test.sh

test-clean:
	@rm -f ./coverage.out

dist: build
	@echo "==> Extracting built binaries from runtime application container into $(BUILD_DIR)"
	@mkdir -p $(BUILD_DIR)
	@docker run -d -v "$$(pwd)/$(BUILD_DIR)":/$(BUILD_DIR) alta3research-go-cert cp $(APP_NAME) /$(BUILD_DIR)

dist-clean:
	@echo "==> Destroying $(BUILD_DIR)"
	@rm -fr ./$(BUILD_DIR)

rebuild: clean run
