package model

import "time"

type Merke struct {
	ID               int
	Ersteller        string
	Erstellungsdatum time.Time
	Kiste_id         int
}
