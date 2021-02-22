COVER_OUTPUT =

test:
	go test -v ./... -coverprofile=${COVER_OUTPUT}coverage.out
cover: test
	go tool cover -html=coverage.out
bench:
	go test -bench . -benchmem