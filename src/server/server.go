package main

import (
	"config"
	"fmt"
	"tools"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	//load config
	config.Load()

	conf := config.Get()

	//init DB
	err := tools.CreateDbConn("default",
		conf.Db.Host,
		conf.Db.Port,
		conf.Db.User, conf.Db.Password, conf.Db.Database,
		conf.Db.MaxConn,
		conf.Db.MaxIdle,
		conf.Db.LifeTime,
	)
	tools.CheckErr(err)

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "../templates",    // Specify what path to load the templates from.
		Extensions: []string{".html"}, // Specify extensions to load for templates.
	}))

	//初始化路由
	InitRoute(m)

	m.RunOnAddr(fmt.Sprintf(":%d", config.Get().Http.Port))
}
