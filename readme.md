## 说明

代码主要是演示使用go-martini搭建基本webserver, 返回html或json, go的模板如何使用, 数据库的 查询, 插入, 更新, 删除

还包含了
- 搭建go开发环境
- go语言的一些学习网站
- gopath的说明
- 一个web server例子
- 编译go代码到exe
- 监控go代码, 如果发现修改自动编译


!!! 使用前先看doc里面的文档



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
server.conf.sample 复制为 server.conf, 然后修改配置信息

### 数据库的表格定义
gorp库读取数据库表的记录, 要先定义这个表的结构,
可以使用gen_model.go自动生成, 例如
```
go run src/gen_model/gen_model.go -dbSource "root:123@tcp(localhost:3306)/dbtest" > src/model/model.go
```

### 编译前准备
编译前需要下载一些第三方库, 参考doc里面的文档, 新开一个控制台, 然后执行
```
go get 缺少的库地址
```
这些缺少的库会在编译时报错, 缺哪个就装哪个


### 自动编译
第一次要安装一个监控程序, 自动监控go的源文件目录, 当有修改自动编译重启
新开一个控制台, 然后执行
```
github.com/luopeizhen/gin
```

安装完成后, 执行watch.bat


### 手动编译
双击setvar.bat, 会打开控制台并将当前目录加入到GOPATH 

然后在控制台执行
```
go install server
```
