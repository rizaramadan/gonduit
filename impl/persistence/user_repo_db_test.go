package persistence

import (
	"github.com/rizaramadan/gonduit/users"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"testing"
)

type MockDb struct {
	User *User
}

func (m *MockDb) Create(v interface{}) *gorm.DB {
	user := v.(*User)
	m.User = user
	return &gorm.DB{Error: nil, RowsAffected: 1}
}

func TestUserRepoDb_Create(t *testing.T) {
	r := new(UserRepoDb)
	mock := new(MockDb)
	r.Db = mock

	var password = "some-password"
	err := r.Create(&users.User{
		Username: "",
		Email:    "",
		Password: password,
		Token:    "",
		Bio:      "",
		Image:    "",
	})

	assert.Nil(t, err)

	var passWithSalt = password + mock.User.Salt
	encryptionErr := bcrypt.CompareHashAndPassword([]byte(mock.User.Password), []byte(passWithSalt))

	assert.Nil(t, encryptionErr)
}
