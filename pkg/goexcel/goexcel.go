package goexcel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Contact struct {
	Firstname string
	Lastname  string
	Email     string
	Phone     string
}

// NewEmail {
// 	Address string
// 	Name string
// }

// Returns a contact struct
func ReadExcelFile(input string) []Contact {

	var contacts []Contact
	var cPersons Contact

	f, err := excelize.OpenFile(input)
	if err != nil {
		fmt.Println(err)
		return contacts
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return contacts
	}

	for _, row := range rows[1:] {
		// append column values to slice
		cPersons.Firstname = row[1]
		cPersons.Lastname = row[2]
		cPersons.Email = row[3]
		cPersons.Phone = row[4]

		contacts = append(contacts, cPersons)
	}

	return contacts
}
