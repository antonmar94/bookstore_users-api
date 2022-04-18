package app

import (
	"github.com/antonmar94/bookstore_users-api/ping"
	"github.com/antonmar94/bookstore_users-api/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
}
