package tests

import (
	"smallbank/server/datasources"
	setup "smallbank/server/db"
	"smallbank/server/models"
	"testing"
)

func Test_CreateTransaction_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 100}, db)

	datasources.CreateTransaction(models.Transaction{AccountID: 1, Amount: 10}, db)

	transaction, err := datasources.FirstTransaction("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if transaction.AccountID != 1 {
		t.Errorf("expected to acount to be 1")
	}

	if transaction.Amount != 10 {
		t.Errorf("expected amount to be 10")
	}

	setup.CleanUpTestDB(db)
}

func Test_CreateTransaction_AccountNotFound(t *testing.T) {
	db := setup.SetupTestDB()

	_, err := datasources.CreateTransaction(models.Transaction{AccountID: 1, Amount: 10}, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "AccountID: record not found"
	if err.Error() != message {
		t.Errorf("expected error to be %q, got %q", err.Error(), message)
	}

	setup.CleanUpTestDB(db)
}

func Test_CreateTransaction_AmountCantBeZero(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 100}, db)

	_, err := datasources.CreateTransaction(models.Transaction{AccountID: 1, Amount: 0}, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "Can't do a transaction on 0 value"
	if err.Error() != message {
		t.Errorf("expected error to be %q, got %q", err.Error(), message)
	}

	setup.CleanUpTestDB(db)
}

func Test_FindTransactions_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 1000}, db)

	datasources.CreateTransaction(models.Transaction{AccountID: 1, Amount: 10}, db)
	datasources.CreateTransaction(models.Transaction{AccountID: 1, Amount: -10}, db)

	transactions, err := datasources.FindTransactionsByAccountId("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if len(transactions) != 2 {
		t.Errorf("expected 2 transactions, got %d", len(transactions))
	}

	setup.CleanUpTestDB(db)
}

func Test_FirstTransactio_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateAccount(models.Account{UserID: 1, Balance: 1000}, db)

	datasources.CreateTransaction(models.Transaction{AccountID: 1, Amount: 10}, db)

	transaction, err := datasources.FirstTransaction("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if transaction.AccountID != 1 {
		t.Errorf("expected transaction with from account 1, got %d", transaction.AccountID)
	}

	setup.CleanUpTestDB(db)
}

func Test_FirstTransaction_TransactionDoesntExistError(t *testing.T) {
	db := setup.SetupTestDB()

	_, err := datasources.FirstTransaction("1", db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "record not found"
	if err.Error() != message {
		t.Errorf("expected error message '%s', got '%s'", message, err.Error())
	}

	setup.CleanUpTestDB(db)
}
