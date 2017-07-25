package main

import "fmt"

func main() {
	var min ,max int
	min,max = Minmax(23,31)
	fmt.Printf("min ",min,max)
}

func Minmax(a int ,b int) (min int ,max int)  {
	if a<b {
		min =a
		max = b
	}else{
		min = b
		max =a
	}
	return min,max
}
