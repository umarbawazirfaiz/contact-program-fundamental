package template

import (
	"contact-program/controller"
	"contact-program/helper"
	"contact-program/model"
	"fmt"
	"reflect"
)

func AddContact() {
	helper.ClearScreen()
	var phone, email string
	fmt.Println("Add Contact")
	fmt.Println("===============")
	name := InputName()
	fmt.Print("Phone: ")
	fmt.Scanln(&phone)
	fmt.Print("Email: ")
	fmt.Scanln(&email)

	var contact model.Contact
	controller.InsertContactHandler(&contact, name, phone, email)

	//Message berhasil
	fmt.Println("")
	fmt.Println("Data berhasil di input.")
	helper.BackHandler()
	Menu()
}

func InputName() string {
	var name string
	fmt.Print("Name: ")
	fmt.Scanln(&name)

	if !ValidateName(&name) {
		fmt.Println("Name tidak boleh kosong")

		InputName()
	}
	return name
}

func ValidateName(name *string) bool {
	var c model.Contact
	typeOf := reflect.TypeOf(c)
	if typeOf.Field(1).Tag.Get("required") == "true" {
		if *name == "" {
			return false
		}
	}
	return true
}
