package repository

import (
	"contact-program/model"
	"context"
	"database/sql"
)

type PhoneRepository interface {
	GetPhonesByContactId(ctx context.Context, contact_id int) ([]model.PhoneData, error)
	InsertPhones(ctx context.Context, tx *sql.Tx, phoneDatas []model.PhoneData, contact_id int) ([]model.PhoneData, error)
	DeleteByContactId(ctx context.Context, tx *sql.Tx, contact_id int) error
}
