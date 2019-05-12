//当前程序包名
package main

//导入其他的包
import std "fmt"

//常量定义
const PI = 3.14

//全局变量的声明与复制
var name = "gopher"

//一般类型声明
type newType int

//结构的声明
type gopher struct{}

//接口的声明
type golang interface{}

//使用 main 函数作为程序的入口点启动
func main() {
	std.Println("Hellow World!")
}
