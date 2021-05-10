package config

import (
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Dbms      string
	User      string
	Pass      string
	Protocol  string
	Dbname    string
	Secretkey string
}

var Config ConfigList

// SQLConnect DB接続
func init() {

	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err.Error())
	}

	Config = ConfigList{
		Dbms:      cfg.Section("db").Key("host").String(),
		User:      cfg.Section("db").Key("user_name").String(),
		Pass:      cfg.Section("db").Key("password").String(),
		Protocol:  cfg.Section("db").Key("protocol").String(),
		Dbname:    cfg.Section("db").Key("db_name").String(),
		Secretkey: cfg.Section("jwt").Key("secretkey").String(),
	}
}
