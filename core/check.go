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
	pKey = "123"
)

//func CheckArgsQuery(c *Context) bool {
//	var (
//		typs = c.Keys["type"].(reflect.Type)
//		argv = reflect.New(typs)
//		args = argv.Interface()
//	)
//	if err := patch.MapUri(args, c.Request.URL.Query()); err != nil {
//		conf.MainLog.Warnf(err.Error())
//		c.JSON(http.StatusOK, cons.ERR_PUB_PARAMS, "invalid query", nil)
//		return false
//	}
//	errs := binding.Validator.ValidateStruct(args)
//	if errs == nil {
//		c.Set("args", args)
//		return true
//	}
//	etip := ""
//	tips := []string{}
//	switch errs.(type) {
//	case validator.ValidationErrors:
//		verr := errs.(validator.ValidationErrors)
//		for _, f := range verr {
//			tip := fmt.Sprintf("field:%s,rule:%s", strings.ToLower(f.Field()), f.Tag())
//			tips = append(tips, tip)
//		}
//	}
//	//c.Debugf("%T, %v\n", errs, errs)
//	if len(tips) == 0 {
//		etip = "invalid query format"
//	} else {
//		etip = strings.Join(tips, "|")
//	}
//	c.JSON(http.StatusOK, cons.ERR_PUB_PARAMS, etip, nil)
//	return false
//}

func CheckArgsBody(c *Context) string {
	typs := c.Keys["type"].(reflect.Type)
	argv := reflect.New(typs)
	args := argv.Interface()
	errs := c.ShouldBindBodyWith(args, binding.JSON)
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
	case *json.UnmarshalTypeError:
		verr := errs.(*json.UnmarshalTypeError)
		tip := fmt.Sprintf("field:%s,rule:type", strings.ToLower(verr.Field))
		tips = append(tips, tip)
	}
	//c.Debugf("%T, %v\n", errs, errs)
	if len(tips) == 0 {
		etip = "invalid json format"
	} else {
		etip = strings.Join(tips, "|")
	}
	return etip
}

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
