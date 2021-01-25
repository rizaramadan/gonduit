package persistence

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rizaramadan/gonduit/users"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func TestUserRepoDb_Create(t *testing.T) {
	r := new(UserRepoDb)

	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	r.Db = gormDB
	var password = "some-password"

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("username","password","salt","email","bio","image") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))
	mock.ExpectCommit()

	user := &users.User{
		Username: "",
		Email:    "",
		Password: password,
		Token:    "",
		Bio:      "",
		Image:    "",
	}

	err = r.Create(user)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	var passWithSalt = password + user.Salt
	encryptionErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passWithSalt))

	assert.Nil(t, encryptionErr)
}
