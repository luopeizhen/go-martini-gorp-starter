package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"tools"
)

type Conf struct {
	Http struct {
		Port int
	}
	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
		MaxConn  int
		MaxIdle  int
		LifeTime int
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
