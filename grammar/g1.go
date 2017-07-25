package main

import "fmt"

/**
	基础语法
		1.一行代表一个语句结束，不需要分号结束。
		2.多个语句要写在同一行，则必须使用; 区分。
	数据类型
		bool : true/false
		1.go不允许不同类型之间的混合使用，如 var a int=5; var b int32=a+a;
	变量
	值类型和引用类型：
		1.int float bool string 基本类型，都属于值类型，使用这些类型的变量，直接指向存在内存中的值。
		2.可以通过 &i 来获取变量i的内存地址，每次的地址可能不一样。
		3.值类型的变量的值存在于 栈中。

 */
var (  // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

func main() {
	//var 可以省略， 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。
	//这种不带声明格式的只能在函数体中出现
	//且必须要被引用，否则编译器报错。

	g, h := 123, "hello2"
	var k int =g;
	g=14;

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(k)

}





