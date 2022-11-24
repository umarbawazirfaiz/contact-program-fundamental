package template

import (
	"contact-program/handler"
	"contact-program/helper"
	"contact-program/repository"
	"database/sql"
	"fmt"
	"os"
)

func Menu(db *sql.DB) {
	// Dependency Injection
	contactRepository := repository.NewContactRepository(db)
	phoneRepository := repository.NewPhoneRepository(db)
	contactHandler := handler.NewContactHandler(db, contactRepository, phoneRepository)
	contactTemplate := NewContactTemplate(db, contactHandler)

	helper.ClearScreen()
	fmt.Println("Menu")
	fmt.Println("=================")
	fmt.Println("1. List Contact")
	fmt.Println("2. Insert Contact")
	fmt.Println("3. Update Contact")
	fmt.Println("4. Delete Contact")
	fmt.Println("5. Exit")

	var menu int
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		contactTemplate.ListContact()
	case 2:
		contactTemplate.AddContact()
	case 3:
		contactTemplate.EditContact()
	case 4:
		contactTemplate.DeleteContact()
	case 5:
		os.Exit(0)
	default:
		Menu(db)
	}
}
