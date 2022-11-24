package model

type PhoneData struct {
	id        int
	contactId int
	phone     string
}

func (phoneData *PhoneData) GetId() *int {
	return &phoneData.id
}

func (phoneData *PhoneData) GetContactId() *int {
	return &phoneData.contactId
}

func (phoneData *PhoneData) GetPhone() *string {
	return &phoneData.phone
}

func (phoneData *PhoneData) SetId(id *int) {
	phoneData.id = *id
}

func (phoneData *PhoneData) SetContactId(contactId *int) {
	phoneData.contactId = *contactId
}

func (phoneData *PhoneData) SetPhone(phone *string) {
	phoneData.phone = *phone
}
