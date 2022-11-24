package main

import (
	"contact-program/database"
	"contact-program/template"
)

func main() {
	db := database.GetConnection()
	template.Menu(db)
}
