package main

import (
	"testing"
)

func TestAddContactWithValidFields(t *testing.T) {
	service := SetupService()

	contact := &Contact{
		Name:  "test",
		Email: "test@test.com",
		Phone: "555-555-5555",
	}

	err := service.AddContact(*contact)

	if err != nil {
		t.Fatal("Could not add a valid contact")
	}
	Teardown()
}

func TestAddContactRequiresValidEmail(t *testing.T) {
	service := SetupService()

	contact := &Contact{
		Name:  "test",
		Email: "test",
		Phone: "555-555-5555",
	}

	err := service.AddContact(*contact)
	if err.Error() != "email must be valid" {
		t.Fatal("Should not have allowed invalid email")
	}

	Teardown()
}

func TestAddContactIfContactExists(t *testing.T) {
	service := SetupService()

	contact := &Contact{
		Name:  "test",
		Email: "test@test.com",
		Phone: "555-555-5555",
	}

	service.AddContact(*contact)

	err := service.AddContact(*contact)

	if err.Error() != "contact already exists" {
		t.Fatal("Should not have allowed duplicate entries")
	}
	Teardown()
}
