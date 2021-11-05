package conf

import (
	"dance/cons"
	//"fmt"

	"github.com/BurntSushi/toml"
	//"github.com/astaxie/beego/logs"
)

type CConfig struct {
	Mode      string
	IDC       string
	ServerID  string
	ServerKey string
	CodeKey   string
	AESKey    string
	Env       map[string]string
}

//
//type V2PInfo struct {
//	PKey     string `gorm:"-"`
//	RKey     string `gorm:"-"`
//	DType    string `gorm:"column:type"`
//	PID      string `gorm:"column:id"`
//	BID      string `gorm:"column:bid"`
//	DClass   string `gorm:"column:class"`
//	KGAppID  string `gorm:"-"`
//	KGAppKey string `gorm:"-"`
//}

//var V2PInfos map[string]*V2PInfo
var Config CConfig

func init() {
	// config
	if _, err := toml.DecodeFile("conf/server.toml", &Config); err != nil {
		panic(err)
	}
	env := map[string]string{}
	if _, err := toml.DecodeFile("conf/idc.toml", &env); err != nil {
		panic(err)
	}
	for k, v := range env {
		Config.Env[k] = v
	}
	Config.IDC = Config.Env["idc"]
	Config.Mode = Config.Env["mode"]
	if _, ok := cons.MapMode[Config.Mode]; !ok {
		panic("invalid server mode")
	}
}
