package model

import "time"

type Kommentar struct {
	ID               int
	Kommentar        string
	Ersteller        string
	Erstellungsdatum time.Time
	Kiste_id         int
}
