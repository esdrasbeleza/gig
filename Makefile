install-deps:
	go mod download
	go get -t github.com/esdrasbeleza/gig

update-submodules:
	git submodule update --init --recursive

setup: install-deps update-submodules 

build:
	go build .

clean:
	rm -f gig

test-ci: setup build
	go test ./...

	
