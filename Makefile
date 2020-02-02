VERSION=1.0.0

all: fmt  release

release:  release-deps 
	gox -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}" .

fmt:
	go fmt ./...

release-deps:
	go get github.com/mitchellh/gox

.PHONY: all  release fmt release-deps
