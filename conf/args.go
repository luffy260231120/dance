package conf

import (
	"os"
)

const (
	DevTest = "test"
	DevProv = "prov"
)

var Dev = ""

func InitArgs() {
	if len(os.Args) != 2 || (os.Args[1] != DevProv && os.Args[1] != DevTest) {
		MainLog.Error("入参错误")
		return
	}
	Dev = os.Args[1]
	MainLog.Infof("目前运行状态:%s", Dev)
}
