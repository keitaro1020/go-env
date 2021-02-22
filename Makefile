test:
	go test -v ./... -coverprofile=coverage.out
cover: test
	go tool cover -html=coverage.out
bench:
	go test -bench . -benchmem