package template

import (
	"contact-program/helper"
	"fmt"
)

func (template *contactTemplate) AddContact() {
	helper.ClearScreen()
	var name, email string
	var phones []string

	fmt.Println("Add Contact")
	fmt.Println("===============")
	template.InputName(&name)
	template.InputEmail(&email)
	template.InputPhone(&phones)

	err := template.contactHandler.InsertContact(name, email, phones)
	if err != nil {
		panic(err)
	}

	//Message berhasil
	fmt.Println("")
	fmt.Println("Data berhasil di input.")
	helper.BackHandler()
	Menu(template.db)
}
