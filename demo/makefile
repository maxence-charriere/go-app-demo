build:
	@GOARCH=wasm GOOS=js go build -o app.wasm ./app
	@go build -o demo ./server

run: build
	PORT=8000 ./demo
	
deploy: build
	gcloud app deploy --project=murlok

