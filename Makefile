.PHONY: check
check:
	go mod tidy
	golint
	goimports -w ./
	gofmt -w ./
	go vet

.PHONY: test-vendor
test-vendor:
	go test -mod=vendor -coverprofile=coverage.out -covermode=count ./...

.PHONY: test
test:
	go test -coverprofile=coverage.out -covermode=count ./...
