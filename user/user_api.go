package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetById(c *gin.Context) {
	id, isGetId := c.GetQuery("id")
	fmt.Printf("GetById:id = %s", id)
	if isGetId {
		dao := new(Dao)
		user, _ := dao.GetUser(id)
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, nil)
	}

}
