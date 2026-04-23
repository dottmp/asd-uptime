package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/dottmp/asd-uptime/components"
)

var static embed.FS

func main() {
	homePage := components.Index()

	// pages
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	log.Fatalln(http.ListenAndServe(":8000", pagesHandler))
}
