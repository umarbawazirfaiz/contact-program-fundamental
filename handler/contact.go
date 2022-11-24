package handler

import (
	"contact-program/helper"
	"contact-program/model"
	"contact-program/repository"
	"context"
	"database/sql"
)

type ContactHandler interface {
	GetContacts() ([]model.Contact, error)
	InsertContact(name, email string, phones []string) error
	DeleteContact(id int) error
	GetContact(id int) (model.Contact, error)
	UpdateContact(id int, name, email string, phones []string) error
}

type contactHandler struct {
	db                *sql.DB
	contactRepository repository.ContactRepository
	phoneRepository   repository.PhoneRepository
}

func NewContactHandler(db *sql.DB, contactRepository repository.ContactRepository, phoneRepository repository.PhoneRepository) *contactHandler {
	return &contactHandler{db, contactRepository, phoneRepository}
}

func (handler *contactHandler) GetContacts() ([]model.Contact, error) {
	ctx := context.Background()

	contacts, err := handler.contactRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	for i, v := range contacts {
		phoneDatas, err := handler.phoneRepository.GetPhonesByContactId(ctx, *v.GetId())
		if err != nil {
			return nil, err
		}
		contacts[i].SetPhoneDatas(phoneDatas)
	}

	return contacts, nil
}

func (handler *contactHandler) InsertContact(name, email string, phone []string) error {
	// make context background
	ctx := context.Background()

	var contact model.Contact
	var phoneDatas []model.PhoneData

	// init struct contact from pameter
	contact.SetName(&name)
	contact.SetEmail(&email)
	contact, err := handler.contactRepository.Insert(ctx, contact)
	if err != nil {
		return err
	}

	for _, v := range phone {
		var phone model.PhoneData
		phone.SetPhone(&v)
		phoneDatas = append(phoneDatas, phone)
	}

	tx, err := handler.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	phoneDatas, err = handler.phoneRepository.InsertPhones(ctx, tx, phoneDatas, *contact.GetId())
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	contact.SetPhoneDatas(phoneDatas)

	return nil
}

func (handler *contactHandler) GetContact(id int) (model.Contact, error) {
	ctx := context.Background()

	contact, err := handler.contactRepository.FindById(ctx, id)
	if err != nil {
		return contact, err
	}

	phoneDatas, err := handler.phoneRepository.GetPhonesByContactId(ctx, id)
	if err != nil {
		return contact, err
	}
	contact.SetPhoneDatas(phoneDatas)

	return contact, nil
}

func (handler *contactHandler) DeleteContact(id int) error {
	ctx := context.Background()

	tx, err := handler.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = handler.phoneRepository.DeleteByContactId(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = handler.contactRepository.Delete(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (handler *contactHandler) UpdateContact(id int, name, email string, phones []string) error {
	ctx := context.Background()

	tx, err := handler.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	contact := model.Contact{}
	contact.SetId(&id)
	contact.SetName(&name)
	contact.SetEmail(&email)
	contact, err = handler.contactRepository.Update(ctx, tx, contact)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = handler.phoneRepository.DeleteByContactId(ctx, tx, *contact.GetId())
	if err != nil {
		tx.Rollback()
		return err
	}

	phoneDatas := helper.PhonesToPhoneDatas(phones)
	_, err = handler.phoneRepository.InsertPhones(ctx, tx, phoneDatas, *contact.GetId())
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
