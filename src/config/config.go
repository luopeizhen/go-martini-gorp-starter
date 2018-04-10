package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"tools"
)

type Db struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	MaxConn  int `json:"max_conn"`
	MaxIdle  int `json:"max_idle"`
	LifeTime int `json:"life_time"`
}

type Conf struct {
	Http struct {
		Port int
	}
	Dbs []Db

	Redis struct {
		Source      string
		MaxActive   int `json:"max_active"`
		MaxIdle     int `json:"max_idle"`
		IdleTimeout int `json:"idle_timeout"`
	}
}

var conf *Conf

func Get() *Conf {
	return conf
}

func Load() {
	conf = &Conf{}
	b, err := ioutil.ReadFile("server.conf")
	tools.CheckErr(err)

	err = json.Unmarshal(b, conf)
	tools.CheckErr(err)

	fmt.Printf("Config Loaded:%+v\n", conf)

}
