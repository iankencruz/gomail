package models

import (
	"database/sql"
	"time"
)

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Created   time.Time
}

type ContactModel struct {
	DB *sql.DB
}

func (c *ContactModel) Insert(firstname string, lastname string, email string, phone string) (int, error) {

	// SQL Query

	// Exec Query

	// Use LastInsertId() to get the of our newly inserted statement

	// The ID returned as a int64 so we convert it into an int type
	return 0, nil
}

func (c *ContactModel) Get(id int) (Contact, error) {
	return Contact{}, nil
}

func (c *ContactModel) GetAllContacts() ([]Contact, error) {
	stmt := `SELECT id, firstname, lastname, email, phone, created FROM contacts`

	// Execute our statement
	rows, err := c.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// Defer DB.Close() to ensure connection pool is terminated properly
	defer rows.Close()

	// Init an empty slice
	var contacts []Contact

	// Use rows.Next() to iterate through the rows in the resultset
	for rows.Next() {
		// Create a pointer to  a new zeroed out contact
		var c Contact

		// Use row.Scan() to copy the values from each field in the rows
		// to the new Contact Object
		// NOTE: Arguments to row.Scan() must be pointers to the place you want to copy
		// the data into & arguments must match the same amount as columns returned by your statement
		err = rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Phone, &c.Created)
		if err != nil {
			return nil, err
		}

		// Append iteration to contact slice
		contacts = append(contacts, c)
	}

	// When rows.Scan() finishes, we call rows.Err() to retrieve any error
	// that was encountered during iteration. NOTE: Important to always call this.
	// Dont assume that a successful iteration was completed over the whole resultset
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything OK, return the Contacts slice
	return contacts, nil
}
