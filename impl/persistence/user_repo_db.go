package persistence

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/rizaramadan/gonduit/users"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int64
	Username string
	Password string
	Salt     string
	Email    string
	Bio      string
	Image    string
}

type UserRepoDb struct {
	Db *gorm.DB
}

//Cost is required by bcrypt to generate hashed password
const Cost = 10

//NewUserRepoDb return userRepo implementation using dbms
func NewUserRepoDb(db *gorm.DB) users.UserRepo {
	fmt.Println("NewUserRepoDb created")
	r := UserRepoDb{
		Db: db,
	}
	return &r
}

//Create is a command to create user.
func (ur *UserRepoDb) Create(u *users.User) error {
	salt := uuid.NewString()
	var passWithSalt = u.Password + salt
	hashed, err := bcrypt.GenerateFromPassword([]byte(passWithSalt), Cost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	u.Salt = salt

	result := ur.Db.Create(&User{
		Username: u.Username,
		Password: u.Password,
		Salt:     salt,
		Email:    u.Email,
		Bio:      u.Bio,
		Image:    u.Image,
	})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return errors.New("cannot create user")
	}
	return nil
}
