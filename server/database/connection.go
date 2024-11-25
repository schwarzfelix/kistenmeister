package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var dbPath = "./Kistenmeister.db"

func ConnectDatabase() error {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		return err
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	} else {
		DB = db
		return nil
	}
}

func CreateTables() error {

	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS Kisten(
			ID INTEGER PRIMARY KEY NOT NULL,
			Name TEXT,
			Beschreibung TEXT,
			Ersteller TEXT,
			Erstellungsdatum DATETIME,
			Änderer TEXT,
			Änderungsdatum DATETIME,
			Verantwortlicher TEXT,
			Ort TEXT
		)
	`)

	if err != nil {
		return err
	}

	_, err1 := DB.Exec(`
		CREATE TABLE IF NOT EXISTS Kommentare(
			ID INTEGER PRIMARY KEY NOT NULL,
			Kommentar TEXT,
			Ersteller TEXT, 
			Erstellungsdatum DATETIME, 
			Kiste_id INTEGER
		)
		
	`)

	if err1 != nil {
		return err1
	}

	_, err2 := DB.Exec(`
		CREATE TABLE IF NOT EXISTS Bilder(
			ID INTEGER PRIMARY KEY NOT NULL,
			Bild BLOB,
			Ersteller TEXT, 
			Erstellungsdatum DATETIME, 
			Kiste_id INTEGER
		)
	`)

	if err2 != nil {
		return err2
	}

	_, err3 := DB.Exec(`
		CREATE TABLE IF NOT EXISTS Merklisteneinträge(
			ID INTEGER PRIMARY KEY NOT NULL,
			Ersteller TEXT, 
			Erstellungsdatum DATETIME, 
			Kiste_id INTEGER
		)
	`)

	if err3 != nil {
		return err3
	}

	_, err4 := DB.Exec(`
		CREATE TABLE IF NOT EXISTS Personen(
			ID INTEGER primary key not null,
			Name TEXT,
			Email TEXT,
			Passwort TEXT,
			Lizenz TEXT,
			Ersteller TEXT,
			Erstellungsdatum DATETIME,
			Änderer TEXT,
			Änderungsdatum DATETIME,
			Active BOOLEAN,
			Token TEXT
		)
	`)

	if err4 != nil {
		return err4
	}

	return nil
}
