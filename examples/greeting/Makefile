serve:
	rm -f webapp/main.wasm
	GOOS=js GOARCH=wasm go build -o webapp/main.wasm webapp/main.go
	go run server/main.go

test:
	go test -tags acceptance
