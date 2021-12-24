package conf

import (
	"errors"
	"os"
)

const (
	DevTest = "test"
	DevProv = "prov"
)

var Dev = ""

func InitArgs() error {
	if len(os.Args) != 2 || (os.Args[1] != DevProv && os.Args[1] != DevTest) {
		MainLog.Error("入参错误，启动失败")
		return errors.New("入参错误，启动失败")
	}
	Dev = os.Args[1]
	MainLog.Infof("目前运行状态:%s", Dev)
	return nil
}
