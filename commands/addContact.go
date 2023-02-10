package commands

import (
	"errors"
	"fmt"

	"github.com/mwolfhoffman/contact-manager/models"
	"github.com/urfave/cli/v2"
)

var contacts []models.Contact

func checkIfExists(newContact models.Contact) bool {
	for i := 0; i < len(contacts); i++ {
		if contacts[i].Name == newContact.Name && contacts[i].Email == newContact.Email && contacts[i].Phone == newContact.Phone {
			return true
		}
	}
	return false
}

func AddContact(c *cli.Context) error {

	newContact := models.Contact{
		Name:  c.Value("name").(string),
		Email: c.Value("email").(string),
		Phone: c.Value("phone").(string),
	}

	if newContact.Name == "" {
		return errors.New("name is required.")
	}

	if newContact.Email == "" && newContact.Phone == "" {
		return errors.New("Either email or phone are required.")
	}

	fmt.Println(newContact)

	if checkIfExists(newContact) == true {
		return errors.New("Contact already exists")
	}

	list := append(contacts, newContact)
	print(list)
	fmt.Println("Contact added to list successfully.")
	return nil
}
