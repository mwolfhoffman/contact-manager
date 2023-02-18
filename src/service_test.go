package src

import (
	"testing"
)

func TestCheckIfExistsReturnsFalseIfUserNotExists(t *testing.T) {
	db := ConnectToDb()
	repo := NewRepository(db)
	service := NewService(repo)

	newContact := &Contact{Name: "test", Email: "test@test.com"}
	got, err := service.CheckIfExists(*newContact)
	want := false

	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	if got != false {
		t.Fatalf("expected %t but got %t", want, got)
	}
}

//	TODO: returns true if exists, check for errors
//	TODO: add contact needs correct args.
