package itest

import (
	"dance/core"
	"dance/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type argsTest struct {
	model.UserAuthWithoutToken
}

func handleTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"aa": "bb"})
}

func handleTest2(c *core.Context) {
	var (
		args = c.Keys["args"].(*argsTest)
	)
	c.JSON(http.StatusOK, 0, "", gin.H{"user_id": args.UserID})
}

func init() {
	core.Engine.GET("/dance/test", handleTest)

	checks := []core.FunCheck{core.CheckSignMD5}
	finish := []core.FunHandle{}
	core.Engine.POST(
		"/dance/test2",
		core.HandleRequest(
			reflect.TypeOf(argsTest{}),
			handleTest2, checks, finish,
		),
	)
}
