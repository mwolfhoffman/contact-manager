package src

import (
	"errors"

	"github.com/urfave/cli/v2"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (cs *Service) CheckIfExists(newContact Contact) (bool, error) {
	exists, err := cs.repo.GetUser(newContact)
	if err != nil {
		return true, errors.New("either email or phone are required")
	}
	if (exists == Contact{}) {
		return false, nil
	}
	return false, nil
}

func (cs *Service) AddContact(c *cli.Context) error {

	newContact := Contact{
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

	contactExists, contactExistsError := cs.CheckIfExists(newContact)
	if contactExists {
		return errors.New("contact already exists")
	}
	if contactExistsError != nil {
		return contactExistsError
	}

	cs.repo.AddContact(&newContact)
	return nil
}

func (cs *Service) List(c *cli.Context) ([]Contact, error) {
	return cs.repo.List()
}
