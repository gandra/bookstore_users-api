package users

import (
	"fmt"
	"github.com/gandra/bookstore/usersapi/datasources/mysql/users_db"
	"github.com/gandra/bookstore/usersapi/utils/date_utils"
	"github.com/gandra/bookstore/usersapi/utils/errors"
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
	current := usersDb[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s is already registered", user.Email), "")
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id), "")
	}

	user.DateCreated = date_utils.GetNowString()

	usersDb[user.Id] = user
	return nil
}
