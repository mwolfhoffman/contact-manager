package src

import (
	"context"

	"gorm.io/gorm"
)

type IRepository interface {
	AddContact(ctx context.Context, contact *Contact)
	List(ctx context.Context) ([]Contact, error)
	GetUser(ctx context.Context, newContact Contact) (Contact, error)
}

type Repository struct {
	repo IRepository
}

func NewRepository(ctx context.Context) *Repository {
	return &Repository{}
}

func (repo *Repository) AddContact(ctx context.Context, contact *Contact) {
	db, _ := ctx.Value("db").(*gorm.DB)
	db.Create(&Contact{Name: contact.Name, Email: contact.Email, Phone: contact.Phone})
}

func (repo *Repository) List(ctx context.Context) ([]Contact, error) {
	var result []Contact
	db, _ := ctx.Value("db").(*gorm.DB)
	err := db.Exec("select * from contacts").Find(&result).Error //  TODO: get all, then add ability to search.
	return result, err
}

func (repo *Repository) GetUser(ctx context.Context, newContact Contact) (Contact, error) {
	var result Contact
	db, _ := ctx.Value("db").(*gorm.DB)
	err := db.Where(&Contact{Name: newContact.Name, Phone: newContact.Phone, Email: newContact.Email}).Find(&result).Error
	return result, err
}
