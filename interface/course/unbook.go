package course

import (
	"dance/conf"
	"dance/cons"
	"dance/core"
	"dance/model"
	"fmt"
	"net/http"
	"reflect"
)

type argsUnBook struct {
	model.UserAuth
	CourseId  int `json:"course_id" binding:"required"`
	MaxNumber int `json:"max_number" binding:"required"` // TODO 最大人数这里应该做一个预约池，以后需要优化
}

func handleUnBook(c *core.Context) {
	var (
		args           = c.Keys["args"].(*argsUnBook)
		key            = fmt.Sprintf(cons.CacheKeyBookUserIds, args.CourseId)
		userIds        = core.GetRedis().ZRange(key, 0, -1).Val()
		needChangeCard = false
		IsWaited       = false
	)

	index := 0
	for _, userId := range userIds {
		index++
		if userId != args.UserID {
			continue
		}
		IsWaited = true

		err := core.GetRedis().ZRem(key, userId).Err()
		if err != nil {
			conf.MainLog.Errorf(err.Error())
			c.JSON(http.StatusOK, cons.ERR_PUB_SYSTEM, "取消失败", nil)
			return
		}

		if index <= args.MaxNumber {
			needChangeCard = true
		}
		break
	}
	if !IsWaited {
		c.JSON(http.StatusOK, cons.ERR_PUB_PARAMS, "已取消 请刷新", nil)
		return
	}

	if needChangeCard {
		// TODO userid 卡次数增加

		if len(userIds) > args.MaxNumber {
			// TODO userIds[args.MaxNumber-1] 这个人扣次数
		}
	}

	c.JSON(http.StatusOK, 0, "取消成功", nil)
}

func init() {
	checks := []core.FunCheck{core.CheckUserId, core.CheckToken}
	finish := []core.FunHandle{}
	core.Engine.POST(
		"/dance/course/unbook",
		core.HandleRequest(
			reflect.TypeOf(argsUnBook{}),
			handleUnBook, checks, finish,
		),
	)
}
