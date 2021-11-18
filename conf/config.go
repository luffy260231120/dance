package conf

import (
	//"fmt"

	"github.com/BurntSushi/toml"
	//"github.com/astaxie/beego/logs"
)

type CConfig struct {
	Listen        string
	Database      string
	Redis         string
	RedisPassword string
	SkipSignCheck int
}

var Config CConfig

func InitConfig() {
	// config
	fileConfig := ""
	if Dev == DevTest {
		fileConfig = "conf/server_test.toml"
	} else {
		fileConfig = "conf/server.toml"
	}
	if _, err := toml.DecodeFile(fileConfig, &Config); err != nil {
		panic(err)
	}
	MainLog.Infof("config file:%v", Config)
}
