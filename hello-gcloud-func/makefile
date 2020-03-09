deploy:
	@cp ../hello/app.wasm .
	@gcloud functions deploy Hello --runtime go113 --trigger-http --project go-app-demo-42

clean:
	@go clean
	@-rm app.wasm