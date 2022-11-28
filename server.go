package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/Jon-MC-dev/files_copy/filepackage"
	"github.com/Masterminds/sprig/v3"
	"github.com/gorilla/mux"
)

//go:embed static/index.html*
var templateFS string

var templates *template.Template
var rootDir filepackage.DirModel
var m map[string]filepackage.DirModel

func main() {
	fmt.Println(templateFS)

	fmt.Println("Hoal mindo")
	rootDir = filepackage.ScanRootDir()
	fmt.Println(filepackage.GetMapDirs())

	// templates = template.Must(templates.ParseGlob("static/*.html"))
	templates = template.Must(
		template.New("base").Funcs(sprig.FuncMap()).Parse(templateFS),
	)

	fmt.Println("Hola mundo")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandle)
	r.HandleFunc("/{dir}", HomeHandle)
	r.HandleFunc("/file/{file}", fileRequest)

	http.ListenAndServe(":3333", r)

}

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hola mundo 2"))

	/*
		templates.ExecuteTemplate(w, "index.html", struct {
			Name string
			City string
		}{
			Name: "MyName",
			City: "MyCity",
		})
	*/
	parms := mux.Vars(r)
	if parms["dir"] == "" {
		templates.ExecuteTemplate(w, "base", rootDir)
	} else {
		if filepackage.GetMapDirs()[parms["dir"]] == nil {
			w.Write([]byte("No se encontro nada"))
		} else {
			puntero := filepackage.GetMapDirs()[parms["dir"]]
			templates.ExecuteTemplate(w, "base", puntero)
		}

	}

}

func fileRequest(w http.ResponseWriter, r *http.Request) {

	parms := mux.Vars(r)

	if parms["file"] != "" {
		if filepackage.GetMapDirs()[parms["file"]] == nil {

			w.Write([]byte("No se encontro nada"))
		} else {
			puntero := filepackage.GetMapDirs()[parms["file"]]
			fileBytes, err := ioutil.ReadFile(puntero.GetDirectory())
			if err != nil {
				panic(err)
			}
			fmt.Println("Se va a descargar el archivo: ", puntero.GetDirectory())
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(fileBytes)
			return
		}

	} else {
		w.Write([]byte("No se encontro nada"))

	}
}
