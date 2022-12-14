package models

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

var (
	// ErrNotFound is returned when a resource cannot be found
	// in the database.
	ErrNotFound modelError = "models: resource not found"

	// ErrIDInvalid is returned when an invalid id is passed to methods
	// like Delete.
	ErrIDInvalid modelError = "models: ID provided was invalid"

	// ErrPasswordIncorrect is returned when an invalid password
	// is used when attempting to authenticate a user.
	ErrPasswordIncorrect modelError = "models: incorrect password provided"

	// ErrPasswordTooShort is returned when a user tries to set
	// a password that is less than 8 characters long
	ErrPasswordTooShort modelError = "models: password must be at least 8 characters long"

	// ErrPasswordRequired is returned when a create is attempted
	// without a user password provided.
	ErrPasswordRequired modelError = "models: password is required"

	// ErrEmailRequired is returned when an email address is
	// not provided when creating a user
	ErrEmailRequired modelError = "models: email address is required"

	// ErrEmailInvalid is returned when an email address provided
	// does not match any of our requirements
	ErrEmailInvalid modelError = "models: email address is not valid"

	// ErrEmailTaken is returned when an update or create is attempted
	// with an email address what is already in use
	ErrEmailTaken modelError = "models: email address is already taken"

	// ErrRememberRequired is returned when a create or update
	// is attempted without a user remember token hash
	ErrRememberRequired modelError = "models: remember token is required"

	// ErrRememberTooShort is returned when a remember token is
	// not at least 32 bytes
	ErrRememberTooShort modelError = "models: remember token is not at least 32 bytes"

	// ErrUserIDRequired is returned for list resource if the user id is absent
	ErrUserIDRequired modelError = "models: user ID is required"

	// ErrTitleEmpty is returned when a create is attempted with empty title
	ErrTitleEmpty modelError = "models: title should not be empty"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	split := strings.Split(s, " ")
	split[0] = cases.Title(language.English, cases.Compact).String(split[0])
	return strings.Join(split, " ")
}

func Title(s string) string {
	return cases.Title(language.English, cases.Compact).String(s)
}

type PublicError interface {
	error
	Public() string
}
