test:
	go test -v -race -covermode=atomic -coverprofile=coverage.out ./...
