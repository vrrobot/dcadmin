package api

import (
	"dcadmin/dao"
	"dcadmin/models"
	"dcadmin/response"
	"log"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		response.Failed("参数错误", c)
		return
	}

	dao.Mgr.AddUser(&user)
	response.Success("添加用户成功", user, c)
}

func ListUser(c *gin.Context) {
	users := dao.Mgr.ListUser()
	response.Success("获取用户列表成功", users, c)
}

func VerifyUser(c *gin.Context) {
	username := c.Query("loginId")
	password := c.Query("loginPwd")
	log.Printf(username)
	log.Printf(password)
	user := dao.Mgr.GetUser(username)

	if user.Password == password {
		response.Success("用户验证成功", "", c)
	} else {
		response.Failed("密码错误", c)
	}
}
