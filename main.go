package main

import (
	"embed"

	handlefiles "github.com/Jon-MC-dev/files_copy/handle_files"
	"github.com/gorilla/websocket"
)

//go:embed templates/*
var content embed.FS
var ScannedFiles *handlefiles.DirectoryNode

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	ScannedFiles = handlefiles.ReeadDirectory()

	return
	// Server_init1()
	Server_init2()
}
