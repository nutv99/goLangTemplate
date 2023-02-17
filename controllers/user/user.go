package user

import (
	"melivecode/jwt-api/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

var hmacSampleSecret []byte

func ReadAll(c *gin.Context) {

	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "ok",
		"mesage": "Read All user Data  Success",
		"users":  users,
	})
}

func Profile(c *gin.Context) {

	userId := c.MustGet("userId").(float64)
	// userId := 1
	var user orm.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusBadRequest, gin.H{
		"status":        "ok",
		"mesage":        "Read  user Data  Success",
		"users Profile": user,
	})
}
