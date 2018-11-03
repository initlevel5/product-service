PHONY: build fmt test install clean

build: fmt
	go build

fmt:
	git ls-files | grep ".go" | xargs gofmt -l -s -w

test:
	go test

install:
	go install

clean:
	$(RM) product-service
