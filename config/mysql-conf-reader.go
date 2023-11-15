package config

import (
	"github.com/fatih/color"
	"gopkg.in/ini.v1"
)

type MysqlConf struct {
	MysqlAddr     string
	MysqlUser     string
	MysqlPassword string
}

func MysqlConfReader(path string) (*MysqlConf, error) {
	conf, error := ini.Load(path)

	if error != nil {
		color.HiRedString("Load target path failed, details: " + error.Error())
		return nil, error
	}

	ms := conf.Section("mysql")

	addr := ms.Key("mysql_addr").String()
	user := ms.Key("mysql_user").String()
	password := ms.Key("mysql_password").String()

	return &MysqlConf{
		MysqlAddr:     addr,
		MysqlUser:     user,
		MysqlPassword: password,
	}, nil
}
