package course

import (
	"dance/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"time"
)

const (
	MaxDays = 7
)

func handleDays(c *core.Context) {

	var (
		days = []gin.H{}
	)
	for offset := 0; offset < MaxDays; offset++ {
		t := time.Now().AddDate(0, 0, offset)
		days = append(days, gin.H{
			"desc":   getDayDesc(t),
			"offset": offset,
		})
	}
	c.JSON(http.StatusOK, 0, "", gin.H{"days": days})
}

func getDayDesc(t time.Time) string {
	week := ""
	switch t.Weekday() {
	case 1:
		week = "周一"
	case 2:
		week = "周二"
	case 3:
		week = "周三"
	case 4:
		week = "周四"
	case 5:
		week = "周五"
	case 6:
		week = "周六"
	case 0:
		week = "周日"
	}
	return fmt.Sprintf("%s %d.%d", week, t.Month(), t.Day())
}

func init() {
	checks := []core.FunCheck{}
	finish := []core.FunHandle{}
	core.Engine.GET(
		"/dance/course/days",
		core.HandleRequest(
			reflect.TypeOf(core.ArgsDefault{}),
			handleDays, checks, finish,
		),
	)
}
