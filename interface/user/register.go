package user

import (
	"dance/conf"
	"dance/cons"
	"dance/core"
	"dance/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"time"
)

type argsRegister struct {
	Phone  string `json:"phone" binding:"required,max=11"`
	Name   string `json:"name" binding:"required,max=64"`
	Sex    int    `json:"sex" binding:"required,enum=1-2"`
	Avatar string `json:"avatar"`
	Bk     string `json:"bk"`
}

func handleRegister(c *core.Context) {
	var (
		args = c.Keys["args"].(*argsRegister)
		now  = time.Now().Format(cons.FORMAT_TIME)
	)
	user := model.UserInfo{
		UserId: args.Phone,
		Token:  "", // TODO
		Name:   args.Name,
		Sex:    args.Sex,
		Bk:     args.Bk,

		CreateTime: now,
		UpdateTime: now,
	}
	if err := core.GetDB().Table("user").Create(&user).Error; err != nil {
		conf.MainLog.Errorf("register %v failed. err:%v", args.Phone, err.Error())
		c.JSON(http.StatusOK, cons.ERR_PUB_SYSTEM, "注册用户出错", nil)
		return
	}
	c.JSON(http.StatusOK, 0, "", gin.H{"user_id": args.Phone})
}

func init() {
	checks := []core.FunCheck{}
	finish := []core.FunHandle{}
	core.Engine.POST(
		"/dance/user/register",
		core.HandleRequest(
			reflect.TypeOf(argsRegister{}),
			handleRegister, checks, finish,
		),
	)
}
