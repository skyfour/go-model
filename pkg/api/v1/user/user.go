package v1

import (
	"github.com/skyfour/go-model/pkg/model"
	"github.com/skyfour/go-model/pkg/service/user"
	"github.com/skyfour/go-model/pkg/utils"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	uid := c.Param("uid")
	utils.Log.Debugln("Code:", uid)

	c.JSON(200,user.GetUser(uid))
}

func InsertUser(c *gin.Context) {
	var user1 model.Users
	c.ShouldBind(&user1)

	c.JSON(200,user.InsertUser(user1))

}

func DeleteUser(c *gin.Context) {
	uid := c.Param("uid")
	utils.Log.Debugln("Code:", uid)
	user.DeleteUser(uid)
	c.JSON(200,"ok")
}
