package main

import "fmt"

// 演示Golang中整数类型的使用

func main(){
	var a int = 1
	fmt.Println("a=",a)
	
	//测试int8范围(-128~127)，其他int16...以此类推
	var j int8 = -128
	fmt.Println("j=",j)
	
	//测试 uint8范围(0~255)，其他uint16...以此类推
	var k uint8 = 255
	fmt.Println("k=",k)
	
	// int,uint,rune,byte使用
	var l int = 8900
	fmt.Println("l=",l)
	
	var b uint = 1
	fmt.Println("b=",b)
	
	var c byte = 255
	fmt.Println("c=",c)
	
}