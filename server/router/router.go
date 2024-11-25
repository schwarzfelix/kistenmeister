package router

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/schwarzfelix/kistenmeister/server/database"
	"github.com/schwarzfelix/kistenmeister/server/model"
)

// Funktion zum Prüfen nach Error -> verwendet in den Router-Funktionen
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SetupRouter() *gin.Engine {

	r := gin.Default()
	// Load HTML templates
	r.LoadHTMLFiles("templates/pictures.html")

	//https://stackoverflow.com/questions/29418478/go-gin-framework-cors
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		//TODO CORS einschränken
		//AllowHeaders:     []string{"Origin Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers"},
		AllowHeaders: []string{"*"},
		//ExposeHeaders:    []string{"Content-Length"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		/*AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},*/
		MaxAge: 12 * time.Hour,
	}))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Lese alle Kisten in der DB
	r.GET("/kisten", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		alleKisten, err := database.GetBoxes(token)
		checkErr(err)

		if len(alleKisten) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Alle Kisten": alleKisten})
		}
	})

	// Gebe alle Personen in der DB aus
	r.GET("/personen", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		allePersonen, err := database.GetPersonen(token)
		checkErr(err)

		if len(allePersonen) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Alle Personen": allePersonen})
		}
	})

	// Lese alle Kommentare in der DB
	r.GET("/kommentare", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		alleKommentare, err := database.GetComments(token)
		checkErr(err)

		if len(alleKommentare) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Alle Kommentare": alleKommentare})
		}
	})

	// Lese alle Merken in der DB
	r.GET("/alleMerklisteneinträge", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		alleMerken, err := database.GetStars(token)
		checkErr(err)

		if len(alleMerken) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Alle Merklisteneinträge": alleMerken})
		}
	})

	// Lese alle Bilder in der DB
	r.GET("/bilder", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		alleBilder, err := database.GetPictures(token)
		checkErr(err)

		if len(alleBilder) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		}
		c.JSON(http.StatusOK, alleBilder)
	})

	// Lese die Daten einer einzelnen Kiste
	r.GET("/kiste/:id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		id := c.Param("id")
		// Prüfe ob id int ist, sonst error ausgeben - check über cast to int
		idi, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}
		singleBox, err := database.GetBox(idi, token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			return
		} else if singleBox == (model.Kiste{}) {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Kiste": singleBox})
		}
	})

	// Lese einzelne Person nach ID
	r.GET("/person", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]
		einzelneperson, err := database.GetPerson(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			return
		} else if einzelneperson == (model.Person{}) {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Person": einzelneperson})
		}
	})

	// Lese alle Kommentare zu gegebener Kiste
	r.GET("/kommentare/:kiste_id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		box_id := c.Param("kiste_id")
		// Prüfe ob id int ist, sonst error ausgeben - check über cast to int
		box_idi, err := strconv.Atoi(box_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}

		allcomments, err := database.GetComment(box_idi, token)
		checkErr(err)

		if len(allcomments) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Kommentare": allcomments})
		}
	})

	// Lese alle Merken zu gegebener Person
	r.GET("/merklisteneinträge", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		allstars, err := database.GetStar(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		} else if len(allstars) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Merklisteneinträge": allstars})
		}
	})

	// Lese alle Bilder zu gegebener Kiste
	r.GET("/bilder/:kiste_id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		box_id := c.Param("kiste_id")
		// Prüfe ob id int ist, sonst error ausgeben
		// check über cast to int
		box_idi, err := strconv.Atoi(box_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}

		allpictures, err := database.GetPicture(box_idi, token)
		checkErr(err)

		if len(allpictures) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		}
		c.JSON(http.StatusOK, allpictures)
	})

	// Einzelne Kiste je id Löschen
	r.DELETE("/kiste/:id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		id := c.Params.ByName("id")
		// Prüfe ob id int ist, sonst error ausgeben
		// check über cast to int
		idi, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}
		// Functionsaufruf zum Löschen der Kiste
		deletedBox, err := database.DeleteBox(idi, token)

		// Lese die Daten der gelöschten Kiste ab
		if err != nil || deletedBox == (model.Kiste{}) {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Kiste gelöscht": deletedBox})
		}
	})

	// Einzelner Kommentar nach id löschen
	r.DELETE("/kommentar/:id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		id := c.Params.ByName("id")
		// Prüfe ob id int ist, sonst error ausgeben
		// check über cast to int
		idi, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}

		// Functionsaufruf zum Löschen des Kommentars
		deletedComment, err := database.DeleteComment(idi, token)

		// Lese die Daten des gelöschten Kommentars ab
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Kommentar gelöscht": deletedComment})
		}
	})

	// Einzelne Merke nach id löschen
	r.DELETE("/merklisteneintrag/:id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		id := c.Params.ByName("id")
		// Prüfe ob id int ist, sonst error ausgeben
		// check über cast to int
		idi, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}

		// Functionsaufruf zum Löschen der Merke
		deletedStar, err := database.DeleteStar(idi, token)

		// Lese die Daten der gelöschten Merke ab
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Merk gelöscht": deletedStar})
		}
	})

	// Einzelnes Bild nach id löschen
	r.DELETE("/bild/:id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		id := c.Params.ByName("id")
		// Prüfe ob id int ist, sonst error ausgeben
		// check über cast to int
		idi, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}

		// Functionsaufruf zum Löschen des Bilds
		deletedPicture, err := database.DeletePicture(idi, token)

		// Lese die Daten des gelöschten Bilds ab
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Bild gelöscht": deletedPicture})
		}
	})

	// Einzelne Person nach id löschen
	r.DELETE("/person/:id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		id := c.Param("id")
		// Prüfe ob id int ist, sonst error ausgeben - check über cast to int
		idi, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}

		// Functionsaufruf zum Löschen der Person
		deletedPerson, err := database.DeletePerson(idi, token)

		// Lese die Daten des gelöschten Bilds ab
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Person gelöscht": deletedPerson})
		}
	})

	// Einzelne Kiste je id aktualisieren
	r.PUT("/kiste/:id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		id := c.Params.ByName("id")
		// Prüfe ob id int ist, sonst error ausgeben
		// check über cast to int
		idi, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}

		name := c.PostForm("Name")
		description := c.PostForm("Beschreibung")
		verantwortlicher := c.PostForm("Verantwortlicher")
		location := c.PostForm("Ort")

		updatedBoxValues := model.Kiste{
			Name:             name,
			Beschreibung:     description,
			Änderungsdatum:   time.Now(),
			Verantwortlicher: verantwortlicher,
			Ort:              location,
		}

		// Functionsaufruf zum aktualisieren der Kiste
		updatedBox, err := database.UpdateBox(idi, updatedBoxValues, token)

		// Lese die Daten der aktualisierten Kiste ab
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Kiste aktualisiert": updatedBox})
		}
	})

	// Einzelne Person je id aktualisieren
	r.PUT("/person", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		name := c.PostForm("Name")
		email := c.PostForm("Email")
		passwort := c.PostForm("Passwort")
		lizenz := c.PostForm("Lizenz")

		updatedPersonValues := model.Person{
			Name:           name,
			Email:          email,
			Passwort:       passwort,
			Lizenz:         lizenz,
			Änderungsdatum: time.Now(),
		}

		// Functionsaufruf zum aktualisieren der Personendaten
		updatedPerson, err := database.UpdatePerson(updatedPersonValues, token)

		// Lese die Daten der aktualisierten Person ab
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Keine Einträge gefunden"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Personendaten aktualisiert": updatedPerson})
		}
	})

	// Neue Kiste erstellen
	r.POST("/kiste", func(c *gin.Context) {

		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		name := c.PostForm("Name")
		description := c.PostForm("Beschreibung")
		location := c.PostForm("Ort")

		newBox := model.Kiste{
			Name:             name,
			Beschreibung:     description,
			Erstellungsdatum: time.Now(),
			Änderungsdatum:   time.Now(),
			Ort:              location,
		}

		newID, err := database.CreateBox(newBox, token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Kiste könnte nicht erstellt werden"})
			c.JSON(http.StatusBadRequest, err)
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Neue Kiste mit ID": newID})
		}
	})

	// Neue Person zur DB hinzufügen
	r.POST("/person", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		name := c.PostForm("Name")
		email := c.PostForm("Email")
		passwort := c.PostForm("Passwort")
		lizenz := c.PostForm("Lizenz")

		newPerson := model.Person{
			Name:             name,
			Email:            email,
			Passwort:         passwort,
			Lizenz:           lizenz,
			Erstellungsdatum: time.Now(),
			Änderungsdatum:   time.Now(),
			Active:           false,
			Token:            "",
		}

		newID, err := database.CreatePerson(newPerson, token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Person könnte nicht zu der Datenbank hinzugefügt werden"})
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Neue Person mit ID": newID})
		}
	})

	// Neuer kommentar zu gegebener Kiste erstellen
	r.POST("/kommentar/:kiste_id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		boxID := c.Params.ByName("kiste_id")
		// Prüfe ob id int ist, sonst error ausgeben
		// check über cast to int
		box_idi, err := strconv.Atoi(boxID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}
		comment := c.PostForm("Kommentar")

		newComment := model.Kommentar{
			Kommentar:        comment,
			Erstellungsdatum: time.Now(),
			Kiste_id:         box_idi,
		}

		newID, err := database.CreateComment(newComment, token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			return
		} else if newID == -1 {
			c.JSON(http.StatusBadRequest, gin.H{"Keine Kiste mit der angegebenen Kiste_id gefunden": nil})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Neuer Kommentar mit ID": newID})
		}
	})

	// Neuer Merk zu gegebener Person erstellen
	r.POST("/merklisteneintrag", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		box_id := c.PostForm("kiste_id")
		// Prüfe ob box_id int ist, sonst error ausgeben
		// check über cast to int
		box_idi, err := strconv.Atoi(box_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}

		newStar := model.Merke{
			Erstellungsdatum: time.Now(),
			Kiste_id:         box_idi,
		}

		newID, err := database.CreateStar(newStar, token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			return
		} else if newID == -1 {
			c.JSON(http.StatusBadRequest, gin.H{"Kein Merk zur angegebenen Person gefunden": nil})
			return
		} else if newID == -2 {
			c.JSON(http.StatusBadRequest, gin.H{"Keine Kiste mit der angegebenen Kiste_id gefunden": nil})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Neuer Merk mit ID": newID})
		}
	})

	// Neues Bild zu gegebener Kiste erstellen
	r.POST("/bild/:kiste_id", func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header missing"})
			return
		}

		// The token should be in the format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Authorization header format"})
			return
		}
		token := tokenParts[1]

		boxID := c.Params.ByName("kiste_id")
		// Prüfe ob id int ist, sonst error ausgeben
		// check über cast to int
		box_idi, err := strconv.Atoi(boxID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Ungültiges ID-Format"})
			return
		}

		// Retrieve the picture file
		file, err := c.FormFile("bild")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Bild-Datei fehlt"})
			return
		}

		// Open the file
		openedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Fehler beim Öffnen der Bild-Datei"})
			return
		}
		defer openedFile.Close()

		// Read the file content into a byte slice
		fileContent, err := io.ReadAll(openedFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Fehler beim Lesen der Bild-Datei"})
			return
		}
		// log.Printf("File content size: %d bytes", len(fileContent))
		newPicture := model.Bild{
			Bild:             fileContent,
			Erstellungsdatum: time.Now(),
			Kiste_id:         box_idi,
		}

		newID, err := database.CreatePicture(newPicture, token)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			return
		} else if newID == -1 {
			c.JSON(http.StatusBadRequest, gin.H{"Keine Kiste mit der angegebenen Kiste_id gefunden": nil})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"Neues Bild mit ID": newID})
		}
	})

	// Login Verfahren
	type LoginInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	r.POST("/login", func(c *gin.Context) {
		var input LoginInput
		err := c.ShouldBindJSON(&input)
		checkErr(err)
		db, err := sql.Open("sqlite3", "./Kistenmeister.db")
		checkErr(err)
		defer db.Close()

		var id int
		var email string
		var password string
		var active bool
		token, err := generateToken()
		checkErr(err)
		err1 := db.QueryRow("SELECT ID, Email, Passwort, Active FROM PERSONEN WHERE Email = ? AND Passwort = ?", input.Email, input.Password).Scan(&id, &email, &password, &active)
		if err1 != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ungültige Anmeldeinformationen"})
			return
		} else if !active {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Aktivieren Sie bitte das Konto über den versendeten Link in deiner E-Mail-ID, bevor Sie sich anmelden können"})
			return
		} else {
			// Functionsaufruf zum aktualisieren des Tokens nach Person ID
			err2 := database.UpdateToken(id, token)
			if err2 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error": err2})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Login erfolgreich", "Token": token})
		}
	})

	// Registrierung Verfahren
	type RegisterInput struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Lizenz   string `json:"lizenz" binding:"required"`
	}

	r.POST("/register", func(c *gin.Context) {
		var reginput RegisterInput
		err := c.ShouldBindJSON(&reginput)
		checkErr(err)
		db, err := sql.Open("sqlite3", "./Kistenmeister.db")
		checkErr(err)
		defer db.Close()

		row := db.QueryRow("SELECT COUNT(*) FROM Personen")
		var count int
		err = row.Scan(&count)
		checkErr(err)

		var max_id int
		if count == 0 {
			// No rows in the table, set maxValue to 0
			max_id = 0
		} else {
			// Query max(Id) from the data table
			query := "SELECT MAX(ID) FROM Personen"
			// Execute the query and scan the result into `max_id`
			err = db.QueryRow(query).Scan(&max_id)
			checkErr(err)
		}

		var new_id = max_id + 1 // store new_id as max(Id) + 1
		var falseactive bool = false
		token, err := generateToken()
		checkErr(err)

		create_stmt, err := db.Prepare("INSERT INTO Personen (ID, Name, Email, Passwort, Lizenz, Ersteller, Erstellungsdatum, Änderer, Änderungsdatum, Active, Token) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
		checkErr(err)
		defer create_stmt.Close()

		_, err = create_stmt.Exec(new_id, reginput.Name, reginput.Email, reginput.Password, reginput.Lizenz, reginput.Name, time.Now(), reginput.Name, time.Now(), falseactive, token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ungültige Eingabedaten"})
			fmt.Println(err)
			return
		}

		err = sendActivationEmail(reginput.Email, token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Aktivierungsmail könnte nicht gesendet werden"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Link zur Registrierung an die Email-Adresse gesendet"})
	})

	r.GET("/activate", func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültiger Aktivierungstoken"})
			return
		}

		db, err := sql.Open("sqlite3", "./Kistenmeister.db")
		checkErr(err)
		defer db.Close()

		update_stmt, err := db.Prepare("UPDATE Personen SET Active = 1 WHERE Token = ?")
		checkErr(err)
		defer update_stmt.Close()

		result, err := update_stmt.Exec(token)
		checkErr(err)

		rowsAffected, err := result.RowsAffected()
		checkErr(err)

		if rowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültiger Aktivierungstoken"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Konto erfolgreich aktiviert"})
	})

	return r
}

// Helper function to generate a random token
func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Function to send activation email
func sendActivationEmail(email, token string) error {
	from := ""
	password := ""
	// Gmail's SMTP server details
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf("Subject: Kontoaktivierung\n\nBitte den unteren Link klicken, um Ihr Benutzerkonto zu aktivieren:\nhttp://localhost:8080/activate?token=%s", token))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, message)
	checkErr(err)
	if err != nil {
		log.Printf("Aktivierungsmail könnte nicht gesendet werden: %v", err)
	}
	return nil
}
