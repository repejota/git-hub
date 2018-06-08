include Makefile.help.mk

BINARY=git-hub
MAIN_PACKAGE=cmd/${BINARY}/main.go
PACKAGES = $(shell go list ./...)
VERSION=`cat VERSION`
BUILD=`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:%h -1`
DIST_FOLDER=dist
DIST_INCLUDE_FILES=README.md LICENSE VERSION

# Setup -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Build & Install

install:		## Build and install package on your system
	go install $(LDFLAGS) -v $(PACKAGES)

.PHONY: version
version:		## Show version information
	@echo $(VERSION)-$(BUILD)

# Testing

.PHONY: test
test:			## Execute package tests 
	go test -v $(PACKAGES)

.PHONY: cover-profile
cover-profile:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	rm -rf coverage.out

.PHONY: cover
cover: cover-profile	
cover: 			## Generate test coverage data
	go tool cover -func=coverage-all.out

.PHONY: cover-html
cover-html: cover-profile
cover-html:		## Generate coverage report
	go tool cover -html=coverage-all.out

.PHONY: coveralls
coveralls:
	goveralls -service circle-ci -repotoken l2ZhxGCBuBL37nk0LneaAL5dJsrBET1Py

# Lint

lint:			## Lint source code
	gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow

# Dependencies

deps:			## Install build dependencies
	go get -u github.com/spf13/cobra/cobra
	go get -u golang.org/x/oauth2
	go get -u gopkg.in/src-d/go-git.v4/...
	go get -u github.com/google/go-github/github

dev-deps: deps
dev-deps:		## Install dev and build dependencies
	go get -u github.com/mattn/goveralls
	go get -u github.com/inconshreveable/mousetrap
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: clean
clean:			## Delete generated development environment
	go clean
	rm -rf ${BINARY}
	rm -rf ${BINARY}.exe
	rm -rf coverage-all.out

# Docs

godoc-serve:		## Serve documentation (godoc format) for this package at port HTTP 9090
	godoc -http=":9090"

# Distribution

dist: dist-linux dist-darwin dist-windows
dist:				## Generate distribution packages

dist-linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-linux-amd64
	zip ${BINARY}-${VERSION}-linux-amd64.zip ${BINARY}-${VERSION}-linux-amd64 ${DIST_INCLUDE_FILES}
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-linux-386
	zip ${BINARY}-${VERSION}-linux-386.zip ${BINARY}-${VERSION}-linux-386 ${DIST_INCLUDE_FILES}

dist-darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-darwin-amd64
	zip ${BINARY}-${VERSION}-darwin-amd64.zip ${BINARY}-${VERSION}-darwin-amd64 ${DIST_INCLUDE_FILES}
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-darwin-386
	zip ${BINARY}-${VERSION}-darwin-386.zip ${BINARY}-${VERSION}-darwin-386 ${DIST_INCLUDE_FILES}

dist-windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-windows-amd64.exe
	zip ${BINARY}-${VERSION}-windows-amd64.zip ${BINARY}-${VERSION}-windows-amd64.exe ${DIST_INCLUDE_FILES}
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ${BINARY}-${VERSION}-windows-386.exe
	zip ${BINARY}-${VERSION}-windows-386.zip ${BINARY}-${VERSION}-windows-386.exe ${DIST_INCLUDE_FILES}
