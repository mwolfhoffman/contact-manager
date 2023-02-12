package commands

import (
	"testing"

	"github.com/mwolfhoffman/contact-manager/commands"
	"github.com/mwolfhoffman/contact-manager/models"
)

//	TODO: returns true if contact does exist, returns error if needed. mocks?

func TestCheckIfExistsReturnsFalseIfContactDoesNotExists(t *testing.T) {
	newContact := models.Contact{Name: "test", Email: "test@test.com"}
	exists, err := commands.CheckIfExists(newContact)
	if err != nil {
		t.Fail()
	}
	t.Errorf("got %d, want %d", exists, false)
}

// https://go.dev/doc/tutorial/add-a-test
