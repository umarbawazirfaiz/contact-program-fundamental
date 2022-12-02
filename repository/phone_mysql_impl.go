package repository

import (
	"contact-program/model"
	"context"
	"database/sql"
)

type phoneRepository struct {
	db *sql.DB
}

func NewPhoneRepository(db *sql.DB) *phoneRepository {
	return &phoneRepository{db}
}

func (repo *phoneRepository) GetPhonesByContactId(ctx context.Context, contact_id int) ([]model.PhoneData, error) {
	var phoneDatas []model.PhoneData
	var query string = "SELECT id, phone FROM phone_data WHERE contact_id=?"

	rows, err := repo.db.QueryContext(ctx, query, contact_id)
	if err != nil {
		return phoneDatas, err
	}
	for rows.Next() {
		var phoneData model.PhoneData
		rows.Scan(phoneData.GetId(), phoneData.GetPhone())
		phoneDatas = append(phoneDatas, phoneData)
	}
	return phoneDatas, err
}

func (repo *phoneRepository) InsertPhones(ctx context.Context, tx *sql.Tx, phoneDatas []model.PhoneData, contact_id int) ([]model.PhoneData, error) {
	var query string = "INSERT INTO phone_data(contact_id, phone) VALUES(?,?)"

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, v := range phoneDatas {
		res, err := stmt.ExecContext(ctx, contact_id, v.GetPhone())
		if err != nil {
			return nil, err
		}
		lastInsertId, _ := res.LastInsertId()
		id := int(lastInsertId)
		v.SetId(&id)
	}
	return phoneDatas, nil
}

func (repo *phoneRepository) DeleteByContactId(ctx context.Context, tx *sql.Tx, contact_id int) error {
	query := "DELETE FROM phone_data WHERE contact_id=?"

	_, err := tx.ExecContext(ctx, query, contact_id)
	if err != nil {
		return err
	}

	return nil
}
