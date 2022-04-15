package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// (1) Create new http router
	router := httprouter.New()

	// (2) Create server
	server := http.Server{
		// (3) Set handler with router
		Handler: router,
		// (4) Set Host
		Addr: "localhost:3000",
	}

	// (5) Run server
	server.ListenAndServe()
}
