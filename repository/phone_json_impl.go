package repository

import (
	"contact-program/model"
	"context"
	"database/sql"
	"encoding/json"
	"os"
)

type phoneJsonRepository struct {
}

func NewPhoneJsonRepository() *phoneJsonRepository {
	return &phoneJsonRepository{}
}

func (repo *phoneJsonRepository) saveToJson(ctx context.Context, phoneDatas *[]model.PhoneData) error {
	file, err := os.Create("database/json/phone.json")
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)

	phoneDatasJson := []model.PhoneDataJson{}
	for _, v := range *phoneDatas {
		phoneJson := v.ToPhoneJson()
		phoneDatasJson = append(phoneDatasJson, phoneJson)
	}

	err = encoder.Encode(phoneDatasJson)
	if err != nil {
		panic(err)
	}

	return nil
}

func (repo *phoneJsonRepository) FindAll(ctx context.Context) ([]model.PhoneData, error) {
	var phoneDatasJson []model.PhoneDataJson

	file, err := os.Open("database/json/phone.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&phoneDatasJson)
	if err != nil {
		panic(err)
	}

	phoneDatas := []model.PhoneData{}
	for _, v := range phoneDatasJson {
		phoneData := v.ToPhone()
		phoneDatas = append(phoneDatas, phoneData)
	}

	return phoneDatas, nil
}

func (repo *phoneJsonRepository) getLastId(ctx context.Context) (int, error) {
	phoneDatas, _ := repo.FindAll(ctx)

	tempId := 0
	for _, v := range phoneDatas {
		if tempId < *v.GetId() {
			tempId = *v.GetId()
		}
	}

	return tempId, nil
}

func (repo *phoneJsonRepository) GetPhonesByContactId(ctx context.Context, contact_id int) ([]model.PhoneData, error) {
	phoneDatas, err := repo.FindAll(ctx)

	phoneDatasByContactId := []model.PhoneData{}
	for _, v := range phoneDatas {
		if *v.GetContactId() == contact_id {
			phoneDatasByContactId = append(phoneDatasByContactId, v)
		}
	}

	return phoneDatasByContactId, err
}

func (repo *phoneJsonRepository) InsertPhones(ctx context.Context, tx *sql.Tx, phoneDatas []model.PhoneData, contact_id int) ([]model.PhoneData, error) {
	phoneDatasDB, err := repo.FindAll(ctx)
	if err != nil {
		return phoneDatasDB, err
	}

	id, _ := repo.getLastId(ctx)
	for _, v := range phoneDatas {
		id++
		v.SetId(&id)
		v.SetContactId(&contact_id)
		phoneDatasDB = append(phoneDatasDB, v)
	}

	repo.saveToJson(ctx, &phoneDatasDB)

	return phoneDatas, err
}

func (repo *phoneJsonRepository) DeleteByContactId(ctx context.Context, tx *sql.Tx, contact_id int) error {
	phoneDatas, err := repo.FindAll(ctx)
	if err != nil {
		return err
	}

	for _, v := range phoneDatas {
		if *v.GetContactId() == contact_id {
			repo.Delete(ctx, *v.GetId())
		}
	}

	return nil
}

func (repo *phoneJsonRepository) Delete(ctx context.Context, id int) error {
	phoneDatas, err := repo.FindAll(ctx)
	if err != nil {
		return err
	}

	for i, v := range phoneDatas {
		if *v.GetId() == id {
			if i == 0 {
				phoneDatas = phoneDatas[1:]
			} else if len(phoneDatas)-1 == i {
				phoneDatas = phoneDatas[:len(phoneDatas)-1]
			} else {
				phoneDatas = append(phoneDatas[:i], phoneDatas[i:]...)
			}
		}
	}

	repo.saveToJson(ctx, &phoneDatas)

	return nil
}
