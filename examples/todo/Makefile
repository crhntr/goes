serve:
	rm -f webapp/main.wasm
	GOOS=js GOARCH=wasm go build -o webapp/main.wasm webapp/main.go

	GOES_EXAMPLES_TODO_DEPENDENCIES_DIR="./dependencies" \
	PORT=8080 \
	 go run server/main.go

test:
	go test -tags acceptance
