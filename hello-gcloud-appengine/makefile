deploy:
	@cd ../hello && make build
	@cp ../hello/app.wasm .
	@-gcloud app deploy . --project go-app-demo-42

run:
	@gcloud app browse --project=go-app-demo-42

clean:
	@go clean
	@-rm app.wasm
