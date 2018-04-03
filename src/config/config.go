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
		Source   string
		MaxConn  int `json:"max_conn"`
		MaxIdle  int `json:"max_idle"`
		LifeTime int `json:"life_time"`
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
