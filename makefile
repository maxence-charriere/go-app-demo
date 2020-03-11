update:
	@cd hello && \
		go get -u ./... && \
		go mod tidy && \
		GOARCH=wasm GOOS=js go build

clean:
	@rm -r hello-*/app.wasm
	go clean ./...