package model

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
