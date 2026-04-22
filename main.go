package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/dottmp/asd-uptime/components"
)

func main() {
	greet := components.Greet("Lizzy The Cat", 2)
	handler := templ.Handler(greet)
	log.Fatalln(http.ListenAndServe(":8080", handler))
}
