# hello-gcloud-func

hello-gcloud-func is a demo that shows how to run a progressive web app created with the [app package](https://github.com/maxence-charriere/go-app) on a [Google Cloud Function](https://cloud.google.com/functions).

## Prerequisites

- [Google Cloud SDK](https://cloud.google.com/sdk) installed on your machine
- A [Google Cloud project](https://console.cloud.google.com/cloud-resource-manager)

## Build and deploy

Go to the hello-gcloud-function directory:

```sh
cd $GOPATH/src/github.com/maxence-charriere/go-app/demo/hello-gcloud-function
```

Build the hello app:

```sh
GOARCH=wasm GOOS=js go build -o app.wasm ../hello
```

The current directory should look like the following:

```sh
# github.com/maxence-charriere/go-app/demo/hello-appengine
.
├── README.md
├── app.wasm
└── hello.go

```

Deploy the Google Cloud Function:

```sh
gcloud functions deploy Hello --runtime go113 --project YOUR_PROJECT_ID
```

See the [live demo](https://demo.murlok.io).

## Contribute

Help to develop the [app](https://github.com/maxence-charriere/go-app) package by becoming a sponsor.
<br>[Become a sponsor](https://opencollective.com/go-app).
