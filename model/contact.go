package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Contact struct {
	id    int
	name  string `required:"true"`
	phone string `type:"number"`
	email string `type:"email"`
}

var Contacts []Contact

func GetLastId() int {
	if Contacts == nil {
		return 0
	} else {
		var tempId int
		for _, v := range Contacts {
			if tempId < v.id {
				tempId = v.id
			}
		}
		return tempId
	}
}

func (c *Contact) SearchById(id *int) error {
	foundContact := make(chan Contact)
	defer close(foundContact)

	go func() {
		for _, v := range Contacts {
			if v.id == *id {
				fmt.Println(v)
				foundContact <- v
				return
			}
		}
		foundContact <- *c
	}()

	*c = <-foundContact
	if c.id == 0 {
		return errors.New("data tidak ditemukan")
	}

	return nil
}

func (c *Contact) SearchByName(name *string) error {
	foundContact := make(chan Contact)
	defer close(foundContact)

	go func() {
		for _, v := range Contacts {
			if strings.EqualFold(v.name, *name) {
				foundContact <- v
				return
			}
		}
		foundContact <- *c
	}()

	*c = <-foundContact
	if c.id == 0 {
		return errors.New("data tidak ditemukan")
	}

	return nil
}

func GetIndex(id *int) (int, error) {
	for i, v := range Contacts {
		if v.id == *id {
			return i, nil
		}
	}

	return 0, errors.New("Id " + strconv.Itoa(*id) + " tidak ditemukan")
}

func (c *Contact) Add(datas ...map[string]interface{}) {
	for _, v := range datas {
		c.id = GetLastId() + 1
		c.name = v["name"].(string)
		c.phone = v["phone"].(string)
		c.email = v["email"].(string)
	}

	Contacts = append(Contacts, *c)
}

func (c *Contact) Edit(datas ...map[string]interface{}) {
	for _, v := range datas {
		c.name = v["name"].(string)
		c.phone = v["phone"].(string)
		c.email = v["email"].(string)
	}
	index, err := GetIndex(&c.id)
	if err == nil {
		Contacts[index].name = c.name
		Contacts[index].phone = c.phone
		Contacts[index].email = c.email
	}
}

func (c *Contact) Delete() {
	var index int
	for i, v := range Contacts {
		if v.id == c.id {
			index = i
		}
	}

	if index == len(Contacts)-1 {
		Contacts = Contacts[:index]
	} else if index == 0 {
		Contacts = Contacts[1:]
	} else {
		Contacts = append(Contacts[:index], Contacts[index+1:]...)
	}
}

func (c *Contact) GetFields() (int, string, string, string) {
	return c.id, c.name, c.phone, c.email
}
