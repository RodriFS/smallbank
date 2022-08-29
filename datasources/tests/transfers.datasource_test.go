package tests

import (
	"smallbank/server/datasources"
	setup "smallbank/server/db"
	"smallbank/server/models"
	"testing"
)

func Test_CreateTransfer_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 100}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	datasources.CreateTransfer(models.Transfer{FromAccountId: 1, ToAccountId: 2, Amount: 10}, db)

	transfer, err := datasources.FirstTransfer("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if transfer.FromAccountId != 1 {
		t.Errorf("expected from acount to be 1")
	}

	if transfer.ToAccountId != 2 {
		t.Errorf("expected to acount to be 2")
	}

	if transfer.Amount != 10 {
		t.Errorf("expected amount to be 10")
	}

	setup.CleanUpTestDB(db)
}

func Test_CreateTransfer_UserFromNotFound(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	_, err := datasources.CreateTransfer(models.Transfer{FromAccountId: 100, ToAccountId: 1, Amount: 10}, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "FromAccountId: record not found"
	if err.Error() != message {
		t.Errorf("expected error to be %q, got %q", err.Error(), message)
	}

	setup.CleanUpTestDB(db)
}

func Test_CreateTransfer_UserToNotFound(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 100}, db)

	_, err := datasources.CreateTransfer(models.Transfer{FromAccountId: 1, ToAccountId: 100, Amount: 10}, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "ToAccountId: record not found"
	if err.Error() != message {
		t.Errorf("expected error to be %q, got %q", err.Error(), message)
	}

	setup.CleanUpTestDB(db)
}

func Test_CreateTransfer_UsersCantBeEqual(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 100}, db)

	_, err := datasources.CreateTransfer(models.Transfer{FromAccountId: 1, ToAccountId: 1, Amount: 10}, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "A user can't make a transfer to itself"
	if err.Error() != message {
		t.Errorf("expected error to be %q, got %q", err.Error(), message)
	}

	setup.CleanUpTestDB(db)
}

func Test_CreateTransfer_AmountCantBeZero(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 100}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	_, err := datasources.CreateTransfer(models.Transfer{FromAccountId: 1, ToAccountId: 2, Amount: 0}, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "Can't transfer 0 or less amount"
	if err.Error() != message {
		t.Errorf("expected error to be %q, got %q", err.Error(), message)
	}

	setup.CleanUpTestDB(db)
}

func Test_CreateTransfer_AmountCantBeNegative(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 100}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	_, err := datasources.CreateTransfer(models.Transfer{FromAccountId: 1, ToAccountId: 2, Amount: -10}, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "Can't transfer 0 or less amount"
	if err.Error() != message {
		t.Errorf("expected error to be %q, got %q", err.Error(), message)
	}

	setup.CleanUpTestDB(db)
}

func Test_CreateTransfer_InsufficientFunds(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 0}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	_, err := datasources.CreateTransfer(models.Transfer{FromAccountId: 1, ToAccountId: 2, Amount: 10}, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "FromAccountId: Insufficient funds"
	if err.Error() != message {
		t.Errorf("expected error to be %q, got %q", err.Error(), message)
	}

	setup.CleanUpTestDB(db)
}

func Test_FindTransfers_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 1000}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	datasources.CreateTransfer(models.Transfer{FromAccountId: 1, ToAccountId: 2, Amount: 10}, db)
	datasources.CreateTransfer(models.Transfer{FromAccountId: 2, ToAccountId: 1, Amount: 10}, db)

	transfers, err := datasources.FindTransfers("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if len(transfers) != 2 {
		t.Errorf("expected 2 transfers, got %d", len(transfers))
	}

	setup.CleanUpTestDB(db)
}

func Test_FirstTransfer_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 1000}, db)
	datasources.CreateAccount(models.Account{UserID: 1}, db)

	datasources.CreateTransfer(models.Transfer{FromAccountId: 1, ToAccountId: 2, Amount: 10}, db)

	transfer, err := datasources.FirstTransfer("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if transfer.FromAccountId != 1 {
		t.Errorf("expected transfer with from account 1, got %d", transfer.FromAccountId)
	}

	if transfer.ToAccountId != 2 {
		t.Errorf("expected transfer with from account 1, got %d", transfer.ToAccountId)
	}

	setup.CleanUpTestDB(db)
}

func Test_FirstTransfer_TransferDoesntExistError(t *testing.T) {
	db := setup.SetupTestDB()

	_, err := datasources.FirstTransfer("1", db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "record not found"
	if err.Error() != message {
		t.Errorf("expected error message '%s', got '%s'", message, err.Error())
	}

	setup.CleanUpTestDB(db)
}
