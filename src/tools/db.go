package tools

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
)

type DbConn struct {
	Name  string
	DbMap *gorp.DbMap
}

var dbConnMap map[string]*DbConn

func init() {
	dbConnMap = make(map[string]*DbConn)
}

func GetDefDb() *DbConn {
	return GetDb("default")
}

func GetDb(name string) *DbConn {
	return dbConnMap[name]
}

func CreateDbConn(name string, host string, port int, user, password, database string, maxConn, maxIdle, lifeTime int) (err error) {
	if _, exist := dbConnMap[name]; exist {
		return
	}

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)

	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdle)
	db.SetConnMaxLifetime(time.Duration(lifeTime) * time.Second)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	dbConnMap[name] = &DbConn{
		Name:  name,
		DbMap: dbmap,
	}

	return

}