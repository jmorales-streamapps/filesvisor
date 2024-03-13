package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jon-MC-dev/files_copy/bd"
	"github.com/Jon-MC-dev/files_copy/functions"
	handlefiles "github.com/Jon-MC-dev/files_copy/handle_files"
	"github.com/gorilla/mux"
	httplogger "github.com/jesseokeya/go-httplogger"
)

func Server_init2() {

	r := mux.NewRouter()
	// build
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(static))))
	// build fix
	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(static)))
	// run --------
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// r.Use(loggingMiddleware)
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
				functions.ServeHTML(w, r, "not_found.html", content, nil)

			}

		}
	})

	r.HandleFunc("/file/{file}", fileRequest)

	fmt.Println("chunk")
	r.HandleFunc("/chunk/chunk", methos).Methods("POST")
	r.HandleFunc("/preferences/chunk", configureSizeChunk).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		functions.ServeHTML(w, r, "not_found.html", content, nil)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), httplogger.Golog(r))
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("Se inicio en localhost:%s\n", PORT)

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

func fileRequest(w http.ResponseWriter, r *http.Request) {
	data := make([]int, 2000)
	for i := 0; i < 2000; i++ {
		data[i] = i + 1
	}

	//functions.ServeHTML(w, r, "info_file.html", content, data)
	//return

	parms := mux.Vars(r)

	if parms["file"] != "" {

		if file, existe := handlefiles.MapScannedFiles[parms["file"]]; existe {
			fileInfo := functions.ReadInfo(file.CompleteUrl)

			fileInfoData := struct {
				Name          string
				Size          int64
				ModTime       int64
				Reference     string
				PrefSizeChunk string
			}{
				Name:          fileInfo.Name(),
				Size:          fileInfo.Size(),
				ModTime:       fileInfo.ModTime().UnixMilli(),
				Reference:     file.Reference,
				PrefSizeChunk: "100000",
			}
			fmt.Println("el total de bytes es: ", fileInfoData.Size)

			prefSizeChunk, _ := bd.GetPreference(bdPreferences, "PrefSizeChunk")
			if prefSizeChunk != "" {
				fileInfoData.PrefSizeChunk = prefSizeChunk
			}
			fmt.Printf("El peso de las preferencias es %s", fileInfoData.PrefSizeChunk)
			fmt.Println(".")
			fmt.Println(prefSizeChunk)
			fmt.Println(".")

			functions.ServeHTML(w, r, "info_file.html", content, fileInfoData)
		} else {
			functions.ServeHTML(w, r, "not_found.html", content, nil)
			// functions.ServeHTML(w, r, "info_file.html", content, nil)

		}

	} else {
		w.Write([]byte("No se encontro nada"))

	}
}

func methos(w http.ResponseWriter, r *http.Request) {

	chunkW := &struct {
		Reference string `json:"reference"`
		From      int64  `json:"from"`
		To        int64  `json:"to"`
	}{
		From:      0,
		Reference: "",
		To:        0,
	}

	err := json.NewDecoder(r.Body).Decode(chunkW)
	if err != nil {
		w.Write([]byte(err.Error()))

		return
	}

	if file, existe := handlefiles.MapScannedFiles[chunkW.Reference]; existe {
		ServeChunks(file, struct {
			Reference string
			From      int64
			To        int64
		}{Reference: chunkW.Reference, From: chunkW.From, To: chunkW.To}, w, r)
		return
	} else {
		functions.ServeHTML(w, r, "not_found.html", content, nil)
		return
	}

}

func ServeChunks(file handlefiles.DirectoryNode, fileRequest struct {
	Reference string
	From      int64
	To        int64
}, w http.ResponseWriter, r *http.Request) {
	fileOpen, err := os.Open(file.CompleteUrl)
	if err != nil {
		functions.ServeHTML(w, r, "not_found.html", content, nil)
		return
	}
	defer fileOpen.Close()
	w.Header().Set("Content-Type", "application/octet-stream")

	length := fileRequest.To - fileRequest.From
	fmt.Println("Length: ", length)

	// Crea un slice para almacenar los bytes extraídos
	extractedData := make([]byte, length)

	_, err = fileOpen.ReadAt(extractedData, fileRequest.From)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error al leer los bytes:", err)
		return
	}

	fmt.Println(fmt.Sprintf("longitud de bytes: %d", len(extractedData)))
	w.Write(extractedData)

}

func configureSizeChunk(w http.ResponseWriter, r *http.Request) {
	fmt.Println("newChunk:::: ", "configuracion de chunk")

	chunkW := &struct {
		NewSize int32 `json:"newSize"`
	}{
		NewSize: 10_000,
	}

	err := json.NewDecoder(r.Body).Decode(chunkW)
	if err != nil {
		w.Write([]byte("0"))
		return
	}

	err = bd.SavePreference(bdPreferences, "PrefSizeChunk", fmt.Sprint(chunkW.NewSize))
	if err != nil {
		w.Write([]byte("1"))
		return
	}

	fmt.Println("newChunk:::: ", chunkW.NewSize)
	w.Write([]byte("-1"))
}
