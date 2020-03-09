# hello-gcloud-func

hello-gcloud-func is a demo that shows how to run a progressive web app created with the [app package](https://github.com/maxence-charriere/go-app) on a [Google Cloud Function](https://cloud.google.com/functions).

## Prerequisites

- [Google Cloud SDK](https://cloud.google.com/sdk) installed on your machine
- A [Google Cloud project](https://console.cloud.google.com/cloud-resource-manager)
- [Firebase CLI](https://firebase.google.com/docs/cli) installed on your machine

## Build and deploy

Go to the hello-gcloud-function directory:

```sh
cd $GOPATH/src/github.com/maxence-charriere/go-app/demo/hello-gcloud-function
```

Copy the hello wasm binary:

```sh
cp ../hello/app.wasm .
```

The current directory should look like the following:

```sh
# github.com/maxence-charriere/go-app/demo/hello-appengine
.
├── README.md
├── app.wasm
├── go.mod
├── go.sum
└── hello.go

```

Deploy the Google Cloud Function:

```sh
gcloud functions deploy Hello --runtime go113 --trigger-http --project YOUR_PROJECT_ID
```

Answer `y` if it asks for:

```sh
API [cloudfunctions.googleapis.com] not enabled on project
[xxx]. Would you like to enable and retry (this will take a
few minutes)? (y/N)?
```

Answer `y` if it asks for:

```sh
Allow unauthenticated invocations of new function [Hello]? (y/N)?
```

You may have this warning:

```sh
WARNING: Function created with limited-access IAM policy. To enable unauthorized access consider "gcloud alpha functions add-iam-policy-binding Hello --member=allUsers --role=roles/cloudfunctions.invoker"
```

If it is the case, delete the function and redeploy it.

## Firebase

Trying to reach the function as it is will result in CORS errors.
You need to use [Firebase](https://firebase.google.com) in order to provide the function under a single host.

### Init Firebase project

```sh
firebase init
```

When ask about initializing a Firebase project, chose `Hosting` one:

```
You're about to initialize a Firebase project in this directory:

  /Users/maxence/dev/src/github.com/maxence-charriere/go-app-demo/hello-gcloud-func

? Which Firebase CLI features do you want to set up for this folder? Press Space to select feature
s, then Enter to confirm your choices.
 ◯ Database: Deploy Firebase Realtime Database Rules
 ◯ Firestore: Deploy rules and create indexes for Firestore
 ◯ Functions: Configure and deploy Cloud Functions
❯◉ Hosting: Configure and deploy Firebase Hosting sites
 ◯ Storage: Deploy Cloud Storage security rules
 ◯ Emulators: Set up local emulators for Firebase features
```

When asked about the project setup, chose the `Add Firebase to an existing Google Cloud Platform project` option and select your Google Cloud project:

```
=== Project Setup

First, let's associate this project directory with a Firebase project.
You can create multiple project aliases by running firebase use --add,
but for now we'll just set up a default project.

? Please select an option:
  Use an existing project
  Create a new project
❯ Add Firebase to an existing Google Cloud Platform project
  Don't set up a default project
```

Answer the next steps with the default options.

### Reroute traffic to the Google Cloud function

- Rename `public/index.html` to `public/index2.html`. It is a hack that will allow the cloud function to be loaded from the root of the firebase URL.
- Open `firebase.json` and modify the rewrite content to point to the `Hello` cloud function:

  ```json
      "rewrites": [
      {
        "source": "**",
        "function": "Hello"
      }
    ]
  ```

- Deploy the firebase modifications:

  ```sh
  firebase deploy
  ```

The app should be now be accessible at your Firebase project URL available in the [Firebase consoler](https://console.firebase.google.com) in your project hosting section.

See the [live demo](https://go-app-demo-42.firebaseapp.com/).

## Contribute

Help to develop the [go-app](https://github.com/maxence-charriere/go-app) package by becoming a sponsor.
<br>[Become a sponsor](https://opencollective.com/go-app).
