package main

import "fmt"

/**
http://www.runoob.com/go/go-program-structure.html
1.当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，
	如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
	标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。
2.命令行执行：go run xxx.go
第一个go程序demo
2017-05-05
主函数入口
 */
func main() {
	fmt.Print("hello world")
}

/**
init方法先于main被执行
 */
func init()  {
	fmt.Println("init ")
}