test:
	@go test ./...

test-cover:
	@go test ./... -cover

coverage:
	@go test ./... -coverprofile=cover.out
	@go tool cover -html=cover.out

clean:
	@rm cover.out ||:
