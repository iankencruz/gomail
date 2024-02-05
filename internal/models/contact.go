package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
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
	stmt := `INSERT INTO contacts (firstname, lastname, email, phone, created) VALUES (?,?,?,?,?)`

	// Exec Query
	result, err := c.DB.Exec(stmt, firstname, lastname, email, phone, time.Now())
	if err != nil {
		return 0, nil
	}

	// Use LastInsertId() to get the of our newly inserted statement
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	fmt.Printf("Inserted Contact with ID: %s || ", strconv.Itoa(int(id)))

	// The ID returned as a int64 so we convert it into an int type
	return int(id), nil
}

func (c *ContactModel) Delete(id int) {
	//Int
	stmt := `DELETE FROM contacts WHERE id = ?`

	_, err := c.DB.Exec(stmt, id)
	if err != nil {
		fmt.Printf("Server Error: %s", err.Error())
		return
	}

	fmt.Printf("Deleted Contact with ID: %s || ", strconv.Itoa(id))
	return
}

func (c *ContactModel) Get(id int) (Contact, error) {
	// SQL statement
	stmt := `SELECT id, firstname, lastname, email, phone, created FROM contacts WHERE id = ?`

	// Create an empty Contact holder
	var ct Contact

	// Query Row on our DB
	err := c.DB.QueryRow(stmt, id).Scan(&ct.ID, &ct.FirstName, &ct.LastName, &ct.Email, &ct.Phone, &ct.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Contact{}, ErrNoRecord
		} else {
			return Contact{}, err
		}
	}

	fmt.Printf("Retrieving Contact with ID: %s || ", strconv.Itoa(id))
	return ct, nil
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
	// that was encountered during iteration.
	// NOTE: Important to always call this. Dont assume that a successful
	// iteration was completed over the whole resultset
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything OK, return the Contacts slice
	return contacts, nil
}
