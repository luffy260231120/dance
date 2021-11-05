package core

import (
	//"dance/conf"
	"dance/cons"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type FunCheck func(c *Context) bool
type FunHandle func(*Context)
type FunTMESvc func(*Context, *http.Client)

type Context struct {
	//*conf.V2PInfo
	*gin.Context
	*log.Entry
	STime      time.Time //请求开始处理的时间
	RTime      time.Time //设备注册时间
	DevID      int
	Code       int
	Info       gin.H
	rOnce      sync.Once
	TokenVer   int   //token版本
	SkipAuth   bool  //是否开启匿名调用api
	SkipToken  bool  //是否跳过token检查
	SkipDid    bool  //是否跳过设备码检查
	ExpireTime int64 //token过期时间

	mu        sync.RWMutex //保护可能被并发访问的字段
	HttpTrace map[string]map[string]string
}

func (c *Context) JSON(code int, info gin.H) {
	// 存在并发调用此函数的情况，请参见QM user info接口
	c.rOnce.Do(func() {
		c.Code = code
		c.Info = info
	})
}

func replyJSONAndLog(c *Context) {
	f := log.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"query":  c.Request.URL.RawQuery,
		"delay":  time.Now().Sub(c.STime).Milliseconds(),
		"code":   c.Code,
	}
	if args, ok := c.Keys["args"]; ok {
		para := reflect.Indirect(reflect.ValueOf(args))
		if para.FieldByName("PID").IsValid() {
			f["pid"] = para.FieldByName("PID").Interface()
		}
		if para.FieldByName("DeviceID").IsValid() {
			f["deviceid"] = para.FieldByName("DeviceID").Interface()
		}
		if para.FieldByName("UserID").IsValid() {
			f["userid"] = para.FieldByName("UserID").Interface()
		}
		if para.FieldByName("SP").IsValid() {
			f["path"] = fmt.Sprintf("%v:%s", para.FieldByName("SP").Interface(), c.Request.URL.Path)
		}
	}
	if c.Info != nil {
		info := c.Info
		result := gin.H{}
		if code, ok := info["error_code"]; ok {
			result["error_code"] = code
		}
		if emsg, ok := info["error_msg"]; ok {
			result["error_msg"] = emsg
		}
		if data, ok := info["data"]; ok {
			dats, _ := json.Marshal(data)
			if len(dats) > 4096 {
				result["result-data"] = string(dats[:4096])
			} else {
				result["result-data"] = string(dats)
			}
		}
		f["result"] = result
	} else {
		f["result"] = nil
	}
	body, _ := c.Get(gin.BodyBytesKey)
	if body != nil {
		f["body"] = string(body.([]byte))
	} else {
		f["body"] = nil
	}
	c.WithFields(f).Info("gateway")
	if c.HttpTrace != nil {
		c.Infof("%s", c.HttpTrace)
	}
}

//func HandleCore(typ reflect.Type, handle FunHandle, checks []FunCheck, finish []FunHandle) func(*gin.Context) {
//	return func(c *gin.Context) {
//		sign := c.GetHeader("signature")
//		if sign == "" {
//			temp := uuid.New()
//			sign = strings.ToLower(hex.EncodeToString(temp[:]))
//		} else {
//			sign = strings.ToLower(sign)
//		}
//		log := log.WithFields(log.Fields{"signature": sign})
//		ctx := Context{Context: c, Entry: log}
//		c.Set("ctx", &ctx)
//		c.Set("type", typ)
//		ctx.STime = time.Now()
//		defer replyJSONAndLog(&ctx)
//		for _, cb := range checks {
//			if !cb(&ctx) {
//				return
//			}
//		}
//		handle(&ctx)
//		for _, cb := range finish {
//			cb(&ctx)
//		}
//	}
//}
//
//func HandleSP(h map[string]FunTMESvc, checks map[string][]FunCheck, finish map[string][]FunHandle) FunHandle {
//	return func(c *Context) {
//		var (
//			args   = c.Keys["args"]
//			vals   = reflect.Indirect(reflect.ValueOf(args))
//			sp     = vals.FieldByName("SP").String()
//			//client = spClient[sp]
//		)
//		c.Header("sp", sp)
//		for _, cb := range checks[sp] {
//			if !cb(c) {
//				return
//			}
//		}
//		h[sp](c, client)
//		for _, cb := range finish[sp] {
//			cb(c)
//		}
//	}
//}

func HandleDummy(c *Context, client *http.Client) {
	c.JSON(http.StatusOK, gin.H{"error_code": cons.ERR_PUB_UNIMPL, "error_msg": ""})
}

var BGCtx *Context
var Engine *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	Engine = gin.New()
	//Engine.Use(RecoveryWithWriter())
	Engine.RedirectTrailingSlash = false

	//pinfo := conf.V2PInfo{
	//	PID:      "200023",
	//	BID:      "10003",
	//	DType:    "智能电视",
	//	DClass:   cons.DTYPE_CLASS_TV,
	//	KGAppID:  "3154",
	//	KGAppKey: "oWuyfzTNfrdRhFYuaWtlcMITsi443aDK",
	//}
	BGCtx = &Context{
		//V2PInfo: &pinfo,
		Entry:   log.WithFields(log.Fields{"signature": "background"}),
		Context: &gin.Context{Keys: map[string]interface{}{"args": struct{}{}}},
	}

}
