package users

import (
	"fmt"
	"github.com/gandra/bookstore/usersapi/datasources/mysql/users_db"
	"github.com/gandra/bookstore/usersapi/utils/date_utils"
	"github.com/gandra/bookstore/usersapi/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "UK_EMAIL"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) values (?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users where id = ?;"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError("Error preparing get user statement", err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to get user for id %d", user.Id), err.Error())
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("Error preparing insert statement", err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email), "")
		}
		return errors.NewInternalServerError("Error inserting user", err.Error())
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("Error when trying to get last insert id", err.Error())
	}

	user.Id = userId
	return nil
}
