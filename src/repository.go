package src

import (
	"context"
	"fmt"
	"os"

	"gorm.io/gorm"
)

type IRepository interface {
	AddContact(ctx context.Context, contact *Contact)
	List(ctx context.Context) ([]Contact, error)
	GetUser(ctx context.Context, newContact Contact) (Contact, error)
	Search(ctx context.Context, name string, email string, phone string) ([]Contact, error)
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

func (repo *Repository) Search(ctx context.Context, name string, email string, phone string) ([]Contact, error) {
	var result []Contact

	fmt.Println(len(os.Args))

	db, _ := ctx.Value("db").(*gorm.DB)
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
		fmt.Println(whereParts[i])
		if i > 0 {
			where += " AND " + whereParts[i]
		} else {
			where += whereParts[i]
		}
	}

	fmt.Println(where)

	err := db.Where(where).Find(&result).Error
	fmt.Println(err)
	return result, err
}
