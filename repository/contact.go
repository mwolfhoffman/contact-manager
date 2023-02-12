package repository

import (
	"database/sql"

	"github.com/mwolfhoffman/contact-manager/db"
	"github.com/mwolfhoffman/contact-manager/models"
)

func AddContact(contact *models.Contact) {
	db.DB.Create(&models.Contact{Name: contact.Name, Email: contact.Email, Phone: contact.Phone})
}

func List() (*sql.Rows, error) {
	res := db.DB.Find(models.Contact{})
	rows, err := res.Rows()
	return rows, err
}

func GetUser(newContact models.Contact) (*sql.Rows, error) {
	res := db.DB.Where(&models.Contact{Name: newContact.Name, Phone: newContact.Phone, Email: newContact.Email})
	rows, err := res.Rows()
	return rows, err
}
