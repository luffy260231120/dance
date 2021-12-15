package course

import (
	"dance/cons"
	"dance/core"
	"dance/model"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

func handleAdd(c *core.Context) {
	var (
		args = c.Keys["args"].(*model.CourseInfo)
		now  = time.Now().Format(cons.FORMAT_TIME)
	)

	args.CreateTime = now
	args.UpdateTime = now

	if err := core.GetDB().Table("course").Create(args).Error; err != nil {
		c.JSON(http.StatusOK, cons.ERR_PUB_SYSTEM, fmt.Sprintf("新增课程出错,err:%v", err.Error()), nil)
		return
	}
	c.JSON(http.StatusOK, 0, "", nil)
}

func init() {
	checks := []core.FunCheck{}
	finish := []core.FunHandle{}
	core.Engine.POST(
		// 添加课程信息
		"/dance/course/add",
		core.HandleRequest(
			reflect.TypeOf(model.CourseInfo{}),
			handleAdd, checks, finish,
		),
	)
}
