# hello-local-external-root

hello-local is a demo that shows how to run a progressive web app created with the [go-app package](https://github.com/maxence-charriere/go-app) on your local machine. It use the static files of the `hello-local` directory.

## TLDR

```sh
cd $GOPATH/src/github.com/maxence-charriere/go-app-demo/hello-local-external-root
make run
```

## Build and run

Go to the hello-local-external-root directory:

```sh
cd $GOPATH/src/github.com/maxence-charriere/go-app-demo/hello-local-external-root
```

Make sure the `hello-local` directory is built:

```sh
cd ../hello-local && make build && cd -
```

Build the server:

```sh
go build
```

The current directory should look like the following:

```sh
# github.com/maxence-charriere/go-app-demo/hello-local-external-root
.
├── README.md
├── go.mod
├── go.sum
├── hello-local-external-root
└── main.go
```

Run the server:

```sh
./hello-local-external-root
```

## Contribute

Help to develop the [go-app](https://github.com/maxence-charriere/go-app) package by becoming a sponsor.
<br>[Become a sponsor](https://opencollective.com/go-app).
