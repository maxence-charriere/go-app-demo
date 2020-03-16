build:
	@GOARCH=wasm GOOS=js go build -o app.wasm ./app
	@go build -o luck ./server

run: build
	PORT=9000 ./
	
deploy: build
	gcloud app deploy --project=murlok

