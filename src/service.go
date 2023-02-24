package src

import (
	"context"
	"errors"

	"github.com/urfave/cli/v2"
)

type IService interface {
}

type Service struct {
	ctx  context.Context
	repo IRepository
}

func NewService(ctx context.Context, repo IRepository) *Service {
	return &Service{
		ctx:  ctx,
		repo: repo,
	}
}

func (service *Service) CheckIfExists(newContact Contact) (bool, error) {
	exists, err := service.repo.GetUser(service.ctx, newContact)
	if err != nil {
		return true, errors.New("either email or phone are required")
	}
	if (exists == Contact{}) {
		return false, nil
	}
	return true, nil
}

func (service *Service) AddContact(c *cli.Context) error {

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

	contactExists, contactExistsError := service.CheckIfExists(newContact)
	if contactExists {
		return errors.New("contact already exists")
	}
	if contactExistsError != nil {
		return contactExistsError
	}

	service.repo.AddContact(service.ctx, &newContact)
	return nil
}

func (service *Service) List(c *cli.Context) ([]Contact, error) {
	return service.repo.List(service.ctx)
}
