package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/dottmp/asd-uptime/components"
)

//go:embed static/*
var static embed.FS

func main() {
	homePage := components.Index()

	// NOTE: Mux is an implementation of go http handlders
	mux := http.NewServeMux()

	// static files
	mux.Handle("/static/", http.FileServer(http.FS(static)))

	// routes
	mux.Handle("/", templ.Handler(homePage))

	fmt.Println("Starting server on :8000")

	log.Fatalln(http.ListenAndServe(":8000", mux))
}
