COVER_OUTPUT =

test:
	go test -v ./... -coverprofile=${COVER_OUTPUT}coverage.out
cover: test
	go tool cover -html=${COVER_OUTPUT}coverage.out
cover_output:
	go tool cover -html=${COVER_OUTPUT}coverage.out -o ${COVER_OUTPUT}coverage.html
bench:
	go test -bench . -benchmem