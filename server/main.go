package main

import (
	"github.com/schwarzfelix/kistenmeister/server/database"
	"github.com/schwarzfelix/kistenmeister/server/router"
)

func main() {
	database.ConnectDatabase()
	database.CreateTables()
	r := router.SetupRouter()
	r.Run(":8080")
}
