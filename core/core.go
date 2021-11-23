package core

import (
	log2 "dance/conf"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type FunCheck func(c *Context) bool
type FunHandle func(*Context)
type FunTMESvc func(*Context, *http.Client)

type Context struct {
	*gin.Context
	Result gin.H
	Code   int
}

func (c *Context) JSON(code int, result gin.H) {
	c.Code = code
	c.Result = result
}

func replyJSONAndLog(c *Context) {
	f := log.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"query":  c.Request.URL.RawQuery,
	}
	if args, ok := c.Keys["args"]; ok {
		para := reflect.Indirect(reflect.ValueOf(args))
		if para.FieldByName("UserID").IsValid() {
			f["userid"] = para.FieldByName("UserID").Interface()
		}
	}
	body, _ := c.Get(gin.BodyBytesKey)
	if body != nil {
		f["body"] = string(body.([]byte))
	} else {
		f["body"] = nil
	}
	f["code"] = c.Code
	if c.Result != nil {
		dats, _ := json.Marshal(c.Result)
		if len(dats) > 4096 {
			f["result"] = string(dats[:4096])
		} else {
			f["result"] = string(dats)
		}
	} else {
		f["result"] = nil
	}
	log2.MainLog.WithFields(f).Info("request")
}

func HandlePost(typ reflect.Type, handle FunHandle, checks []FunCheck, finish []FunHandle) func(*gin.Context) {
	return func(c *gin.Context) {
		sign := c.GetHeader("signature")
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log2.MainLog.Errorf("handle post err:%v", err.Error())
			return
		}
		if sign == "" {
			temp := uuid.New()
			sign = strings.ToLower(hex.EncodeToString(temp[:]))
		} else {
			sign = strings.ToLower(sign)
		}
		ctx := Context{Context: c}
		c.Set("ctx", &ctx)
		c.Set("type", typ)
		c.Set(gin.BodyBytesKey, body)
		typs := c.Keys["type"].(reflect.Type)
		argv := reflect.New(typs)
		args := argv.Interface()
		errs := c.ShouldBindBodyWith(args, binding.JSON)
		if errs == nil {
			c.Set("args", args)
		} else {
			log2.MainLog.Errorf("get args failed.err:%v", err.Error())
			return
		}
		defer replyJSONAndLog(&ctx)

		for _, cb := range checks {
			if !cb(&ctx) {
				return
			}
		}
		handle(&ctx)
		for _, cb := range finish {
			cb(&ctx)
		}
		c.JSON(ctx.Code, ctx.Result)
	}
}

var Engine *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	Engine = gin.New()
	Engine.Use(RecoveryWithWriter())
	Engine.RedirectTrailingSlash = false
}
