package main

import "github.com/maxence-charriere/go-app/v6/pkg/app"

func main() {
	app.Route("/", &home{})
	app.Route("/select", &gameSelect{})
	app.Run()
}
