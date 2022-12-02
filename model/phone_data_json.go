package model

type PhoneDataJson struct {
	Id        int
	ContactId int
	Phone     string
}

func (m *PhoneDataJson) ToPhone() PhoneData {
	return PhoneData{
		id:        m.Id,
		contactId: m.ContactId,
		phone:     m.Phone,
	}
}
