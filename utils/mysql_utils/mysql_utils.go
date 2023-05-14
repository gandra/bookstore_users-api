package mysql_utils

import (
	"github.com/gandra/bookstore/usersapi/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response", err.Error())
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data", sqlErr.Error())
	}
	return errors.NewInternalServerError("error processing request", sqlErr.Error())
}
