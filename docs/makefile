build:
	go run main.go
	@cd ../hello && make build
	@cp ../hello/app.wasm ./web/

clean:
	@go clean
	@-rm app.wasm
