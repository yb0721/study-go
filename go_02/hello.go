package main
import "fmt"

// 此方法一般用于声明全局变量
var (
	x int
	y bool
)

func main() {
	var a string = "Runoob"
	fmt.Println(a)
	
	var b, c int = 1, 2
	fmt.Println(b, c)

	//声明一个变量并初始化
	var a1 = "RUNOOB"
	fmt.Println(a1)

	//没有初始化就为0
	var b1 int
	fmt.Println(b1)

	//bool默认值为False
	var c1 bool
	fmt.Println(c1)

	//省略var声明变量
	//var f string = "Runoob"
	f := "Runoob" 
	fmt.Println(f)

	fmt.Println(x, y)
}
