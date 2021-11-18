package main

import (
	"dance/conf"
	"dance/core"
	_ "dance/itest"
	"fmt"
	_ "github.com/gin-contrib/gzip"
)

func main() {
	conf.InitArgs()
	conf.InitConfig()
	conf.InitLog()
	hosts := fmt.Sprintf(conf.Config.Listen)
	core.Engine.Run(hosts)
}
