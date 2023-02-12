package commands

import (
	"errors"

	"github.com/mwolfhoffman/contact-manager/models"
	"github.com/mwolfhoffman/contact-manager/repository"
	"github.com/urfave/cli/v2"
)

func checkIfExists(newContact models.Contact) (bool, error) { //	TODO: TEST!!!
	exists, err := repository.GetUser(newContact)
	if err != nil {
		return true, errors.New("either email or phone are required")
	}
	if (exists == models.Contact{}) {
		return false, nil
	}
	return false, nil
}

func AddContact(c *cli.Context) error {

	newContact := models.Contact{
		Name:  c.Value("name").(string),
		Email: c.Value("email").(string),
		Phone: c.Value("phone").(string),
	}

	if newContact.Name == "" {
		return errors.New("name is required")
	}

	if newContact.Email == "" && newContact.Phone == "" {
		return errors.New("either email or phone are required")
	}

	contactExists, contactExistsError := checkIfExists(newContact)
	if contactExists {
		return errors.New("contact already exists")
	}
	if contactExistsError != nil {
		return contactExistsError
	}

	repository.AddContact(&newContact)
	return nil
}

func List(c *cli.Context) ([]models.Contact, error) {
	return repository.List()
}
