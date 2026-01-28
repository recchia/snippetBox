package mock

import (
	"time"

	"github.com/recchia/snippetbox/internal/models"
	"github.com/recchia/snippetbox/internal/models/mysql"
)

type UserModel struct{}

var mockUser = mysql.User{
	ID:      1,
	Name:    "Alice",
	Email:   "alice@example.com",
	Created: time.Now(),
}

func (m *UserModel) Get(id int) (mysql.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return mysql.User{}, models.ErrNoRecord
	}
}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
	if id == 1 {
		if currentPassword != "pa$$word" {
			return models.ErrInvalidCredentials
		}

		return nil
	}

	return models.ErrNoRecord
}
