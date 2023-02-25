package src

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type IRepository interface {
	AddContact(ctx context.Context, contact *Contact)
	List(ctx context.Context) ([]Contact, error)
	GetUser(ctx context.Context, newContact Contact) (Contact, error)
	Search(ctx context.Context, name string, email string, phone string) ([]Contact, error)
	Update(ctx context.Context, c *cli.Context) (Contact, error)
}

type Repository struct {
	repo IRepository
}

func NewRepository() *Repository {
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

func (repo *Repository) Search(ctx context.Context, name string, email string, phone string) ([]Contact, error) {
	var result []Contact
	db, ok := ctx.Value("db").(*gorm.DB)
	if ok == false {
		fmt.Printf("error getting db from context: %v", db)
	}
	var where string
	var whereParts []string

	if len(name) > 0 {
		whereParts = append(whereParts, "name = '"+name+"'")
	}
	if len(email) > 0 {
		whereParts = append(whereParts, "email = '"+email+"'")
	}
	if len(phone) > 0 {
		whereParts = append(whereParts, "phone = '"+phone+"'")
	}

	for i := 0; i < len(whereParts); i++ {
		if i > 0 {
			where += " AND " + whereParts[i]
		} else {
			where += whereParts[i]
		}
	}

	err := db.Where(where).Find(&result).Error
	return result, err
}

func (repo *Repository) Update(ctx context.Context, c *cli.Context) (Contact, error) {
	var result Contact
	id := c.Value("id")
	contact := Contact{
		Name:  c.Value("name").(string),
		Email: c.Value("email").(string),
		Phone: c.Value("phone").(string),
	}
	db, _ := ctx.Value("db").(*gorm.DB)
	db.Model(Contact{}).Where("id = ?", id).Updates(Contact{Name: contact.Name, Email: contact.Email, Phone: contact.Phone})
	err := db.Where("id = ?", id).Find(&result).Error
	return result, err
}
