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
	for _, db := range conf.Dbs {
		err := tools.CreateDbConn(db.Database, db.Host, db.Port, db.User, db.Password, db.MaxConn, db.MaxIdle, db.LifeTime)
		tools.CheckErr(err)
	}

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "../templates",    // Specify what path to load the templates from.
		Extensions: []string{".html"}, // Specify extensions to load for templates.
	}))

	//初始化路由
	InitRoute(m)

	m.RunOnAddr(fmt.Sprintf(":%d", config.Get().Http.Port))
}
