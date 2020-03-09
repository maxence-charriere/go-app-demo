# hello

hello is a demo that shows how to use the [go-app package](https://github.com/maxence-charriere/go-app) to build a GUI.

## Build

Go to the hello directory:

```sh
cd $GOPATH/src/github.com/maxence-charriere/go-app-demo/hello
```

```sh
GOARCH=wasm GOOS=js go build -o app.wasm
```

Note that `app.wasm` binary requires to be moved at the server location that will serve it. See the other hello examples:

- [hello-docker](https://github.com/maxence-charriere/go-app-demo/tree/v6/hello-docker)
- [hello-local](https://github.com/maxence-charriere/go-app-demo/tree/v6/hello-local)

## Contribute

Help to develop the [go-app](https://github.com/maxence-charriere/go-app) package by becoming a sponsor.
<br>[Become a sponsor](https://opencollective.com/go-app).
