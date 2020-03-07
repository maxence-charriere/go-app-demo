.PHONY: demo

local:
	GOARCH=wasm GOOS=js go build -o ./hello-local/app.wasm ./hello
	go build  -o ./hello-local/hello-local ./hello-local
	cd ./hello-local && ./hello-local

clean:
	@go clean -v ./...
