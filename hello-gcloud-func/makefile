deploy:
	@cd ../hello && make build
	@cp ../hello/app.wasm .
	@gsutil \
    	-h "Cache-Control:no-cache,max-age=0" \
    	-h "Content-Type:application/wasm" \
    	cp ./app.wasm gs://goapp-gcloud-function/app.wasm
	@gcloud functions deploy Hello --runtime go113 --trigger-http --project go-app-demo-42
	@cd ./infra/firebase && firebase deploy

run:
	@open https://go-app-demo-42.firebaseapp.com

clean:
	@go clean
	@-rm app.wasm