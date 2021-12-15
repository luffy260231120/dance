package main

import (
	"dance/conf"
	"dance/core"
	_ "dance/interface/course"
	_ "dance/interface/itest"
	_ "dance/interface/teacher"
	_ "dance/interface/user"
	"fmt"
	_ "github.com/gin-contrib/gzip"
)

func main() {
	conf.InitArgs()
	conf.InitConfig()
	conf.InitLog()

	core.InitDB()
	core.InitRedis()
	hosts := fmt.Sprintf(conf.Config.Listen)
	core.Engine.Run(hosts)
}
