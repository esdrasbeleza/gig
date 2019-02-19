packr2:
	go get -u github.com/gobuffalo/packr/v2/...

deps:
	dep ensure

setup: packr2 deps 

build:
	packr2 build

install:
	packr2 install

clean:
	packr2 clean
	rm -rf vendor/
	rm -f gig
