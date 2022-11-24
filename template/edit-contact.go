package template

import (
	"contact-program/helper"
	"contact-program/model"
	"fmt"
	"reflect"
	"strings"
)

func (template *contactTemplate) EditContact() {
	helper.ClearScreen()
	fmt.Println("Edit Contact")
	fmt.Println("===============")
	var search int
	fmt.Print("Input id yang akan di ubah: ")
	fmt.Scanln(&search)
	contact, err := template.contactHandler.GetContact(search)
	if err != nil {
		fmt.Println("id tidak ditemukan")
		var jeda string
		fmt.Scanln(&jeda)
		template.EditContact()
	} else {
		fmt.Println("Data Ditemukan")
		fmt.Println("===============")
		fmt.Println("Nama:", *contact.GetName())
		fmt.Println("Phone:", helper.PhoneToString(contact.GetPhoneDatas()))
		fmt.Println("Email:", *contact.GetEmail())
		fmt.Println("===============")
	}

	fmt.Println("")
	fmt.Println("")

	var name, email string
	var phones []string
	fmt.Println("Form Contact")
	fmt.Println("===============")
	template.InputName(&name)
	template.InputEmail(&email)
	template.InputPhone(&phones)
	template.contactHandler.UpdateContact(*contact.GetId(), name, email, phones)

	//Message berhasil
	fmt.Println("")
	fmt.Println("Data berhasil di update.")
	helper.BackHandler()
	Menu(template.db)
}

func (template *contactTemplate) InputName(name *string) {
	var input string
	fmt.Print("Name: ")
	fmt.Scanln(&input)

	if !ValidateName(&input) {
		fmt.Println("Name tidak boleh kosong")

		template.InputName(&input)
	}
	*name = input
}

func (template *contactTemplate) InputEmail(email *string) {
	fmt.Print("Email: ")
	fmt.Scanln(email)
}

func (template *contactTemplate) InputPhone(phones *[]string) {
	var phone string
	fmt.Print("Phone: ")
	fmt.Scanln(&phone)

	*phones = append(*phones, phone)
	var lagi string
	fmt.Print("Input phone lagi: ")
	fmt.Scanln(&lagi)
	if strings.ToLower(lagi) == "y" {
		template.InputPhone(phones)
	}
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
