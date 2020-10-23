安装Go语言
=========
https://golang.google.cn/dl/

执行
====
``` go
go run xxx.go
```

fmt
===
使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数

错误信息
========
``` go
cannot run non-main package
```

错误原因
========
main方法只能放在package main中，go run 是执行命令，必须要一个main用来调用

电子书
======
https://studygolang.com/subject/2