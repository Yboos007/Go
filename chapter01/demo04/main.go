package main

import "fmt"

func main(){
	//该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
	i = 30
	// i = 50.1 //int,原因是不能改变数据类型
	
	fmt.Println("i=",i)
}