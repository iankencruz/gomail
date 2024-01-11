package models

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

func (c *Contact) GetContacts() ([]Contact, error) {
	var contacts []Contact

	contacts = append(contacts, Contact{
		ID:        1,
		FirstName: "Ian",
		LastName:  "Dela Cruz",
		Email:     "iancruz@test.com",
		Phone:     "0423382922",
	})

	return contacts, nil
}
