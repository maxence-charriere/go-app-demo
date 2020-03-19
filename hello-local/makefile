build:
	@cd ../hello && make build
	@cp ../hello/app.wasm .
	@go build

run: build
	./hello-local

clean:
	@go clean -v ./...
	@-rm app.wasm
