package course

import (
	"dance/conf"
	"dance/cons"
	"dance/core"
	"dance/model"
	"fmt"
	"github.com/chenhg5/collection"
	"github.com/go-redis/redis"
	"net/http"
	"reflect"
	"time"
)

type argsBook struct {
	model.UserAuth
	Status    BookStatus `json:"status" binding:"required"`
	CourseId  int        `json:"course_id" binding:"required"`
	MaxNumber int        `json:"max_number" binding:"required"` // TODO 最大人数这里应该做一个预约池，以后需要优化
	StartTime string     `json:"start_time" binding:"required"`
}

func handleBook(c *core.Context) {
	var (
		args          = c.Keys["args"].(*argsBook)
		key           = fmt.Sprintf(cons.CacheKeyBookUserIds, args.CourseId)
		userIds       = core.GetRedis().ZRange(key, 0, -1).Val()
		number        = len(userIds)
		startTimeT, _ = time.ParseInLocation(cons.FORMAT_TIME, args.StartTime, time.Local)
	)

	if collection.Collect(userIds).Contains(args.UserID) {
		c.JSON(http.StatusOK, cons.ERR_PUB_PARAMS, "重复预约", nil)
		return
	}

	err := core.GetRedis().ZAdd(key, redis.Z{
		Score:  float64(startTimeT.Unix() - time.Now().Unix()),
		Member: args.UserID,
	}).Err()
	if err != nil {
		conf.MainLog.Errorf(err.Error())
		c.JSON(http.StatusOK, cons.ERR_PUB_SYSTEM, "预约失败 未知错误", nil)
		return
	}

	if number >= args.MaxNumber {
		c.JSON(http.StatusOK, cons.ERR_PUB_PARAMS, "无法预约 开始排队", nil)
		return
	}

	// TODO 扣卡次数

	c.JSON(http.StatusOK, 0, "预约成功", nil)
}

func init() {
	checks := []core.FunCheck{core.CheckUserId, core.CheckToken}
	finish := []core.FunHandle{}
	core.Engine.POST(
		"/dance/course/book",
		core.HandleRequest(
			reflect.TypeOf(argsBook{}),
			handleBook, checks, finish,
		),
	)
}
