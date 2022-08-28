package tests

import (
	"github.com/google/go-cmp/cmp"
	"smallbank/server/datasources"
	setup "smallbank/server/db"
	"smallbank/server/models"
	"testing"
)

func Test_CreateAccount_EmptyResult(t *testing.T) {
	db := setup.SetupTestDB()

	account := models.Account{
		UserID: 1,
	}

	datasources.CreateUser(models.User{}, db)
	got, err := datasources.CreateAccount(account, db)

	want := models.Account{
		Transactions: nil,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	setup.CleanUpTestDB(db)
}
