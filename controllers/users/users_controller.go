package users

import (
	"fmt"
	"github.com/gandra/bookstore/usersapi/domain/users"
	"github.com/gandra/bookstore/usersapi/services"
	"github.com/gandra/bookstore/usersapi/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body", err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userIdEntry := c.Param("user_id")
	userId, userErr := strconv.ParseInt(userIdEntry, 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("invalid user id", fmt.Sprintf("user id not an number: %s", userIdEntry))
		c.JSON(restErr.Status, restErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userIdEntry := c.Param("user_id")
	userId, userErr := strconv.ParseInt(userIdEntry, 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("invalid user id", fmt.Sprintf("user id not an number: %s", userIdEntry))
		c.JSON(restErr.Status, restErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body", err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
