update:
	@cd hello && \
		go get -u ./... && \
		go mod tidy && \
		GOARCH=wasm GOOS=js go build -o app.wasm
	@cd hello-docker && \
		go get -u ./... && \
		go mod tidy
	@cd hello-gcloud-appengine && \
		go get -u ./... && \
		go mod tidy
	@cd hello-gcloud-func && \
		go get -u ./... && \
		go mod tidy
	@cd hello-local && \
		go get -u ./... && \
		go mod tidy
	@cd hello-local-external-root && \
		go get -u ./... && \
		go mod tidy
	@cd luck && \
		go get -u ./... && \
		go mod tidy
	@cd demo && \
		go get -u ./... && \
		go mod tidy

clean:
	@-rm -r */app.wasm
	@-rm ./luck/luck
	@-rm ./demo/demo
	go clean ./...