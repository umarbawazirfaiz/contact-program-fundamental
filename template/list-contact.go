package template

import (
	"contact-program/helper"
	"contact-program/model"
	"fmt"
)

func ListContact() {
	helper.ClearScreen()
	fmt.Println("==========================================================")
	fmt.Println("ID\tNama\t\tPhone\t\tEmail")
	fmt.Println("==========================================================")
	if len(model.Contacts) == 0 {
		fmt.Println("Data kosong")
	} else {
		for _, v := range model.Contacts {
			id, name, phone, email := v.GetFields()
			fmt.Printf("%v\t%v\t\t%v\t\t%v\n", id, name, phone, email)
		}
	}
	fmt.Println("==========================================================")
	helper.BackHandler()
	Menu()
}
