package src

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/urfave/cli/v2"
)

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

func (service *Service) AddContact(newContact Contact) error {

	if newContact.Name == "" {
		return errors.New("name is required")
	}

	if newContact.Email == "" && newContact.Phone == "" {
		return errors.New("either email or phone are required")
	}

	if newContact.Email != "" {

		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		isValid := emailRegex.MatchString(newContact.Email)
		if !isValid {
			return errors.New("email must be valid")
		}
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

func (service *Service) Search(c *cli.Context) error {
	searchParams := Contact{
		Name:  c.Value("name").(string),
		Email: c.Value("email").(string),
		Phone: c.Value("phone").(string),
	}

	if len(searchParams.Name) == 0 && len(searchParams.Email) == 0 && len(searchParams.Phone) == 0 {
		return errors.New("you must provide at name, email, and or phone")
	}

	result, err := service.repo.Search(service.ctx, searchParams.Name, searchParams.Email, searchParams.Phone)
	if err != nil {
		return errors.New("error occured")
	}
	fmt.Println(result)
	return nil
}

func (service *Service) Edit(c *cli.Context) error {
	res, err := service.repo.Update(service.ctx, c)
	if err == nil {
		fmt.Printf("Successfully updated %v", res)
	}
	return err
}
