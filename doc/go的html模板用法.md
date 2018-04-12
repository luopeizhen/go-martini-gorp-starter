
go-martini输出html是使用go的标准库 text/template包实现的, 这个go的模板和php, nodejs, jsp那些的原理其实是差不多的, 流程为

1. 在template文件内标记一些变量, 
1. 程序要输出时, 代码传入变量到模板文件, 然后执行渲染
1. 渲染理论上是把变量的值替换模板里面的标记, 当然还可以做一些几个文件组合和循环等工作
1. 渲染完成后输出

代码里面有演示模板如何嵌套, 使用传入的变量和循环

其他用法可参考:
> https://studygolang.com/articles/3071


