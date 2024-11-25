package model

import "time"

type Kiste struct {
	ID               int
	Name             string
	Beschreibung     string
	Ersteller        string
	Erstellungsdatum time.Time
	Änderer          string
	Änderungsdatum   time.Time
	Verantwortlicher string
	Ort              string
}
