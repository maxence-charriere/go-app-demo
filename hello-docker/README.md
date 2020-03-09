# hello-docker

hello-docker is a demo that shows how to deploy a progressive web app created with the [app package](https://github.com/maxence-charriere/go-app) in a Docker container.

## Prerequisites

- [Docker](https://www.docker.com) installed on your machine
- A [Docker file](https://github.com/maxence-charriere/go-app-demo/tree/v6/hello-docker/dockerfile)

## Build and run Docker contrainer

Go to the hello-docker directory:

```sh
cd $GOPATH/src/github.com/maxence-charriere/go-app-demo/hello-docker
```

Build the hello app:

```sh
GOARCH=wasm GOOS=js go build -o app.wasm ../hello
```

The current directory should look like the following:

```sh
# github.com/maxence-charriere/go-app-demo/hello-docker
.
├── README.md
├── app.wasm
├── dockerfile
└── main.go
```

Build Docker image:

```sh
docker build -f dockerfile -t hello-docker ..
```

Run the Docker container:

```sh
docker run -p 7000:7000 hello-docker
```

## Contribute

Help to develop the [go-app](https://github.com/maxence-charriere/go-app) package by becoming a sponsor.
<br>[Become a sponsor](https://opencollective.com/go-app).
