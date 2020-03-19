# hello-local

hello-local is a demo that shows how to run a progressive web app created with the [go-app package](https://github.com/maxence-charriere/go-app) on your local machine.

## TLDR

```sh
cd $GOPATH/src/github.com/maxence-charriere/go-app-demo/hello-local
make run
```

## Build and run

Go to the hello-local directory:

```sh
cd $GOPATH/src/github.com/maxence-charriere/go-app-demo/hello-local
```

Make sure the `hello` directory is built:

```sh
cd ../hello && make build && cd -
```

Copy the hello wasm binary:

```sh
cp ../hello/app.wasm .
```

Build the server:

```sh
go build
```

The current directory should look like the following:

```sh
# github.com/maxence-charriere/go-app-demo/hello-local
.
├── README.md
├── app.wasm
├── go.mod
├── go.sum
├── hello-local
└── main.go
```

Run the server:

```sh
./hello-local
```

## Contribute

Help to develop the [go-app](https://github.com/maxence-charriere/go-app) package by becoming a sponsor.
<br>[Become a sponsor](https://opencollective.com/go-app).
