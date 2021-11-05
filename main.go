package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	//"os"

	_ "github.com/gin-contrib/gzip"

	"dance/conf"
	//"dance/cons"
	"dance/core"
	//_ "dance/iauthorize" // 授权相关
	//_ "dance/iauthorize/contact"
	//_ "dance/ipay" // 支付购买相关
	//_ "dance/isdk"
	//_ "dance/isong"
	//_ "dance/istat"
	//_ "dance/ivip"
	//"dance/patch"
	//log "github.com/sirupsen/logrus"
)

//
//func init() {
//	var pretty bool
//	log.SetLevel(log.DebugLevel)
//	log.SetReportCaller(true)
//	if conf.Config.Mode == cons.MODE_DEV {
//		pretty = true
//		log.SetOutput(os.Stdout)
//	} else {
//		log.SetOutput(conf.LoggerDebug)
//	}
//	format := patch.LogFormatter{&log.JSONFormatter{PrettyPrint: pretty}}
//	log.SetFormatter(&format)
//}

func main1() {
	hosts := fmt.Sprintf(conf.Config.Env["listen"])
	//endless.ListenAndServe(hosts, core.Engine)
	core.Engine.Run(hosts)
}

func main() {
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})

	r.Run(":1234")
}
