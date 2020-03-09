build:
	GOARCH=wasm GOOS=js go build -o app.wasm ../hello
	cp ../hello/app.wasm .
	go build

run: build
	./hello-local

clean:
	@go clean -v ./...
	@-rm app.wasm
