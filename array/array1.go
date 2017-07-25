package main

import "fmt"

//数组的长度不同 是作为不同的数据类型。

func main() {

	a := [...]int{1,2}
	fmt.Println(a)

	var arrkeyvalue = [5]string{3:"chars",4:"ron"}
	fmt.Println(arrkeyvalue)

}
