package template

import (
	"contact-program/handler"
	"contact-program/helper"
	"database/sql"
	"fmt"

	"github.com/buger/goterm"
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

	box := goterm.NewBox(100|goterm.PCT, (len(contacts)+1)*2-(len(contacts)-1), 0)
	table := goterm.NewTable(0, 5, 1, ' ', 0)
	fmt.Fprintf(table, "ID\t| Nama\t| Phone\t| Email\n")
	if len(contacts) == 0 {
		fmt.Fprintf(table, "Data kosong")
	} else {
		for _, v := range contacts {
			fmt.Fprintf(table, "%v\t| %v\t| %s\t| %v\n", *v.GetId(), *v.GetName(), helper.PhoneToString(v.GetPhoneDatas()), *v.GetEmail())
		}
	}
	fmt.Fprint(box, table)

	fmt.Println("Data Contact")
	fmt.Println(box)

	helper.BackHandler()
	Menu(c.db)
}
