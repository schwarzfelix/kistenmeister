package database

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schwarzfelix/kistenmeister/server/model"
)

var c *gin.Context

// Funktion zum Prüfen, ob angegebenen Token in der Datenbank existiert
func GetToken(token string) error {
	tokenRow := DB.QueryRow("SELECT Count(*) FROM Personen WHERE Token = ?;", token)
	var tokenCount int
	err := tokenRow.Scan(&tokenCount)
	if err != nil {
		return err
	} else if tokenCount < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiger Token"})
		return err
	} else {
		return nil
	}
}

// Funktion zum Prüfen, ob Person mit dem angegebenen Token ein Pro-Anwender ist
func GetLizenz(token string) error {
	lizenzRow := DB.QueryRow("SELECT Lizenz FROM Personen WHERE Token = ?;", token)
	var lizenz string
	err := lizenzRow.Scan(&lizenz)
	if err != nil {
		return err
	} else if lizenz != "Pro" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Sie benötigen für die Ausführung dieser Funktionalität die Pro-Lizenz"})
		return err
	} else {
		return nil
	}
}

// Funktion zum Abfragen aller Kisten
func GetBoxes(token string) ([]model.Kiste, error) {
	if GetToken(token) != nil {
		return nil, GetToken(token)
	}
	rows, err := DB.Query("SELECT ID, Name, Beschreibung, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Verantwortlicher, Ort FROM Kisten;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleKisten := make([]model.Kiste, 0)

	for rows.Next() {
		einzelneKiste := model.Kiste{}
		err = rows.Scan(&einzelneKiste.ID, &einzelneKiste.Name, &einzelneKiste.Beschreibung, &einzelneKiste.Ersteller, &einzelneKiste.Erstellungsdatum, &einzelneKiste.Änderer, &einzelneKiste.Änderungsdatum, &einzelneKiste.Verantwortlicher, &einzelneKiste.Ort)

		if err != nil {
			return nil, err
		}

		alleKisten = append(alleKisten, einzelneKiste)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleKisten, err
}

// Funktion zum Abfragen aller Personen in der DB
func GetPersonen(token string) ([]model.Person, error) {
	if GetToken(token) != nil {
		return nil, GetToken(token)
	}
	rows, err := DB.Query("SELECT ID, Name, Email, Lizenz, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Active, Token FROM Personen;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allePersonen := make([]model.Person, 0)

	for rows.Next() {
		einzelnePerson := model.Person{}
		err = rows.Scan(&einzelnePerson.ID, &einzelnePerson.Name, &einzelnePerson.Email, &einzelnePerson.Lizenz, &einzelnePerson.Ersteller, &einzelnePerson.Erstellungsdatum, &einzelnePerson.Änderer, &einzelnePerson.Änderungsdatum, &einzelnePerson.Active, &einzelnePerson.Token)

		if err != nil {
			return nil, err
		}

		allePersonen = append(allePersonen, einzelnePerson)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return allePersonen, err
}

// Funktion zum Abfragen aller kommentare
func GetComments(token string) ([]model.Kommentar, error) {
	if GetToken(token) != nil {
		return nil, GetToken(token)
	}
	rows, err := DB.Query("SELECT ID, Kommentar, Ersteller, Erstellungsdatum, Kiste_id FROM Kommentare;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleKommentare := make([]model.Kommentar, 0)

	for rows.Next() {
		einzelneKommentar := model.Kommentar{}
		err = rows.Scan(&einzelneKommentar.ID, &einzelneKommentar.Kommentar, &einzelneKommentar.Ersteller, &einzelneKommentar.Erstellungsdatum, &einzelneKommentar.Kiste_id)

		if err != nil {
			return nil, err
		}

		alleKommentare = append(alleKommentare, einzelneKommentar)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleKommentare, err
}

// Funktion zum Abfragen aller kommentare
func GetStars(token string) ([]model.Merke, error) {
	if GetToken(token) != nil {
		return nil, GetToken(token)
	}
	rows, err := DB.Query("SELECT ID, Ersteller, Erstellungsdatum, Kiste_id FROM Merklisteneinträge;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleMerken := make([]model.Merke, 0)

	for rows.Next() {
		einzelneMerke := model.Merke{}
		err = rows.Scan(&einzelneMerke.ID, &einzelneMerke.Ersteller, &einzelneMerke.Erstellungsdatum, &einzelneMerke.Kiste_id)

		if err != nil {
			return nil, err
		}

		alleMerken = append(alleMerken, einzelneMerke)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleMerken, err
}

// Funktion zum Abfragen aller Bilder
func GetPictures(token string) ([]model.Bild, error) {
	if GetToken(token) != nil {
		return nil, GetToken(token)
	}
	rows, err := DB.Query("SELECT ID, Bild, Ersteller, Erstellungsdatum, Kiste_id FROM Bilder;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleBilder := make([]model.Bild, 0)

	for rows.Next() {
		einzelnesBild := model.Bild{}
		err = rows.Scan(&einzelnesBild.ID, &einzelnesBild.Bild, &einzelnesBild.Ersteller, &einzelnesBild.Erstellungsdatum, &einzelnesBild.Kiste_id)

		if err != nil {
			return nil, err
		}

		alleBilder = append(alleBilder, einzelnesBild)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleBilder, err
}

// Funktion zum Abfragen einer einzelnen Kiste je nach id
func GetBox(id int, token string) (model.Kiste, error) {
	if GetToken(token) != nil {
		return model.Kiste{}, GetToken(token)
	}
	row := DB.QueryRow("SELECT ID, Name, Beschreibung, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Verantwortlicher, Ort FROM Kisten WHERE ID = ?;", id)
	singleBox := model.Kiste{}
	err := row.Scan(&singleBox.ID, &singleBox.Name, &singleBox.Beschreibung, &singleBox.Ersteller, &singleBox.Erstellungsdatum, &singleBox.Änderer, &singleBox.Änderungsdatum, &singleBox.Verantwortlicher, &singleBox.Ort)
	if err != nil {
		return model.Kiste{}, err
	}
	return singleBox, err
}

// Funktion zum Abfragen einer einzelnen Person je nach ID
func GetPerson(token string) (model.Person, error) {
	if GetToken(token) != nil {
		return model.Person{}, GetToken(token)
	}
	row := DB.QueryRow("SELECT ID, Name, Email, Lizenz, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Active, Token FROM Personen WHERE Token = ?;", token)
	singlePerson := model.Person{}
	err := row.Scan(&singlePerson.ID, &singlePerson.Name, &singlePerson.Email, &singlePerson.Lizenz, &singlePerson.Ersteller, &singlePerson.Erstellungsdatum, &singlePerson.Änderer, &singlePerson.Änderungsdatum, &singlePerson.Active, &singlePerson.Token)
	if err != nil {
		return model.Person{}, err
	}
	return singlePerson, err
}

// Funktion zum Abfragen aller Kommentare einer Kiste je nach kiste_id
func GetComment(id int, token string) ([]model.Kommentar, error) {
	if GetToken(token) != nil {
		return nil, GetToken(token)
	}
	rows, err := DB.Query("SELECT ID, Kommentar, Ersteller, Erstellungsdatum, Kiste_id FROM Kommentare WHERE Kiste_id = ?;", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleKommentarezuKiste := make([]model.Kommentar, 0)

	for rows.Next() {
		einzelneKommentar := model.Kommentar{}
		err = rows.Scan(&einzelneKommentar.ID, &einzelneKommentar.Kommentar, &einzelneKommentar.Ersteller, &einzelneKommentar.Erstellungsdatum, &einzelneKommentar.Kiste_id)

		if err != nil {
			return nil, err
		}

		alleKommentarezuKiste = append(alleKommentarezuKiste, einzelneKommentar)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleKommentarezuKiste, err
}

// Funktion zum Abfragen aller Merken einer Person
func GetStar(token string) ([]model.Merke, error) {
	if GetToken(token) != nil {
		return nil, GetToken(token)
	}
	personRow := DB.QueryRow("SELECT Name FROM Personen WHERE Token = ?;", token)
	var person string
	err := personRow.Scan(&person)
	if err != nil {
		return nil, err
	}
	rows, err := DB.Query("SELECT ID, Ersteller, Erstellungsdatum, Kiste_id FROM Merklisteneinträge WHERE Ersteller = ?;", person)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleMerkenzuPerson := make([]model.Merke, 0)

	for rows.Next() {
		einzelneMerke := model.Merke{}
		err = rows.Scan(&einzelneMerke.ID, &einzelneMerke.Ersteller, &einzelneMerke.Erstellungsdatum, &einzelneMerke.Kiste_id)

		if err != nil {
			return nil, err
		}

		alleMerkenzuPerson = append(alleMerkenzuPerson, einzelneMerke)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleMerkenzuPerson, err
}

// Funktion zum Abfragen aller Bilder einer Kiste je nach kiste_id
func GetPicture(id int, token string) ([]model.Bild, error) {
	if GetToken(token) != nil {
		return nil, GetToken(token)
	}
	rows, err := DB.Query("SELECT ID, Bild, Ersteller, Erstellungsdatum, Kiste_id FROM bilder WHERE Kiste_id = ?;", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	alleBilderzuKiste := make([]model.Bild, 0)

	for rows.Next() {
		einzelnesBild := model.Bild{}
		err = rows.Scan(&einzelnesBild.ID, &einzelnesBild.Bild, &einzelnesBild.Ersteller, &einzelnesBild.Erstellungsdatum, &einzelnesBild.Kiste_id)

		if err != nil {
			return nil, err
		}

		alleBilderzuKiste = append(alleBilderzuKiste, einzelnesBild)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return alleBilderzuKiste, err
}

// Funktion zum Löschen einer einzelnen Kiste je nach id
func DeleteBox(id int, token string) (model.Kiste, error) {
	if GetLizenz(token) != nil {
		return model.Kiste{}, GetLizenz(token)
	}
	if GetToken(token) != nil {
		return model.Kiste{}, GetToken(token)
	}
	row := DB.QueryRow("SELECT ID, Name, Beschreibung, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Verantwortlicher, Ort FROM Kisten WHERE id = ?;", id)

	deletedBox := model.Kiste{}
	err1 := row.Scan(&deletedBox.ID, &deletedBox.Name, &deletedBox.Beschreibung, &deletedBox.Ersteller, &deletedBox.Erstellungsdatum, &deletedBox.Änderer, &deletedBox.Änderungsdatum, &deletedBox.Verantwortlicher, &deletedBox.Ort)

	if err1 != nil {
		return model.Kiste{}, err1
	}

	deletebox, err := DB.Begin()

	if err != nil {
		return deletedBox, err
	}

	// Beim Löschen einer Kiste müssen die dazugehörigen Bilder und Kommentare auch gelöscht werden.
	delete_stmt, err := DB.Prepare("DELETE FROM Kisten WHERE ID = ?;")
	delete_stmt1, err1 := DB.Prepare("DELETE FROM Kommentare WHERE Kiste_id = ?;")
	delete_stmt2, err2 := DB.Prepare("DELETE FROM Bilder WHERE Kiste_id = ?;")
	delete_stmt3, err3 := DB.Prepare("DELETE FROM Merklisteneinträge WHERE Kiste_id = ?;")

	if err != nil {
		return deletedBox, err
	}

	defer delete_stmt.Close()

	if err1 != nil {
		return deletedBox, err1
	}

	defer delete_stmt1.Close()

	if err2 != nil {
		return deletedBox, err2
	}

	defer delete_stmt2.Close()

	if err3 != nil {
		return deletedBox, err3
	}

	defer delete_stmt3.Close()

	_, err = delete_stmt.Exec(id)

	if err != nil {
		return deletedBox, err
	}

	_, err = delete_stmt1.Exec(id)

	if err != nil {
		return deletedBox, err
	}

	_, err = delete_stmt2.Exec(id)

	if err != nil {
		return deletedBox, err
	}

	_, err = delete_stmt3.Exec(id)

	if err != nil {
		return deletedBox, err
	}

	deletebox.Commit()
	return deletedBox, err1

}

// Funktion zum Löschen eines einzelnen Kommentars je nach id
func DeleteComment(id int, token string) (model.Kommentar, error) {
	if GetLizenz(token) != nil {
		return model.Kommentar{}, GetLizenz(token)
	}
	if GetToken(token) != nil {
		return model.Kommentar{}, GetToken(token)
	}
	row := DB.QueryRow("SELECT ID, Kommentar, Ersteller, Erstellungsdatum, Kiste_id FROM Kommentare WHERE ID = ?;", id)

	deletedComment := model.Kommentar{}
	err1 := row.Scan(&deletedComment.ID, &deletedComment.Kommentar, &deletedComment.Ersteller, &deletedComment.Erstellungsdatum, &deletedComment.Kiste_id)

	if err1 != nil {
		return model.Kommentar{}, err1
	}

	deletecomment, err := DB.Begin()

	if err != nil {
		return deletedComment, err
	}

	delete_stmt, err := DB.Prepare("DELETE FROM Kommentare WHERE ID = ?;")

	if err != nil {
		return deletedComment, err
	}

	defer delete_stmt.Close()

	_, err = delete_stmt.Exec(id)

	if err != nil {
		return deletedComment, err
	}

	deletecomment.Commit()
	return deletedComment, err1
}

// Funktion zum Löschen einer einzelnen Merke je nach id
func DeleteStar(id int, token string) (model.Merke, error) {
	if GetToken(token) != nil {
		return model.Merke{}, GetToken(token)
	}
	row := DB.QueryRow("SELECT ID, Ersteller, Erstellungsdatum, Kiste_id FROM Merklisteneinträge WHERE ID = ?;", id)

	deletedStar := model.Merke{}
	err1 := row.Scan(&deletedStar.ID, &deletedStar.Ersteller, &deletedStar.Erstellungsdatum, &deletedStar.Kiste_id)

	if err1 != nil {
		return model.Merke{}, err1
	}

	deletestar, err := DB.Begin()

	if err != nil {
		return deletedStar, err
	}

	delete_stmt, err := DB.Prepare("DELETE FROM Merklisteneinträge WHERE ID = ?;")

	if err != nil {
		return deletedStar, err
	}

	defer delete_stmt.Close()

	_, err = delete_stmt.Exec(id)

	if err != nil {
		return deletedStar, err
	}

	deletestar.Commit()
	return deletedStar, err1
}

// Funktion zum Löschen eines einzelnen Bilds je nach id
func DeletePicture(id int, token string) (model.Bild, error) {
	if GetLizenz(token) != nil {
		return model.Bild{}, GetLizenz(token)
	}
	if GetToken(token) != nil {
		return model.Bild{}, GetToken(token)
	}
	row := DB.QueryRow("SELECT ID, Bild, Ersteller, Erstellungsdatum, Kiste_id FROM Bilder WHERE ID = ?;", id)

	deletedPicture := model.Bild{}
	err1 := row.Scan(&deletedPicture.ID, &deletedPicture.Bild, &deletedPicture.Ersteller, &deletedPicture.Erstellungsdatum, &deletedPicture.Kiste_id)

	if err1 != nil {
		return model.Bild{}, err1
	}

	deletepicture, err := DB.Begin()

	if err != nil {
		return deletedPicture, err
	}

	delete_stmt, err := DB.Prepare("DELETE FROM Bilder WHERE id = ?;")

	if err != nil {
		return deletedPicture, err
	}

	defer delete_stmt.Close()

	_, err = delete_stmt.Exec(id)

	if err != nil {
		return deletedPicture, err
	}

	deletepicture.Commit()
	return deletedPicture, err1
}

// Funktion zum Löschen einer einzelnen Person je nach id
func DeletePerson(id int, token string) (model.Person, error) {
	if GetToken(token) != nil {
		return model.Person{}, GetToken(token)
	}
	row := DB.QueryRow("SELECT ID, Name, Email, Passwort, Lizenz, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Active, Token FROM Personen WHERE ID = ?;", id)

	deletedPerson := model.Person{}
	err1 := row.Scan(&deletedPerson.ID, &deletedPerson.Name, &deletedPerson.Email, &deletedPerson.Passwort, &deletedPerson.Lizenz, &deletedPerson.Ersteller, &deletedPerson.Erstellungsdatum, &deletedPerson.Änderer, &deletedPerson.Änderungsdatum, &deletedPerson.Active, &deletedPerson.Token)

	if err1 != nil {
		return model.Person{}, err1
	}

	deleteperson, err := DB.Begin()

	if err != nil {
		return deletedPerson, err
	}

	delete_stmt, err := DB.Prepare("DELETE FROM Personen WHERE ID = ?;")

	if err != nil {
		return deletedPerson, err
	}

	defer delete_stmt.Close()

	_, err = delete_stmt.Exec(id)

	if err != nil {
		return deletedPerson, err
	}

	deleteperson.Commit()
	return deletedPerson, err1
}

// Funktion zum Aktualiseren einer Kiste
func UpdateBox(id int, updatedBoxValues model.Kiste, token string) (model.Kiste, error) {
	if GetLizenz(token) != nil {
		return model.Kiste{}, GetLizenz(token)
	}
	if GetToken(token) != nil {
		return model.Kiste{}, GetToken(token)
	}
	creatorRow := DB.QueryRow("Select Name from Personen where Token = ?", token)
	var creator string
	err := creatorRow.Scan(&creator)
	if err != nil {
		return model.Kiste{}, err
	}
	row := DB.QueryRow("SELECT ID, Name, Beschreibung, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Verantwortlicher, Ort FROM Kisten WHERE ID = ?;", id)

	currentBox := model.Kiste{}
	err1 := row.Scan(&currentBox.ID, &currentBox.Name, &currentBox.Beschreibung, &currentBox.Ersteller, &currentBox.Erstellungsdatum, &currentBox.Änderer, &currentBox.Änderungsdatum, &currentBox.Verantwortlicher, &currentBox.Ort)

	if err1 != nil {
		return model.Kiste{}, err1
	}

	updatebox, err := DB.Begin()
	if err != nil {
		return currentBox, err
	}

	update_stmt, err := updatebox.Prepare("UPDATE Kisten SET Name = ?, Beschreibung = ?, Ersteller = ?, Erstellungsdatum = ?, Änderer = ?, Änderungsdatum = ?, Verantwortlicher = ?, Ort = ? WHERE ID = ?;")

	if err != nil {
		return currentBox, err
	}

	defer update_stmt.Close()

	_, err = update_stmt.Exec(updatedBoxValues.Name, updatedBoxValues.Beschreibung, &currentBox.Ersteller, &currentBox.Erstellungsdatum, creator, updatedBoxValues.Änderungsdatum, updatedBoxValues.Verantwortlicher, updatedBoxValues.Ort, id)

	if err != nil {
		return currentBox, err
	}

	updatebox.Commit()

	row_updated := DB.QueryRow("SELECT ID, Name, Beschreibung, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Verantwortlicher, Ort FROM Kisten WHERE ID = ?;", id)

	updatedBox := model.Kiste{}
	err2 := row_updated.Scan(&updatedBox.ID, &updatedBox.Name, &updatedBox.Beschreibung, &currentBox.Ersteller, &currentBox.Erstellungsdatum, &updatedBox.Änderer, &updatedBox.Änderungsdatum, &updatedBox.Verantwortlicher, &updatedBox.Ort)

	if err2 != nil {
		return model.Kiste{}, err2
	} else {
		return updatedBox, nil
	}
}

// Funktion zum Aktualiseren der Daten einer Person
func UpdatePerson(updatedPersonValues model.Person, token string) (model.Person, error) {
	if GetLizenz(token) != nil {
		return model.Person{}, GetLizenz(token)
	}
	creatorRow := DB.QueryRow("Select Name from Personen where Token = ?", token)
	var creator string
	err := creatorRow.Scan(&creator)
	if err != nil {
		return model.Person{}, err
	}

	row := DB.QueryRow("SELECT ID, Name, Email, Passwort, Lizenz, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Active, Token FROM Personen WHERE Token = ?;", token)

	currentPerson := model.Person{}
	err1 := row.Scan(&currentPerson.ID, &currentPerson.Name, &currentPerson.Email, &currentPerson.Passwort, &currentPerson.Lizenz, &currentPerson.Ersteller, &currentPerson.Erstellungsdatum, &currentPerson.Änderer, &currentPerson.Änderungsdatum, &currentPerson.Active, &currentPerson.Token)

	if err1 != nil {
		return model.Person{}, err1
	}

	updateperson, err := DB.Begin()
	if err != nil {
		return currentPerson, err
	}

	update_stmt, err := updateperson.Prepare("UPDATE Personen SET Name = ?, Email = ?, Passwort = ?, Lizenz = ?, Ersteller = ?, Erstellungsdatum = ?, Änderer = ?, Änderungsdatum = ?, Active = ?, Token=?  WHERE Token = ?;")

	if err != nil {
		return currentPerson, err
	}

	defer update_stmt.Close()

	_, err = update_stmt.Exec(updatedPersonValues.Name, updatedPersonValues.Email, &updatedPersonValues.Passwort, &updatedPersonValues.Lizenz, &currentPerson.Ersteller, &currentPerson.Erstellungsdatum, creator, &updatedPersonValues.Änderungsdatum, &currentPerson.Active, &currentPerson.Token, token)

	if err != nil {
		return currentPerson, err
	}

	updateperson.Commit()

	row_updated := DB.QueryRow("SELECT ID, Name, Email, Passwort, Lizenz, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Active, Token FROM Personen WHERE Token = ?;", token)

	updatedPerson := model.Person{}
	err2 := row_updated.Scan(&updatedPerson.ID, &updatedPerson.Name, &updatedPerson.Email, &updatedPerson.Passwort, &updatedPerson.Lizenz, &updatedPerson.Ersteller, &updatedPerson.Erstellungsdatum, &updatedPerson.Änderer, &updatedPerson.Änderungsdatum, &updatedPerson.Active, &updatedPerson.Token)

	if err2 != nil {
		return model.Person{}, err2
	} else {
		return updatedPerson, nil
	}
}

// Funktion zum Erstellen einer neuen Kiste -- Gibt neue id automatisch zurück
func CreateBox(newBox model.Kiste, token string) (int, error) {
	if GetLizenz(token) != nil {
		return -1, GetLizenz(token)
	}
	if GetToken(token) != nil {
		return -1, GetToken(token)
	}
	creatorRow := DB.QueryRow("Select Name from Personen where Token = ?", token)
	var creator string
	err := creatorRow.Scan(&creator)
	if err != nil {
		return -1, err
	}
	// Query to check if the table has any rows
	row := DB.QueryRow("SELECT COUNT(*) FROM Kisten")
	var count int
	err = row.Scan(&count)
	if err != nil {
		return -1, err
	}
	var max_id int
	if count == 0 {
		// No rows in the table, set maxValue to 0
		max_id = 0
	} else {
		// Query max(Id) from the data table
		query := "SELECT MAX(ID) FROM Kisten"
		// Execute the query and scan the result into `max_id`
		err := DB.QueryRow(query).Scan(&max_id)
		if err != nil {
			return -1, err
		}
	}

	var new_id int = max_id + 1 // store new_id as max(Id) + 1
	// prepare with 9 placeholders
	create_stmt, err := DB.Prepare("INSERT INTO Kisten (ID, Name, Beschreibung, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Verantwortlicher, Ort) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return -1, err
	}
	defer create_stmt.Close()
	_, err1 := create_stmt.Exec(new_id, newBox.Name, newBox.Beschreibung, creator, newBox.Erstellungsdatum, creator, newBox.Änderungsdatum, creator, newBox.Ort)

	if err1 != nil {
		return -1, err1
	} else {
		return new_id, nil
	}
}

// Funktion zum Erstellen eines neuen Kommentars -- Gibt neue id automatisch zurück
func CreateComment(newComment model.Kommentar, token string) (int, error) {
	if GetToken(token) != nil {
		return -1, GetToken(token)
	}
	erstellerRow := DB.QueryRow("SELECT Name FROM Personen WHERE Token = ?", token)
	var ersteller string
	err := erstellerRow.Scan(&ersteller)

	if err != nil {
		return -1, err
	}
	// Prüfen, ob eine Kiste mit der angegebenen Kiste_id in der DB vorhanden ist
	check := DB.QueryRow("SELECT COUNT(*) FROM Kisten WHERE ID = ?;", newComment.Kiste_id)
	var check_count int
	err = check.Scan(&check_count)
	if err != nil {
		return -1, err
	}
	if check_count == 0 {
		return -1, nil
	} else {
		// Query to check if the table has any rows
		row := DB.QueryRow("SELECT COUNT(*) FROM Kommentare;")
		var count int
		err := row.Scan(&count)
		if err != nil {
			return -1, err
		}
		var max_id int
		if count == 0 {
			// No rows in the table, set maxValue to 0
			max_id = 0
		} else {
			// Query max(Id) from the data table
			query := "SELECT MAX(ID) FROM Kommentare;"
			// Execute the query and scan the result into `max_id`
			err = DB.QueryRow(query).Scan(&max_id)
			if err != nil {
				return -1, err
			}
		}

		var new_id = max_id + 1 // store new_id as max(Id) + 1
		// prepare with 5 placeholders
		create_stmt, err := DB.Prepare("INSERT INTO Kommentare (ID, Kommentar, Ersteller, Erstellungsdatum, Kiste_id) VALUES (?, ?, ?, ?, ?);")
		if err != nil {
			return -1, err
		}

		defer create_stmt.Close()

		_, err1 := create_stmt.Exec(new_id, newComment.Kommentar, ersteller, newComment.Erstellungsdatum, newComment.Kiste_id)

		if err1 != nil {
			return -1, err1
		} else {
			return new_id, nil
		}
	}
}

// Funktion zum Merken einer Kiste zu gegebener Person
func CreateStar(newStar model.Merke, token string) (int, error) {
	if GetToken(token) != nil {
		return -1, GetToken(token)
	}
	erstellerRow := DB.QueryRow("SELECT Name FROM Personen WHERE Token = ?", token)
	var ersteller string
	err := erstellerRow.Scan(&ersteller)

	if err != nil {
		return -1, err
	}
	// Prüfen, ob eine Kiste mit der angegebenen Kiste_id in der DB vorhanden ist
	check := DB.QueryRow("SELECT COUNT(*) FROM Kisten WHERE ID = ?;", newStar.Kiste_id)
	var check_count int
	err = check.Scan(&check_count)
	if err != nil {
		return -1, err
	}
	if check_count == 0 {
		return -2, nil
	} else {
		// Query to check if the table has any rows
		row := DB.QueryRow("SELECT COUNT(*) FROM Merklisteneinträge")
		var count int
		err := row.Scan(&count)
		if err != nil {
			return -1, err
		}
		var max_id int
		if count == 0 {
			// No rows in the table, set maxValue to 0
			max_id = 0
		} else {
			// Query max(Id) from the data table
			query := "SELECT MAX(ID) FROM Merklisteneinträge"
			// Execute the query and scan the result into `max_id`
			err = DB.QueryRow(query).Scan(&max_id)
			if err != nil {
				return -1, err
			}
		}

		var new_id = max_id + 1 // store new_id as max(Id) + 1
		// prepare with 5 placeholders
		create_stmt, err := DB.Prepare("INSERT INTO Merklisteneinträge (ID, Ersteller, Erstellungsdatum, Kiste_id) VALUES (?, ?, ?, ?);")
		if err != nil {
			return -1, err
		}
		defer create_stmt.Close()

		_, err1 := create_stmt.Exec(new_id, ersteller, newStar.Erstellungsdatum, newStar.Kiste_id)

		if err1 != nil {
			return -1, err1
		} else {
			return new_id, nil
		}
	}
}

// Funktion zum Erstellen eines neuen Bilds -- Gibt neue id automatisch zurück
func CreatePicture(newPicture model.Bild, token string) (int, error) {
	if GetToken(token) != nil {
		return -1, GetToken(token)
	}
	erstellerRow := DB.QueryRow("SELECT Name FROM Personen WHERE Token = ?", token)
	var ersteller string
	err := erstellerRow.Scan(&ersteller)

	if err != nil {
		return -1, err
	}
	// Prüfen, ob eine Kiste mit der angegebenen Kiste_id in der DB vorhanden ist
	check := DB.QueryRow("SELECT COUNT(*) FROM Kisten WHERE ID = ?;", newPicture.Kiste_id)
	var check_count int
	err = check.Scan(&check_count)
	if err != nil {
		return -1, err
	}
	if check_count == 0 {
		return -1, nil
	} else {
		// Query to check if the table has any rows
		row := DB.QueryRow("SELECT COUNT(*) FROM Bilder")
		var count int
		err := row.Scan(&count)
		if err != nil {
			return -1, err
		}
		var max_id int
		if count == 0 {
			// No rows in the table, set maxValue to 0
			max_id = 0
		} else {
			// Query max(Id) from the data table
			query := "SELECT MAX(ID) FROM Bilder"
			// Execute the query and scan the result into `max_id`
			err = DB.QueryRow(query).Scan(&max_id)
			if err != nil {
				return -1, err
			}
		}

		var new_id int = max_id + 1 // store new_id as max(Id) + 1
		// prepare with 4 placeholders
		create_stmt, err := DB.Prepare("INSERT INTO Bilder (ID, Bild, Ersteller, Erstellungsdatum, Kiste_id) VALUES (?, ?, ?, ?, ?);")
		if err != nil {
			return -1, err
		}
		defer create_stmt.Close()
		_, err1 := create_stmt.Exec(new_id, newPicture.Bild, ersteller, newPicture.Erstellungsdatum, newPicture.Kiste_id)

		if err1 != nil {
			return -1, err1
		} else {
			return new_id, nil
		}
	}
}

// Funktion zum Hinzufügen einer neuen Person zur Datenbank
func CreatePerson(newPerson model.Person, token string) (int, error) {
	if GetLizenz(token) != nil {
		return -1, GetLizenz(token)
	}

	erstellerRow := DB.QueryRow("SELECT Name FROM Personen WHERE Token = ?", token)
	var ersteller string
	err := erstellerRow.Scan(&ersteller)

	if err != nil {
		return -1, err
	}
	// Query to check if the table has any rows
	row := DB.QueryRow("SELECT COUNT(*) FROM Personen")
	var count int
	err = row.Scan(&count)

	if err != nil {
		return -1, err
	}

	var max_id int
	if count == 0 {
		// No rows in the table, set maxValue to 0
		max_id = 0
	} else {
		// Query max(Id) from the data table
		query := "SELECT MAX(ID) FROM Personen"
		// Execute the query and scan the result into `max_id`
		err = DB.QueryRow(query).Scan(&max_id)
		if err != nil {
			return -1, err
		}
	}

	var new_id = max_id + 1 // store new_id as max(Id) + 1
	// prepare with 9 placeholders
	create_stmt, err := DB.Prepare("INSERT INTO Personen (ID, Name, Email, Passwort, Lizenz, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Active, Token) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return -1, err
	}
	defer create_stmt.Close()
	_, err1 := create_stmt.Exec(new_id, newPerson.Name, newPerson.Email, newPerson.Passwort, newPerson.Lizenz, ersteller, newPerson.Erstellungsdatum, ersteller, newPerson.Änderungsdatum, newPerson.Active, newPerson.Token)

	if err1 != nil {
		return -1, err1
	} else {
		return new_id, nil
	}
}

// Funktion zum Aktualiseren der Daten einer Person
func UpdateToken(id int, token string) error {

	updateperson, err := DB.Begin()
	if err != nil {
		return err
	}

	update_stmt, err := updateperson.Prepare("UPDATE Personen SET Token = ? WHERE ID = ?;")

	if err != nil {
		return err
	}

	defer update_stmt.Close()

	_, err = update_stmt.Exec(token, id)

	if err != nil {
		return err
	}

	updateperson.Commit()

	row_updated := DB.QueryRow("SELECT ID, Name, Email, Passwort, Lizenz, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Active, Token FROM Personen WHERE ID = ?;", id)

	updatedPerson := model.Person{}
	err2 := row_updated.Scan(&updatedPerson.ID, &updatedPerson.Name, &updatedPerson.Email, &updatedPerson.Passwort, &updatedPerson.Lizenz, &updatedPerson.Ersteller, &updatedPerson.Erstellungsdatum, &updatedPerson.Änderer, &updatedPerson.Änderungsdatum, &updatedPerson.Active, &updatedPerson.Token)

	if err2 != nil {
		return err2
	} else {
		return nil
	}
}
