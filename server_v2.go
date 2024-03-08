package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jon-MC-dev/files_copy/functions"
	handlefiles "github.com/Jon-MC-dev/files_copy/handle_files"
	"github.com/gorilla/mux"
)

func Server_init2() {

	r := mux.NewRouter()
	// build
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(static))))
	// build fix
	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(static)))
	// run --------
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Use(loggingMiddleware)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		functions.ServeHTML(w, r, "page_nav.html", content, ScannedFiles)
	})

	r.HandleFunc("/{dir}", func(w http.ResponseWriter, r *http.Request) {
		parms := mux.Vars(r)
		if parms["dir"] == "" {
			functions.ServeHTML(w, r, "page_nav.html", content, ScannedFiles)
		} else {
			// fmt.Println(handlefiles.MapScannedFiles)
			if valor, existe := handlefiles.MapScannedFiles[parms["dir"]]; existe {
				functions.ServeHTML(w, r, "page_nav.html", content, valor)
			} else {
				functions.ServeHTML(w, r, "not_found.html", content, valor)

			}

		}
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
