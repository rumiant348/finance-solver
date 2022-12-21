package models

import "github.com/jinzhu/gorm"

// UserDB interface

// UserDB is used to interact with the users database
//
// For pretty much all single users queries
// If the user is found, we will return a nil error
// If the user is not found, we will return ErrNotFound
// If there is another error, we will return an error with
// more information about what went wrong. This may not be
// an error generated by the models package.
//
// For single user queries, any error but ErrNotFound should
// probably result in a 500 error until we make "public"
// facing errors
type UserDB interface {
	// ByID ByEmail ByRemember - methods for querying a single user
	ByID(in uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	// Create Update Delete - methods for altering users
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

// User type
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}