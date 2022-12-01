package repository

import (
	"contact-program/model"
	"context"
	"database/sql"
	"encoding/json"
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
	file, err := os.Create("/database/json/contact.json")
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(&contacts)

	return nil
}

func (repo *contactJsonRepository) FindAll(ctx context.Context) ([]model.Contact, error) {
	var contacts []model.Contact

	file, err := os.Open("databse/json/contact.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		panic(err)
	}

	return contacts, nil
}

func (repo *contactJsonRepository) Insert(ctx context.Context, contact model.Contact) (model.Contact, error) {
	contacts, err := repo.FindAll(ctx)
	if err != nil {
		return contact, err
	}

	id, err := repo.getLastId(ctx)
	if err != nil {
		return contact, err
	}
	contact.SetId(&id)

	contacts = append(contacts, contact)

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
		if v.GetId() == contact.GetId() {
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
		if v.GetId() == &contact_id {
			contact = v
		}
	}

	return contact, nil
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
