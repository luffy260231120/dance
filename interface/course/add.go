package course

import (
	"dance/core"
	"reflect"
)

type argsAdd struct {
	Type      int    `json:"type" binding:"required"`
	Name      string `json:"name" binding:"required"`
	TeacherId int    `json:"teacher_id" binding:"required"`
	Bk        string `json:"bk" binding:"required"`
	OpenTime  string `json:"open_time" binding:"required,time"`
	StartTime string `json:"start_time" binding:"required,time"`
	EndTime   string `json:"end_time" binding:"required,time"`
	Class     string `json:"class"`
}

func handleAdd(c *core.Context) {

}

func init() {
	checks := []core.FunCheck{}
	finish := []core.FunHandle{}
	core.Engine.POST(
		// 添加课程信息
		"/dance/course/add",
		core.HandleRequest(
			reflect.TypeOf(argsAdd{}),
			handleAdd, checks, finish,
		),
	)
}
