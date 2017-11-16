all: clean vet lint test

test:
	sh test.sh

clean:
	go clean ./...

vet:
	go vet ./...

lint:
	golint $(go list ./...)
