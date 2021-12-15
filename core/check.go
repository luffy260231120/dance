package core

import (
	"bytes"
	"crypto/md5"
	"dance/conf"
	"dance/cons"
	"dance/patch"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

const (
	pKey      = "123"
	Anonymous = "anonymous"
	Password  = "password"
)

func CheckArgsQueryBody(c *Context) string {
	var (
		typs = c.Keys["type"].(reflect.Type)
		argv = reflect.New(typs)
		args = argv.Interface()
		ctyp = c.Request.Header.Get("Content-Type")
		mime = "application/json"
	)
	if len(ctyp) >= len(mime) && ctyp[:len(mime)] == mime {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			conf.MainLog.Warnf(err.Error())
			return "body read failed"
		}
		c.Set(gin.BodyBytesKey, body)
		decoder := json.NewDecoder(bytes.NewReader(body))
		if err := decoder.Decode(args); err != nil && err != io.EOF {
			conf.MainLog.Warnf(err.Error())
			return "invalid json"
		}
	} else {
		c.Set(gin.BodyBytesKey, []byte(""))
	}
	if err := patch.MapForm(args, c.Request.URL.Query()); err != nil {
		conf.MainLog.Warnf(err.Error())
		return "invalid query"
	}
	errs := binding.Validator.ValidateStruct(args)
	if errs == nil {
		c.Set("args", args)
		return ""
	}
	etip := ""
	tips := []string{}
	switch errs.(type) {
	case validator.ValidationErrors:
		verr := errs.(validator.ValidationErrors)
		for _, f := range verr {
			tip := fmt.Sprintf("field:%s,rule:%s", strings.ToLower(f.Field()), f.Tag())
			tips = append(tips, tip)
		}
	}
	//c.Debugf("%T, %v\n", errs, errs)
	if len(tips) == 0 {
		etip = "invalid json format"
	} else {
		etip = strings.Join(tips, "|")
	}
	return etip
}

func CheckSignMD5(c *Context) bool {
	if conf.Config.SkipSignCheck == 1 {
		return true
	}
	var (
		md5cal = md5.New()
	)
	body := c.Keys[gin.BodyBytesKey]
	md5cal.Write(body.([]byte))
	md5cal.Write([]byte(pKey))
	calsign := fmt.Sprintf("%X", md5cal.Sum(nil))
	reqsign := c.GetHeader("signature")
	if strings.ToUpper(reqsign) != calsign {
		c.JSON(http.StatusOK, cons.ERR_PUB_PARAMS, "签名错误", nil)
		return false
	}
	return true
}

func CheckManager(c *Context) bool {
	//reqsign := c.GetHeader("signature")
	//if strings.ToUpper(reqsign) != calsign {
	//	c.JSON(http.StatusOK, gin.H{"error_code": cons.ERR_PUB_PARAMS, "error_msg": "invalid signature"})
	//	return false
	//}
	return true
}

// 免登陆
func CheckVisitAuth(c *Context) bool {
	var (
		args   = reflect.Indirect(reflect.ValueOf(c.Keys["args"]))
		userid = args.FieldByName("UserID").String()
		token  = args.FieldByName("Token").String()
	)
	if userid != Anonymous || token != Password {
		return true
	}
	args.FieldByName("UserID").SetString("")
	args.FieldByName("Token").SetString("")
	c.SkipAuth = true
	return true
}

func CheckUserId(c *Context) bool {
	if c.SkipAuth {
		return true
	}
	// TODO
	return true
}

func CheckToken(c *Context) bool {
	if c.SkipAuth {
		return true
	}
	// TODO
	return true
}
