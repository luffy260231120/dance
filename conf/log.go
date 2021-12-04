package conf

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"time"
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
		path := "./logs/log"
		LogMaxAge := 24 * 60 //hour
		LogRotatTm := 60     //min
		writer, _ := rotatelogs.New(
			//path+".%Y%m%d%H%M",
			path+".%Y%m%d%H",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithMaxAge(time.Duration(LogMaxAge)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(LogRotatTm)*time.Minute),
		)
		MainLog.SetOutput(writer)
	}
}
