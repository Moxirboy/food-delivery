package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"food-delivery/pkg/logger"
	"food-delivery/pkg/utils"
)

type employeeRepository struct {
	db  *sql.DB
	log logger.Logger
}

func (e *employeeRepository) CheckField(
	ctx context.Context,
	field, value string,
) (bool, error) {
	var existsClient int

	if field == "position" {
		query := fmt.Sprintf(checkFieldEmployee, "position")
		row := e.db.QueryRowContext(ctx, query, value)
		if err := row.Scan(&existsClient); err != nil {
			e.log.Error(
				"Error while checking field 'position' of employee",
				err.Error(),
			)
			return false, err
		}
	} else {
		return false, utils.ErrInvalidField
	}

	if existsClient > 0 {
		return false, nil
	}

	return true, nil
}

func (e *employeeRepository) CreateEmployee(
	ctx context.Context,
	employee *models.Employee,
) (string, error) {
	var (
		employeeID string
	)

	row := e.db.QueryRowContext(
		ctx,
		createEmployee,
		employee.FirstName,
		employee.LastName,
		employee.PhoneNumber,
		employee.Position,
		employee.Password,
	)

	if err := row.Scan(&employeeID); err != nil {
		e.log.Error(
			"Error while inserting into employee_details",
			err.Error(),
		)
		return "", err
	}

	return employeeID, nil
}

func (e *employeeRepository) Login(
	ctx context.Context,
	login, password string,
) (*models.Employee, error) {
	var (
		ID        string
		position  string
		firstName string
		lastName  sql.NullString
		code      string
	)

	row := e.db.QueryRowContext(
		ctx,
		getEmployeeByAuthCred,
		login,
		password,
	)

	if err := row.Scan(
		&ID,
		&firstName,
		&lastName,
		&position,
		&code,
	); err != nil {
		if err != sql.ErrNoRows {
			e.log.Error(
				"Error while selecting from employee_details to login",
				err.Error(),
			)
		}

		return nil, err
	}

	return &models.Employee{
		ID:          ID,
		FirstName:   firstName,
		LastName:    lastName.String,
		Code:        code,
		Position:    position,
		PhoneNumber: login,
	}, nil
}
