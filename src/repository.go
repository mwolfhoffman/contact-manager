package src

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) AddContact(contact *Contact) {
	repo.db.Create(&Contact{Name: contact.Name, Email: contact.Email, Phone: contact.Phone})
}

func (repo *Repository) List() ([]Contact, error) {
	var result []Contact
	err := repo.db.Exec("select * from contacts").Find(&result).Error //  TODO: get all, then add ability to search.
	return result, err
}

func (repo *Repository) GetUser(newContact Contact) (Contact, error) {
	var result Contact
	err := repo.db.Where(&Contact{Name: newContact.Name, Phone: newContact.Phone, Email: newContact.Email}).Find(&result).Error
	return result, err
}
