package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// var EmbedBD embed.FS

func OpenBD() *sql.DB {
	db, err := sql.Open("sqlite3", "preferences.bd")

	if err != nil {
		fmt.Println("Error al abrir la base de datos:", err)
		return nil
	}

	fmt.Println("Se creo la bd con exito")
	return db

}

func CloseBD(db *sql.DB) {
	db.Close()
}
