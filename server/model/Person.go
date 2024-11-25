package model

import "time"

type Person struct {
	ID               int
	Name             string
	Email            string
	Passwort         string
	Lizenz           string
	Ersteller        string
	Erstellungsdatum time.Time
	Änderer          string
	Änderungsdatum   time.Time
	Active           bool
	Token            string
}
