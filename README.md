# alta3research-go-cert

---

This application is a submission for certification for the Alta3 Research GoLang proficiency course.

For convenience the runtime distribution a development environment are wrapped up in containers.
Runtime distribution: Dockerfile
Development environment: Dockerfile-dev

These are optional conveniences which remove any runtime / build dependency, toolchain versioning
and other environmental faffing you would normally have to do to compile from src.
You can choose to use these or just interact with the sources directly on your host.

Additionally there is a Makefile wrapper to interact with these two containers.

&nbsp;
## Usage

---
1. Recommended approach: use the Makefile convenience wrapper:
	$ make dist
	$ ./dist/swisscheese

2. Enter the runtime distribution container:
	2.1. Using the Makefile convenience wrapper:
		$ make run
 		$ ./swisscheese
	2.2. Build out the container from sources:
		$ docker build -f Dockerfile -t alta3research-go-cert .	
		$ docker run -it alta3research-go-cert
		$ ./swisscheese

3. Manually build the application source directly on your host machine - if this is the method you choose, you're on your own.

&nbsp;
## Test suite

---
Running the test suite is wrapped up in test.sh which perform all the necessary steps.

1. Recommended approach: use the Makefile convenience wrapper:
	$ make test

2. If you have a working environment to build and run the project on your host machine then simply invoke:
	$ ./test.sh

&nbsp;
## Contributing

---

A container based development environment is used to maintain this project.

1. Recommended approach: use the Makefile convenience wrapper:
	$ make run_dev
	$ ... profit

2. Build out the container from sources:
	$ docker build -f Dockerfile-dev -t alta3research-go-cert_dev.
	$ docker run -it alta3research-go-cert_dev
