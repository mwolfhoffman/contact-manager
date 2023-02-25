package main

import (
	"testing"
)

func TestAddContact(t *testing.T) {
	repo, ctx := SetupRepo()
	want := &Contact{Name: "test123", Email: "test@test.com", Phone: "555-555-5555"}
	repo.AddContact(ctx, want)

	got, err := repo.Search(ctx, "test123", "test@test.com", "555-555-5555")

	if err != nil {
		t.Fatal("Error occurred")
	}

	if len(got) != 1 {
		t.Fatal("failed")
	}

	if got[0].Name != want.Name || got[0].Email != want.Email || got[0].Phone != want.Phone {
		t.Fatalf("expected %v but got %v", want, got)
	}

	Teardown()
}

func TestSearch(t *testing.T) {
	repo, ctx := SetupRepo()
	got, err := repo.Search(ctx, "test", "", "")

	if err != nil {
		t.Fatal("Error occurred")
	}

	if len(got) > 0 {
		t.Fatal("failed")
	}
	Teardown()
}

func TestSearchIfContactExists(t *testing.T) {
	newContact := &Contact{Name: "test", Email: "test@test.com"}
	repo, ctx := SetupRepo()
	repo.AddContact(ctx, newContact)

	got, err := repo.Search(ctx, "test", "", "")

	if err != nil {
		t.Fatal("Error occurred")
	}

	if len(got) != 1 {
		t.Fatal("failed")
	}
	Teardown()
}
