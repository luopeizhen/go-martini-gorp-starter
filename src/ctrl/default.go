package ctrl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"model"
	"net/http"
	"strconv"
	"tools"

	"github.com/martini-contrib/render"
)

func getAllItems() (rs []model.Table1) {
	dbConn := tools.GetDefDb()

	_, err := dbConn.DbMap.Select(&rs, "select * from table1")
	tools.CheckErr(err)

	return
}

//默认Home页
func DefaultGetHome(req *http.Request, r render.Render) {

	items := getAllItems()

	r.HTML(200, "default/home", map[string]interface{}{
		"title": "I am title",
		"items": items,
	})
}

//增加一个记录
func DefaultPostHome(req *http.Request, r render.Render) {
	type Input struct {
		Name string
	}

	b, err := ioutil.ReadAll(req.Body)
	tools.CheckErr(err)

	var data Input
	err = json.Unmarshal(b, &data)
	tools.CheckErr(err)

	fmt.Printf("Input item:%+v\n", data)

	dbConn := tools.GetDefDb()
	sql := fmt.Sprintf("insert into `table1` (`name`) values('%s')", data.Name)
	res, err := dbConn.DbMap.Exec(sql)
	tools.CheckErr(err)

	lastId, err := res.LastInsertId()
	tools.CheckErr(err)

	fmt.Println("lastId:", lastId)

	r.JSON(200, nil)
}

//修改一个记录
func DefaultPutHome(req *http.Request, r render.Render) {
	type Input struct {
		Id   int
		Name string
	}

	b, err := ioutil.ReadAll(req.Body)
	tools.CheckErr(err)

	var data Input
	err = json.Unmarshal(b, &data)
	tools.CheckErr(err)

	fmt.Printf("Input item:%+v\n", data)

	dbConn := tools.GetDefDb()
	sql := fmt.Sprintf("update `table1` set `name`='%s' where `id`=%d", data.Name, data.Id)
	res, err := dbConn.DbMap.Exec(sql)
	tools.CheckErr(err)

	affectedCount, err := res.RowsAffected()
	tools.CheckErr(err)

	fmt.Println("affectedCount:", affectedCount)

	r.JSON(200, nil)
}

//删除
func DefaultDeleteHome(req *http.Request, r render.Render) {

	req.ParseForm()

	id, _ := strconv.ParseInt(req.Form.Get("id"), 10, 64)

	dbConn := tools.GetDefDb()
	sql := fmt.Sprintf("delete from table1 where id=%d", id)
	res, err := dbConn.DbMap.Exec(sql)
	tools.CheckErr(err)

	affectedCount, err := res.RowsAffected()
	tools.CheckErr(err)

	fmt.Println("affectedCount:", affectedCount)

	r.JSON(200, nil)

}
