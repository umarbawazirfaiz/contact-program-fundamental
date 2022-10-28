package template

import (
	"contact-program/controller"
	"contact-program/helper"
	"contact-program/model"
	"fmt"
)

func EditContact() {
	helper.ClearScreen()
	fmt.Println("Edit Contact")
	fmt.Println("===============")
	var id int
	fmt.Print("Input id yang akan di ubah: ")
	fmt.Scanln(&id)
	contact, err := model.SearchById(&id)
	if err != nil {
		EditContact()
	} else {
		_, name, phone, email := contact.GetFields()
		fmt.Println("Data Ditemukan")
		fmt.Println("===============")
		fmt.Println("Nama:", name)
		fmt.Println("Phone:", phone)
		fmt.Println("Email:", email)
		fmt.Println("===============")
	}

	fmt.Println("")
	fmt.Println("")

	var name, phone, email string
	fmt.Println("Form Contact")
	fmt.Println("===============")
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Print("Phone: ")
	fmt.Scanln(&phone)
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	controller.UpdateContactHandler(&contact, name, phone, email)

	//Message berhasil
	fmt.Println("")
	fmt.Println("Data berhasil di update.")
	helper.BackHandler()
	Menu()
}
