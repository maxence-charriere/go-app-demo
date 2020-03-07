.PHONY: demo

local:
	GOARCH=wasm GOOS=js go build -o app.wasm ../hello
	go build
	./hello-local

clean:
	@go clean -v ./...
