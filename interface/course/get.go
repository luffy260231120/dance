package course

import (
	"dance/conf"
	"dance/cons"
	"dance/core"
	"dance/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"time"
)

type argsGet struct {
	model.UserAuth
	model.Pagination
	Offset int `form:"offset"`
}

type BookStatus int

const (
	CantBook BookStatus = iota
	CanBook
	Booked
	CanWait
	Waiting
	UnLogin
)

type respGet struct {
	model.CourseInfo
	TeacherName string     `gorm:"column:name" json:"name"`
	Sex         int        `gorm:"column:sex" json:"sex"` // 1-女 2-男
	Avatar      string     `gorm:"column:avatar" json:"avatar"`
	Number      int        `json:"number"`
	Status      BookStatus `json:"status"` // 0-不可预约 1-可预约 2-已预约 3-可排队 4-排队中 5-未登录
}

func handleGet(c *core.Context) {

	var (
		args    = c.Keys["args"].(*argsGet)
		cTab    = core.GetDB().Table("course")
		courses = []*respGet{}
		day     = time.Now().AddDate(0, 0, args.Offset).Format(cons.FORMAT_DATE)
	)

	// 1 获取当天所有课程 + 老师信息
	err := cTab.Select("teacher.name,teacher.sex, teacher.avatar, course.id, course.type,course.title,course.max_number,course.bk,course.open_time,course.start_time,course.class").
		Joins("left join teacher on course.teacher_id = teacher.teacher_id").
		Where("course.date = ?", day).
		Limit(args.Size).Offset(args.Size * (args.Page - 1)).Order("start_time").Find(&courses).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, 0, "", gin.H{"courses": courses})
			return
		}
		conf.MainLog.Errorf(err.Error())
		c.JSON(http.StatusOK, cons.ERR_PUB_SYSTEM, "获取课程信息失败", nil)
		return
	}

	// 2 获取预约人数 判断status
	for _, course := range courses {
		key := fmt.Sprintf(cons.CacheKeyBookUserIds, course.ID)
		userIds := core.GetRedis().ZRange(key, 0, -1).Val()
		course.Number = len(userIds)

		// status
		openTime, _ := time.ParseInLocation(cons.FORMAT_TIME, course.OpenTime, time.Local)
		startTime, _ := time.ParseInLocation(cons.FORMAT_TIME, course.StartTime, time.Local)
		if time.Now().Before(openTime) || time.Now().After(startTime) {
			course.Status = CantBook
			continue
		}

		if args.UserID == "" {
			course.Status = UnLogin
			continue
		}

		if course.Number == 0 {
			course.Status = CanBook
			continue
		}

		var index = 0
		for _, userId := range userIds {
			index++
			if userId == args.UserID {
				if index <= course.MaxNumber {
					course.Status = Booked
				} else {
					course.Status = Waiting
				}
				continue
			}
		}

		if course.Status == 0 { // 还未设置状态
			if course.Number < course.MaxNumber {
				course.Status = CanBook
			} else {
				course.Status = CanWait
			}
		}
	}

	c.JSON(http.StatusOK, 0, "", gin.H{"courses": courses})
}

func init() {
	checks := []core.FunCheck{core.CheckVisitAuth, core.CheckUserId, core.CheckToken}
	finish := []core.FunHandle{}
	core.Engine.GET(
		"/dance/course/get",
		core.HandleRequest(
			reflect.TypeOf(argsGet{}),
			handleGet, checks, finish,
		),
	)
}
