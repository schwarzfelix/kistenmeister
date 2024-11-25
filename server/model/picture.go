package model

import "time"

type Bild struct {
	ID               int
	Bild             []byte
	Ersteller        string
	Erstellungsdatum time.Time
	Kiste_id         int
}
