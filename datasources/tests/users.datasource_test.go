package tests

import (
	"smallbank/server/datasources"
	setup "smallbank/server/db"
	"smallbank/server/models"
	"testing"
)

func Test_CreateUser_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	user := models.User{
		Name: "Patrick",
		Last: "Swayze",
		Phone: models.Phone{
			Code:   34,
			Number: 533986781,
		},
	}

	_, err := datasources.CreateUser(user, db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	var createdUser *models.User
	result := db.First(&createdUser, 1)

	if result.Error != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if createdUser.Name != user.Name {
		t.Errorf("expected name %s, got %s", user.Name, createdUser.Name)
	}

	if createdUser.Last != user.Last {
		t.Errorf("expected last name %s, got %s", user.Last, createdUser.Last)
	}

	if createdUser.Phone.Code != user.Phone.Code {
		t.Errorf("expected phone code %d, got %d", user.Phone.Code, createdUser.Phone.Code)
	}

	if createdUser.Phone.Number != user.Phone.Number {
		t.Errorf("expected phone number %d, got %d", user.Phone.Number, createdUser.Phone.Number)
	}

	setup.CleanUpTestDB(db)
}

func Test_FindUser_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)
	datasources.CreateUser(models.User{}, db)

	users, err := datasources.FindUsers(db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if len(users) != 2 {
		t.Errorf("expected 2 userss, got %d", len(users))
	}

	setup.CleanUpTestDB(db)
}

func Test_FirstUser_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)

	user, err := datasources.FirstUser("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if user.ID != 1 {
		t.Errorf("expected  user with id 1, got %d", user.ID)
	}

	setup.CleanUpTestDB(db)
}

func Test_FirstUser_UserDoesntExistError(t *testing.T) {
	db := setup.SetupTestDB()

	_, err := datasources.FirstUser("1", db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "record not found"
	if err.Error() != message {
		t.Errorf("expected error message '%s', got '%s'", message, err.Error())
	}

	setup.CleanUpTestDB(db)
}

func Test_UpdateUser_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)

	user := map[string]any{
		"Name":   "Sam",
		"Last":   "Rockwell",
		"Code":   12,
		"Number": int32(729749264),
	}
	err := datasources.UpdateUser("1", user, db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	updatedUser, err := datasources.FirstUser("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if updatedUser.Name != user["Name"] {
		t.Errorf("expected name to be %s, got %s", user["Name"], updatedUser.Name)
	}

	if updatedUser.Last != user["Last"] {
		t.Errorf("expected last name to be %s, got %s", user["Last"], updatedUser.Last)
	}

	if updatedUser.Phone.Code != user["Code"] {
		t.Errorf("expected phone code to be %d, got %d", user["Code"], updatedUser.Phone.Code)
	}

	if updatedUser.Phone.Number != user["Number"] {
		t.Errorf("expected phone number to be %d, got %d", user["Number"], updatedUser.Phone.Number)
	}

	setup.CleanUpTestDB(db)
}

func Test_UpdateUser_UserDoesntExistError(t *testing.T) {
	db := setup.SetupTestDB()

	user := map[string]any{
		"Name": "Sam",
	}
	err := datasources.UpdateUser("1", user, db)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	message := "record not found"
	if err.Error() != message {
		t.Errorf("expected error message '%s', got '%s'", message, err.Error())
	}

	setup.CleanUpTestDB(db)
}

func Test_DeleteUser_NoError(t *testing.T) {
	db := setup.SetupTestDB()

	datasources.CreateUser(models.User{}, db)

	users, err := datasources.FindUsers(db)
	if len(users) != 1 || err != nil {
		t.Errorf("Error setting up test")
	}

	err = datasources.DeleteUser("1", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	users, err = datasources.FindUsers(db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	if len(users) != 0 {
		t.Errorf("expected users to be deleted, got %d users(s)", len(users))
	}

	setup.CleanUpTestDB(db)
}

func Test_DeleteUser_UserDoesntExistNoErrorEither(t *testing.T) {
	db := setup.SetupTestDB()

	err := datasources.DeleteUser("100", db)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}

	setup.CleanUpTestDB(db)
}
