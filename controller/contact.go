package controller

import (
	"contact-program/interfaces"
	"strconv"
)

func UpdateContactHandler(co interfaces.ContactInterface, name, phone, email string) {
	datas := map[string]interface{}{
		"name":  name,
		"phone": phone,
		"email": email,
	}
	co.Edit(datas)
}

func InsertContactHandler(co interfaces.ContactInterface, name, phone, email string) {
	datas := map[string]interface{}{
		"name":  name,
		"phone": phone,
		"email": email,
	}
	co.Add(datas)
}

func DeleteContactHandler(co interfaces.ContactInterface) {
	co.Delete()
}

func InsertCustomerHandler(co interfaces.ContactInterface, name string, age int) {
	datas := map[string]interface{}{
		"name": name,
		"age":  age,
	}
	co.Add(datas)
}

func SearchContact(co interfaces.ContactInterface, search string) error {
	ses, err := strconv.Atoi(search)
	if err == nil {
		err = co.SearchById(&ses)
		if err != nil {
			return err
		}
		return nil
	} else {
		err = co.SearchByName(&search)
		if err != nil {
			return err
		}
		return nil
	}
}
