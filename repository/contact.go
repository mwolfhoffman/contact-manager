package repository

import (
	"github.com/mwolfhoffman/contact-manager/db"
	"github.com/mwolfhoffman/contact-manager/models"
)

func AddContact(contact *models.Contact) {
	db.DB.Create(&models.Contact{Name: contact.Name, Email: contact.Email, Phone: contact.Phone})
}

func List() ([]models.Contact, error) {
	var result []models.Contact
	err := db.DB.Exec("select * from contacts").Find(&result).Error //  TODO: get all, then add ability to search.
	return result, err
}

func GetUser(newContact models.Contact) (models.Contact, error) {
	var result models.Contact
	err := db.DB.Where(&models.Contact{Name: newContact.Name, Phone: newContact.Phone, Email: newContact.Email}).Find(&result).Error
	return result, err
}
