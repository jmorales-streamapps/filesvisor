package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/Jon-MC-dev/files_copy/functions"
	handlefiles "github.com/Jon-MC-dev/files_copy/handle_files"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

//go:embed templates/*
var content embed.FS

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Server_init2() {

	handlefiles.ReeadDirectory()

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		functions.ServeHTML(w, r, "test.html", content)
	})
	// r.HandleFunc("/{dir}", HomeHandle)
	// r.HandleFunc("/file/{file}", fileRequest)

	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
	if err != nil {
		fmt.Printf("Se inicio en localhost:%s", PORT)
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
