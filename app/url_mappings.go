package app

import (
	"github.com/gandra/bookstore/usersapi/controllers/ping"
	"github.com/gandra/bookstore/usersapi/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// Users
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
