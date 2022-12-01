package model

type Contact struct {
	id         int
	name       string `required:"true"`
	phoneDatas []PhoneData
	email      string `type:"email"`
}

type ContactJson struct {
	Id    int
	Name  string
	Email string
}

func (contactJson *ContactJson) ToContact() Contact {
	contact := Contact{
		id:    contactJson.Id,
		name:  contactJson.Name,
		email: contactJson.Email,
	}

	return contact
}

func (contact *Contact) ToContactJson() ContactJson {
	contactJson := ContactJson{
		Id:    contact.id,
		Name:  contact.name,
		Email: contact.email,
	}

	return contactJson
}

func (contact *Contact) GetId() *int {
	return &contact.id
}

func (contact *Contact) GetName() *string {
	return &contact.name
}

func (contact *Contact) GetEmail() *string {
	return &contact.email
}

func (contact *Contact) GetPhoneDatas() []PhoneData {
	return contact.phoneDatas
}

func (contact *Contact) SetId(id *int) {
	contact.id = *id
}

func (contact *Contact) SetName(name *string) {
	contact.name = *name
}

func (contact *Contact) SetEmail(email *string) {
	contact.email = *email
}

// set phone data to struct contact
func (contact *Contact) SetPhoneDatas(phoneDatas []PhoneData) {
	contact.phoneDatas = phoneDatas
}
