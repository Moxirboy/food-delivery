package models

type Position string

const (
	PositionAdmin   Position = "ADMIN"
	PositionCourier Position = "COURIER"
	PositionUser    Position = "USER"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Position
	At
}
