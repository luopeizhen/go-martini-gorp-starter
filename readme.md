## 说明
这个例子包含了如何搭建martini并返回html或json, go的模板如何使用, 数据库的 查询, 插入, 更新, 删除


### 数据库准备

使用前创建数据库dbtest

导入下面的sql

```sql
CREATE TABLE `table1` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

ALTER TABLE `table1`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `table1`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
```

### 配置
server.conf.sample 改名为 server.conf 放在bin和执行文件放在一起

### 数据库的表格定义
可以使用gen_model.go自动生成, 例如
```
go run src/gen_model/gen_model.go -dbSource "root:123@tcp(localhost:3306)/dbtest" > src/model/model.go
```


### 编译
编译前要把当前目录加入GOPATH, vscode里面已经设置了支持windows编译

vscode里面的默认task是在win10里面编译

linux命令行编译
```
env GOPATH="`pwd`:$GOPATH" go install server
```