package conf

import (
	"github.com/sirupsen/logrus"
	"os"
)

// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
var MainLog = logrus.New()

func InitLog() {
	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	if Dev == DevTest {
		MainLog.Out = os.Stdout
	} else {
		file, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			MainLog.Out = file
		} else {
			MainLog.Info("Failed to log to file")
		}
	}
}
