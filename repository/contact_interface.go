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
