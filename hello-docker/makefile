build:
	GOARCH=wasm GOOS=js go build -o app.wasm ../hello
	docker build -f dockerfile -t hello-docker ..

run: build
	docker run -p 7000:7000 hello-docker

clean:
	@go clean
	@rm app.wasm
