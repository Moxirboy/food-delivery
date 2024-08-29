package models

type Employee struct {
	ID        string
	FirstName string `validate:"required"`
	LastName  string
	Email     string `validate:"required,email"`
	BirthDate string `validate:"required"`
	Position  string `validate:"required"`
	Password  string

	At
	By
}
