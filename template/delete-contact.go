package template

import (
	"contact-program/controller"
	"contact-program/helper"
	"contact-program/model"
	"fmt"
)

func DeleteContact() {
	helper.ClearScreen()
	fmt.Println("Delete Contact")
	fmt.Println("===============")
	var search string
	fmt.Print("Input id/nama yang akan di delete: ")
	fmt.Scanln(&search)
	var contact model.Contact
	err := controller.SearchContact(&contact, search)
	if err != nil {
		fmt.Println(err.Error())
		var jeda string
		fmt.Scanln(&jeda)
		DeleteContact()
	} else {
		_, name, phone, email := contact.GetFields()
		fmt.Println("Data Ditemukan")
		fmt.Println("===============")
		fmt.Println("Nama:", name)
		fmt.Println("Phone:", phone)
		fmt.Println("Email:", email)
		fmt.Println("===============")
	}

	confirmDelete()
	controller.DeleteContactHandler(&contact)

	//Message berhasil
	fmt.Println("")
	fmt.Println("Data berhasil di didelete.")
	helper.BackHandler()
	Menu()
}

func confirmDelete() {
	fmt.Print("Apakah yakin ingin dihapus(y/t)")
	var confirm string
	fmt.Scanln(&confirm)
	switch confirm {
	case "y":
	case "t":
		Menu()
	default:
		confirmDelete()
	}
}
