serve: assets
	mkdir -p assets
	rm -f assets/main.wasm assets/wasm_exec.js
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" assets/

	GOOS=js GOARCH=wasm go build -o assets/main.wasm webapp/main.go
	go run server/main.go

test:
	go test -tags acceptance
