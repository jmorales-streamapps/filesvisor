package main

import (
	"database/sql"
	"embed"

	"github.com/Jon-MC-dev/files_copy/bd"
	handlefiles "github.com/Jon-MC-dev/files_copy/handle_files"
	"github.com/gorilla/websocket"
)

//go:embed static/*
var static embed.FS

//go:embed templates/*
var content embed.FS

// / go:embed files_embed/*
// var embedBD embed.FS

var ScannedFiles *handlefiles.DirectoryNode

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var bdPreferences *sql.DB

func main() {
	// bd.EmbedBD = embedBD
	bdPreferences = bd.OpenBD()

	defer bd.CloseBD(bdPreferences)
	bd.CreateTable(bdPreferences)

	ScannedFiles = handlefiles.ReeadDirectory()

	//return
	// Server_init1()
	Server_init2()
}
