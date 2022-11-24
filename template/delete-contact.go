package template

import (
	"contact-program/helper"
	"database/sql"
	"fmt"
)

func (c *contactTemplate) DeleteContact() {
	helper.ClearScreen()
	fmt.Println("Delete Contact")
	fmt.Println("===============")
	var search int
	fmt.Print("Input id yang akan di delete: ")
	fmt.Scanln(&search)
	contact, err := c.contactHandler.GetContact(search)
	if err != nil {
		fmt.Println("id tidak ditemukan")
		var jeda string
		fmt.Scanln(&jeda)
		c.DeleteContact()
	} else {
		fmt.Println("Data Ditemukan")
		fmt.Println("===============")
		fmt.Println("Nama:", *contact.GetName())
		fmt.Println("Phone:", helper.PhoneToString(contact.GetPhoneDatas()))
		fmt.Println("Email:", *contact.GetEmail())
		fmt.Println("===============")
	}

	confirmDelete(c.db)
	c.contactHandler.DeleteContact(*contact.GetId())

	//Message berhasil
	fmt.Println("")
	fmt.Println("Data berhasil di didelete.")
	helper.BackHandler()
	Menu(c.db)
}

func confirmDelete(db *sql.DB) {
	fmt.Print("Apakah yakin ingin dihapus(y/t)")
	var confirm string
	fmt.Scanln(&confirm)
	switch confirm {
	case "y":
	case "t":
		Menu(db)
	default:
		confirmDelete(db)
	}
}
