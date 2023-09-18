package goexcel

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Contact struct {
	Firstname string
	Lastname  string
	Email     string
	Phone     int
}

// Returns a contact struct
func ReadExcelFile(input string) []Contact {

	var contacts []Contact

	// store values
	var fname []string
	var lname []string
	var email []string
	var phone []int

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
		// fmt.Println(row[1])
		fname = append(fname, row[1])
		lname = append(lname, row[2])
		email = append(email, row[3])
		// phone = append(phone, row[4])
		if i, err := strconv.Atoi(row[4]); err == nil {
			phone = append(phone, i)
		}

	}

	contacts.Firstname = fname
	contacts.Lastname = lname
	contacts.Email = email
	contacts.Phone = phone

	return contacts
}
