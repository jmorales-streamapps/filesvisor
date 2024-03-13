package bd

import "database/sql"

var table = "preferences"

func SavePreference(db *sql.DB, key, value string) error {
	query := "INSERT OR REPLACE INTO " + table + " (key, value) VALUES (?, ?);"
	_, err := db.Exec(query, key, value)
	return err
}

func GetPreference(db *sql.DB, key string) (string, error) {
	var value string
	query := "SELECT value FROM " + table + " WHERE key = ?;"
	err := db.QueryRow(query, key).Scan(&value)
	return value, err
}

func DeletePreference(db *sql.DB, key string) error {
	query := "DELETE FROM " + table + " WHERE key = ?;"
	_, err := db.Exec(query, key)
	return err
}

func CreateTable(db *sql.DB) error {
	query := `
        CREATE TABLE IF NOT EXISTS ` + table + `(
            key TEXT PRIMARY KEY,
            value TEXT
        );
    `
	_, err := db.Exec(query)
	return err
}
