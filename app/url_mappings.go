package app

import (
	"github.com/gandra/bookstore/usersapi/controllers/ping"
	"github.com/gandra/bookstore/usersapi/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// Users
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
