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
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) values (?, ?, ?, ?);"
)

var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDb[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

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
