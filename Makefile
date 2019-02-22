install-packr2:
	go get -u github.com/gobuffalo/packr/v2/...

install-deps:
	dep ensure

update-submodules:
	git submodule update --init --recursive

setup: install-packr2 install-deps update-submodules 

build:
	packr2 build

install: build
	packr2 install

clean:
	packr2 clean
	rm -rf vendor/
	rm -f gig
