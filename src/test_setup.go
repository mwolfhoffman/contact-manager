package src

import (
	"context"
	"os"

	"gorm.io/gorm"
)

func SetupRepo() (IRepository, context.Context) {
	ctx := context.Background()
	db := ConnectToDb(os.Getenv("TEST_DB_CONN_STRING"))
	ctx = context.WithValue(ctx, "db", db)
	repo := NewRepository()
	return repo, ctx
}

func SetupService() Service {
	ctx := context.Background()
	db := ConnectToDb(os.Getenv("TEST_DB_CONN_STRING"))
	ctx = context.WithValue(ctx, "db", db)
	repo := NewRepository()
	service := NewService(ctx, repo)
	return *service
}

func Teardown() {
	_, ctx := SetupRepo()
	db, _ := ctx.Value("db").(*gorm.DB)
	db.Exec("delete from contacts")
}
