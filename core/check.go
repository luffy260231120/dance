package core

import (
	"crypto/md5"
	"dance/conf"
	"dance/cons"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	pKey = "123"
)

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
