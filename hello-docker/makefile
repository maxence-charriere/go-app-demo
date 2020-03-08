deps:
	-go mod init
	go get -u github.com/maxence-charriere/go-app@v1.0.1
	@go mod tidy

build: deps
	cd ../hello && \
		GOARCH=wasm GOOS=js go build -o app.wasm && \
		mv app.wasm ../hello-docker/app.wasm
	docker build -t hello-docker .

run: build
	docker run -p 7000:7000 hello-docker

clean:
	@go clean
	@rm go.mod
	@rm go.sum
	@rm app.wasm
