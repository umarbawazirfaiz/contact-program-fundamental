package template

import (
	"contact-program/handler"
	"contact-program/helper"
	"database/sql"
	"fmt"
)

type contactTemplate struct {
	db             *sql.DB
	contactHandler handler.ContactHandler
}

func NewContactTemplate(db *sql.DB, contactHandler handler.ContactHandler) *contactTemplate {
	return &contactTemplate{db, contactHandler}
}

func (c *contactTemplate) ListContact() {
	helper.ClearScreen()

	contacts, err := c.contactHandler.GetContacts()
	if err != nil {
		panic(err)
	}
	fmt.Println("==========================================================")
	fmt.Println("ID\tNama\t\tPhone\t\tEmail")
	fmt.Println("==========================================================")
	if len(contacts) == 0 {
		fmt.Println("Data kosong")
	} else {
		for _, v := range contacts {
			fmt.Printf("%v\t%v\t\t%s\t\t%v\n", *v.GetId(), *v.GetName(), helper.PhoneToString(v.GetPhoneDatas()), *v.GetEmail())
		}
	}
	fmt.Println("==========================================================")
	helper.BackHandler()
	Menu(c.db)
}
