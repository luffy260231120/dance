package core

import (
	log2 "dance/conf"
	"dance/cons"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reflect"
)

type FunCheck func(c *Context) bool
type FunHandle func(*Context)
type FunTMESvc func(*Context, *http.Client)

type Context struct {
	*gin.Context
	Result   gin.H
	HttpCode int
	Code     int
	Msg      string
}

func (c *Context) JSON(httpCode, code int, msg string, result gin.H) {
	c.HttpCode = httpCode
	c.Code = code
	c.Msg = msg
	c.Result = result
}

func replyJSONAndLog(ctx *Context) {
	var (
		f = log.Fields{
			"method": ctx.Request.Method,
			"path":   ctx.Request.URL.Path,
			"query":  ctx.Request.URL.RawQuery,
		}
		c      = ctx.Context
		result = gin.H{"code": ctx.Code, "msg": ctx.Msg}
	)

	//if args, ok := c.Keys["args"]; ok {
	//	para := reflect.Indirect(reflect.ValueOf(args))
	//	if para.FieldByName("UserID").IsValid() {
	//		f["userid"] = para.FieldByName("UserID").Interface()
	//	}
	//}
	body, _ := ctx.Get(gin.BodyBytesKey)
	if body != nil {
		f["body"] = string(body.([]byte))
	} else {
		f["body"] = nil
	}
	f["code"] = ctx.Code
	f["code_msg"] = ctx.Msg
	if ctx.Result != nil {
		result["data"] = ctx.Result
		dats, _ := json.Marshal(ctx.Result)
		if len(dats) > 4096 {
			f["result"] = string(dats[:4096])
		} else {
			f["result"] = string(dats)
		}
	}
	log2.MainLog.WithFields(f).Info("request")
	c.JSON(ctx.HttpCode, result)
}

type ArgsDefault struct{}

func HandleRequest(typ reflect.Type, handle FunHandle, checks []FunCheck, finish []FunHandle) func(*gin.Context) {
	return func(c *gin.Context) {
		//sign := c.GetHeader("sign")

		//if sign == "" {
		//	temp := uuid.New()
		//	sign = strings.ToLower(hex.EncodeToString(temp[:]))
		//} else {
		//	sign = strings.ToLower(sign)
		//}
		ctx := Context{Context: c}
		//c.Set("sign", sign)
		c.Set("ctx", &ctx)
		c.Set("type", typ)

		defer replyJSONAndLog(&ctx)
		if msg := CheckArgsQueryBody(&ctx); msg != "" {
			ctx.JSON(ctx.HttpCode, cons.ERR_PUB_PARAMS, msg, nil)
			return
		}

		for _, cb := range checks {
			if !cb(&ctx) {
				c.JSON(ctx.HttpCode, gin.H{"code": ctx.Code, "msg": ctx.Msg, "data": ctx.Result})
				return
			}
		}
		handle(&ctx)
		for _, cb := range finish {
			cb(&ctx)
		}
	}
}

var Engine *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	Engine = gin.New()
	Engine.Use(RecoveryWithWriter())
	Engine.RedirectTrailingSlash = false
}
