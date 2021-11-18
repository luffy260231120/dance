package core

import (
	log2 "dance/conf"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type FunCheck func(c *gin.Context) bool
type FunHandle func(*gin.Context)
type FunTMESvc func(*gin.Context, *http.Client)

func replyJSONAndLog(c *gin.Context) {
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
		c.Set("ctx", c)
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
		defer replyJSONAndLog(c)

		for _, cb := range checks {
			if !cb(c) {
				return
			}
		}
		handle(c)
		for _, cb := range finish {
			cb(c)
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
