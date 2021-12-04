package user

import (
	"dance/conf"
	"dance/cons"
	"dance/core"
	"dance/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"reflect"
)

const Password = "123456"

type argsLogin struct {
	Phone    string `json:"phone" binding:"required,max=11"`
	Password string `json:"password" binding:"required,max=11"`
}

func handleLogin(c *core.Context) {
	var (
		args = c.Keys["args"].(*argsLogin)
		user = &model.UserInfo{}
	)

	if err := core.GetDB().Table("user").Where("user_id=?", args.Phone).First(user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusOK, cons.ERR_PUB_PARAMS, "登录失败", nil)
			return
		}
		conf.MainLog.Errorf("登录未知错误,err:%v", err.Error())
		return
	}

	c.JSON(http.StatusOK, 0, "", gin.H{
		"userId": user.UserId,
		"token":  user.Token,
		"name":   user.Name,
		"sex":    user.Sex,
	})
}

func checkPassword(c *core.Context) bool {
	var (
		args = c.Keys["args"].(*argsLogin)
	)

	if args.Password == Password {
		return true
	}
	c.JSON(http.StatusOK, cons.ERR_PUB_PARAMS, "登录失败 密码错误", nil)
	return false
}

func init() {
	checks := []core.FunCheck{core.CheckSignMD5, checkPassword}
	finish := []core.FunHandle{}
	core.Engine.POST(
		"/dance/user/login",
		core.HandleRequest(
			reflect.TypeOf(argsLogin{}),
			handleLogin, checks, finish,
		),
	)
}
