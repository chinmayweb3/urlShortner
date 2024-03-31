start:
	cd ./server && go run ./cmd/main.go
test:
	cd ./server && go test -v -cover -count=1 ./...