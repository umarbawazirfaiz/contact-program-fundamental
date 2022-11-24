package repository

import (
	"contact-program/model"
	"context"
	"database/sql"
)

type ContactRepository interface {
	FindAll(ctx context.Context) ([]model.Contact, error)
	Insert(ctx context.Context, contact model.Contact) (model.Contact, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, id int) (model.Contact, error)
	Update(ctx context.Context, tx *sql.Tx, contact model.Contact) (model.Contact, error)
}

type contactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *contactRepository {
	return &contactRepository{db}
}

func (repo *contactRepository) FindAll(ctx context.Context) ([]model.Contact, error) {
	var query string = "SELECT id, name, email FROM contact"
	var contacts []model.Contact

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return contacts, err
	}
	for rows.Next() {
		var contact model.Contact
		rows.Scan(contact.GetId(), contact.GetName(), contact.GetEmail())
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (repo *contactRepository) Insert(ctx context.Context, contact model.Contact) (model.Contact, error) {
	var query string = "INSERT INTO contact(name, email) VALUES(?,?)"

	res, err := repo.db.ExecContext(ctx, query, contact.GetName(), contact.GetEmail())
	if err != nil {
		return contact, nil
	}
	lastInsertId, _ := res.LastInsertId()
	id := int(lastInsertId)
	contact.SetId(&id)

	return contact, nil
}

func (repo *contactRepository) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	var query string = "DELETE FROM contact WHERE id=?"

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *contactRepository) FindById(ctx context.Context, contact_id int) (model.Contact, error) {
	var query string = "SELECT id, name, email FROM contact WHERE id=?"
	var contact model.Contact

	row := repo.db.QueryRowContext(ctx, query, contact_id)
	err := row.Scan(contact.GetId(), contact.GetName(), contact.GetEmail())
	if err != nil {
		return contact, err
	}

	return contact, nil
}

// method to update contact
func (repo *contactRepository) Update(ctx context.Context, tx *sql.Tx, contact model.Contact) (model.Contact, error) {
	query := "UPDATE contact SET name=?, email=? WHERE id=?"

	_, err := tx.ExecContext(ctx, query, contact.GetName(), contact.GetEmail(), contact.GetId())
	if err != nil {
		return contact, err
	}

	return contact, nil
}
