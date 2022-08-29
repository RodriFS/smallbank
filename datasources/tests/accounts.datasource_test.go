package tests

import (
	"smallbank/server/datasources"
	setup "smallbank/server/db"
	"smallbank/server/models"
	"testing"
)

func Test_CreateAccount_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	account := models.Account{
		UserID:   1,
		Balance:  10,
		Currency: "USD",
	}

	datasources.CreateUser(models.User{}, db)
	_, err := datasources.CreateAccount(account, db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	var createdAccount *models.Account
	result := db.First(&createdAccount, 1)

	if result.Error != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if createdAccount.UserID != account.UserID {
		t.Errorf("expected user %d, got %d", account.UserID, createdAccount.UserID)
	}

	if createdAccount.Balance != account.Balance {
		t.Errorf("expected balance %d, got %d", account.Balance, createdAccount.Balance)
	}

	if createdAccount.Currency != account.Currency {
		t.Errorf("expected currency %s, got %s", account.Currency, createdAccount.Currency)
	}

	if !createdAccount.Active {
		t.Errorf("expected account to be active")
	}

	setup.CleanUpTestDB(db)
}

func Test_CreateAccount_UserDoesntExistError(t *testing.T) {
	db := setup.SetupTestDB()

	account := models.Account{
		UserID:   1,
		Balance:  10,
		Currency: "USD",
	}

	_, err := datasources.CreateAccount(account, db)

	if err == nil {
		t.Errorf("expected error, got %q", err)
	}

	message := "User: record not found"
	if err.Error() != message {
		t.Errorf("expected error message '%s', got '%s'", message, err.Error())
	}

	setup.CleanUpTestDB(db)
}

func Test_FindAccounts_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	accounts, err := datasources.FindAccounts(db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if len(accounts) != 2 {
		t.Errorf("expected 2 accounts, got %d", len(accounts))
	}

	setup.CleanUpTestDB(db)
}

func Test_FirstAccount_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	account, err := datasources.FirstAccount("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if account.ID != 1 {
		t.Errorf("expected account with id 1, got %d", account.ID)
	}

	setup.CleanUpTestDB(db)
}

func Test_FirstAccount_AccountDoesntExistError(t *testing.T) {
	db := setup.SetupTestDB()

	_, err := datasources.FirstAccount("1", db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "record not found"
	if err.Error() != message {
		t.Errorf("expected error message '%s', got '%s'", message, err.Error())
	}

	setup.CleanUpTestDB(db)
}

func Test_UpdateAccount_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	account := map[string]any{
		"Balance":  int64(10),
		"Currency": "USD",
		"Active":   false,
	}
	err := datasources.UpdateAccount("1", account, db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	updatedAccount, err := datasources.FirstAccount("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if updatedAccount.Balance != account["Balance"] {
		t.Errorf("expected balance to be %d, got %d", account["Balance"], updatedAccount.Balance)
	}

	if updatedAccount.Currency != account["Currency"] {
		t.Errorf("expected balance to be %q, got %q", account["Currency"], updatedAccount.Currency)
	}

	if updatedAccount.Active != account["Active"] {
		t.Errorf("expected balance to be %t, got %t", account["Active"], updatedAccount.Active)
	}

	setup.CleanUpTestDB(db)
}

func Test_UpdateAccount_AccountDoesntExistError(t *testing.T) {
	db := setup.SetupTestDB()

	account := map[string]any{
		"Balance":  int64(10),
		"Currency": "USD",
		"Active":   false,
	}
	err := datasources.UpdateAccount("1", account, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "record not found"
	if err.Error() != message {
		t.Errorf("expected error message '%s', got '%s'", message, err.Error())
	}

	setup.CleanUpTestDB(db)
}

func Test_DeleteAccount_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	accounts, err := datasources.FindAccounts(db)
	if len(accounts) != 1 || err != nil {
		t.Errorf("Error setting up test")
	}

	err = datasources.DeleteAccount("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	accounts, err = datasources.FindAccounts(db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if len(accounts) != 0 {
		t.Errorf("expected account to be deleted, got %d account(s)", len(accounts))
	}

	setup.CleanUpTestDB(db)
}

func Test_DeleteAccount_AccountDoesntExistNoErrorEither(t *testing.T) {
	db := setup.SetupTestDB()

	err := datasources.DeleteAccount("100", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	setup.CleanUpTestDB(db)
}
