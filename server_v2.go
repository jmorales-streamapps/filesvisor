package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jon-MC-dev/files_copy/functions"
	"github.com/gorilla/mux"
)

func Server_init2() {

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		functions.ServeHTML(w, r, "page_nav.html", content, ScannedFiles)
	})
	// r.HandleFunc("/{dir}", HomeHandle)
	// r.HandleFunc("/file/{file}", fileRequest)

	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
	if err != nil {
		fmt.Printf("Se inicio en localhost:%s", PORT)
	} else {
		fmt.Println(err)
	}

}

// Middleware de logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Realizar acciones de logging antes de pasar al siguiente middleware o controlador
		log.Printf("[%s] %s %s\n", r.Method, r.RequestURI, r.RemoteAddr)

		// Llamar al siguiente middleware o controlador en la cadena
		next.ServeHTTP(w, r)
	})
}
