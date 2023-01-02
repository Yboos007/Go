package main

import "fmt"

// 演示Golang中+的使用

func main(){
	var i = 1
	var j = 2
	var r = i + j
	fmt.Println("r=",r)
	
	var str1 = "hello"
	var str2 = "world"
	var str3 = str1 + str2
	fmt.Println("str3=",str3)
	
}