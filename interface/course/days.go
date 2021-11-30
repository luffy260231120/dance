package course

import (
	"dance/core"
	"reflect"
)

type argsDays struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func handleRegister(c *core.Context) {

	//user := model.UserInfo{
	//	UserId:     args.Phone,
	//	Token:      "", // TODO
	//	Name:       args.Name,
	//	Sex:        args.Sex,
	//	Bk:         args.Bk,
	//	CreateTime: now,
	//	UpdateTime: now,
	//}
	//if err := core.GetDB().Table("user").Create(&user).Error; err != nil {
	//	conf.MainLog.Errorf("register %v failed. err:%v", args.Phone, err.Error())
	//	c.JSON(http.StatusOK, cons.ERR_PUB_SYSTEM, "注册用户出错", nil)
	//	return
	//}
	//c.JSON(http.StatusOK, 0, "", gin.H{"user_id": args.Phone})
}

func init() {
	checks := []core.FunCheck{}
	finish := []core.FunHandle{}
	core.Engine.GET(
		"/dance/course/days",
		core.HandlePost(
			reflect.TypeOf(argsDays{}),
			handleRegister, checks, finish,
		),
	)
}
