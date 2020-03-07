build:
	GOARCH=wasm GOOS=js go build -o app.wasm ../hello
	go build

run: build
	./hello-local

clean:
	@go clean -v ./...
	@go rm app.wasm
