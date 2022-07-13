copy_vim_files:
	@if [ ! -d "./.vim" ]; then cp -r -L "$$HOME/.vim" ./.vim; fi
	@if [ ! -f "./.vimrc" ]; then cp -L "$$HOME/.vimrc" ./.vimrc; fi

remove_vim_files:
	@rm -rf ./.vim
	@rm -f ./.vimrc

build:
	@docker build -f Dockerfile -t alta3research-go-cert .

run: build
	@docker run -it alta3research-go-cert ./main

build_dev: copy_vim_files
	@docker build -f Dockerfile-dev -t alta3research-go-cert_dev .

run_dev: build_dev
	@docker run -it -v "$$(pwd)":/app alta3research-go-cert_dev /bin/bash

clean: dist-clean remove_vim_files
	-@docker rmi -f alta3research-go-cert > /dev/null 2>&1 || true
	-@docker rmi -f alta3research-go-cert_dev > /dev/null 2>&1 || true

test:
	PYTHONPATH=. pytest

dist: build
	@mkdir -p dist
	@docker run -d -v "$$(pwd)/dist":/dist alta3research-go-cert cp main /dist

dist-clean:
	@rm -fr ./dist

rebuild: clean run
