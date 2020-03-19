build:
	@cd ../hello && make build
	@cp ../hello/app.wasm .
	@docker build -t hello-docker .

run: build
	docker run -p 7000:7000 hello-docker

clean:
	@go clean
	@rm app.wasm
