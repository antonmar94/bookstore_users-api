package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implementame")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implementame")

}

/* func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implementame")
} */
