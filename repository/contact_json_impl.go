package repository

import (
	"contact-program/model"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type contactJsonRepository struct {
}

func NewContactJsonRepository() *contactJsonRepository {
	return &contactJsonRepository{}
}

func (repo *contactJsonRepository) getLastId(ctx context.Context) (int, error) {
	contacts, _ := repo.FindAll(ctx)

	tempId := 0
	for _, v := range contacts {
		if tempId < *v.GetId() {
			tempId = *v.GetId()
		}
	}

	return tempId, nil
}

func (repo *contactJsonRepository) saveToJson(ctx context.Context, contacts *[]model.Contact) error {
	file, err := os.Create("database/json/contact.json")
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)

	contactsJson := []model.ContactJson{}
	for _, v := range *contacts {
		contactJson := v.ToContactJson()
		contactsJson = append(contactsJson, contactJson)
	}

	err = encoder.Encode(contactsJson)
	if err != nil {
		panic(err)
	}

	return nil
}

func (repo *contactJsonRepository) FindAll(ctx context.Context) ([]model.Contact, error) {
	var contactsJson []model.ContactJson

	file, err := os.Open("database/json/contact.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contactsJson)
	if err != nil {
		panic(err)
	}

	contacts := []model.Contact{}
	for _, v := range contactsJson {
		contact := v.ToContact()
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (repo *contactJsonRepository) Insert(ctx context.Context, contact model.Contact) (model.Contact, error) {
	contacts, err := repo.FindAll(ctx)
	if err != nil {
		return contact, err
	}

	lastId, err := repo.getLastId(ctx)
	if err != nil {
		return contact, err
	}
	id := lastId + 1
	contact.SetId(&id)

	fmt.Println(contact)

	contacts = append(contacts, contact)
	fmt.Println(contacts)

	repo.saveToJson(ctx, &contacts)

	return contact, nil
}

func (repo *contactJsonRepository) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	contacts, err := repo.FindAll(ctx)
	if err != nil {
		return err
	}

	contact, err := repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	for i, v := range contacts {
		if *v.GetId() == *contact.GetId() {
			if i == 0 {
				contacts = contacts[1:]
			} else if len(contacts)-1 == i {
				contacts = contacts[:len(contacts)-1]
			} else {
				contacts = append(contacts[:i], contacts[i:]...)
			}
		}
	}

	repo.saveToJson(ctx, &contacts)

	return nil
}

func (repo *contactJsonRepository) FindById(ctx context.Context, contact_id int) (model.Contact, error) {
	var contact model.Contact

	contacts, err := repo.FindAll(ctx)
	if err != nil {
		return contact, err
	}

	for _, v := range contacts {
		if *v.GetId() == contact_id {
			return v, nil
		}
	}

	return contact, errors.New("data tidak ditemukan")
}

func (repo *contactJsonRepository) Update(ctx context.Context, tx *sql.Tx, contact model.Contact) (model.Contact, error) {
	contacts, err := repo.FindAll(ctx)
	if err != nil {
		return contact, err
	}

	find_contact, err := repo.FindById(ctx, *contact.GetId())
	if err != nil {
		return contact, err
	}

	find_contact.SetEmail(contact.GetEmail())
	find_contact.SetName(contact.GetName())

	for i, v := range contacts {
		if v.GetId() == find_contact.GetId() {
			contacts[i] = find_contact
		}
	}

	repo.saveToJson(ctx, &contacts)

	return find_contact, nil
}
