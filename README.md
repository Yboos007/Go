## 学习Go语法和编码规范

1、Go入门
https://tour.go-zh.org/list

包括基础、接口和方法以及并发三部分。

2、如何使用Go编程
https://go-zh.org/doc/code.html

3、实效Go编程
https://go-zh.org/doc/effective_go.html

4、Go项目的FAQ
https://go-zh.org/doc/faq

5、Go维护的维基百科
https://github.com/golang/go/wiki

6、更多学习
https://github.com/golang/go/wiki/Learn

四、Go资源-标准库等
1、Go标准库文档
https://go-zh.org/pkg/

2、Go命令文档
https://go-zh.org/doc/cmd

3、官方Go语言规范
https://go-zh.org/ref/spec

4、Go内存模型
https://go-zh.org/ref/mem

5、Go发行历史
https://go-zh.org/doc/devel/release.html

五、Go官方博客
主要是Go项目的官方博客：https://blog.golang.org/

1、函数、文本、共享内存、web应用开发等
函数 - Go 语言中的一等公民
生成任意文本：马尔可夫链算法
通过通信共享内存
编写Web应用 - 构建简单的Web应用
（2）、Go语言的一些用法
JSON-RPC：有关接口的故事
Go 的声明语法
Defer、Panic 和 Recover
Go 并发模式：超时，继续
Go 切片：用法和本质
GIF 解码器：Go 接口练习
Go 与错误处理
组织 Go 代码
3、Go包
JSON 和 Go - 使用json包。
数据一坨 - gob包的设计与使用
反射法则 - reflect包基础。
Go 图像包 - image包基础。
Go 图像绘制包 - image/draw包基础。
4、Go工具
关于 Go 命令 - 为什么写它？它是什么？它不是什么？它怎么用？
C? Go? Cgo! - 使用cgo连接 C 代码。
使用GDB调试Go代码
Godoc：编写 Go 代码文档 - 为godoc编写好的文档。
Go 程序性能分析
数据竞态检测器 - 测试竞态条件下的 Go 程序。
Go 竞态检测器介绍 - 对竞态检测器的介绍。
Go 汇编器快速指南 - 对 Go 使用的汇编的介绍。



## 了解golang



![image-20230101122832033](E:\MD文件图片集\image-20230101122832033.png)

对上图说明



1、go文件后缀为 .go

2、package main

```
表示该main.go文件所在的包是main,在go中，每个文件都必须归属于一个包
```

3、import "fmt"

```
表示，引入一个包，包名为fmt,引入该包后，就可以使用fmt的函数，比如fmt.Println
```

4、func main(){

​	}

```
func是一个关键字，表示一个函数
main是函数名，是一个主函数，即程序的入口
```

5、fmt.Println

```
调用fmt包的函数，Println输出‘Hello,World’
```

通过go build命令对该文件进行编译，生成.exe文件

![image-20230101123557695](E:\MD文件图片集\image-20230101123557695.png)

运行程序文件

![image-20230101123627177](E:\MD文件图片集\image-20230101123627177.png)

也可以使用 go run 执行.go文件，相当于执行脚本文件（执行速度慢）

![image-20230101124008658](E:\MD文件图片集\image-20230101124008658.png)





## 转义字符(escape char)

##### 常用转义字符

\t ：一个制表位，实现对齐

\n:  换行符

\\\:   一个\

\\":  一个"

\r: 一个回车



## 使用细节

1、对于行注释和块注释，被注释的文字，不会被go识别

2、块注释里面不允许有块注释嵌套【需要注意】



## 代码的规范风格

1、正确的注释和注释风格

​	1.1 Go官方推荐使用行注释来注释整个方法和语句

​	1.2 带看Go源码

2、正确的缩进和空白

​	2.1 使用一次**tab操作**，实现缩进默认整体向右边移动，使用**shift+tab**实现整体向左移动

​	2.2 使用gofmt 来进行格式化【加-w 格式化后写入源文件】

3、运算符两边习惯性各加一个空格。例如：2 = 4 * 5 。

```
	var num = 2 + 4 * 5
```

4、Go语言的代码风格

```
package main

import "fmt"

func main() {

	fmt.Println("Jerry")
}
```

上面写法正确的

```
package main

import "fmt"

func main() 
{

	fmt.Println("Jerry")
}
```

Go语言不允许这样写代码，写法是错误的

Go设计者思想：**一个问题尽量只有一个解决办法**

5、一行最长不超过80个字符，超过的请使用换行展示，尽量保持格式优雅

```
	fmt.Println("Jerry \n HelloWorld")
```



## 变量介绍

### 一、基本概念

变量相当于内存中一个数据存储空间的表示，你可以把变量看作是一个房间的门牌号，通过门牌号我们可以找到房间，同样的道理，通过变量名可以访问到变量（**值**）

### 二、变量的使用步骤



##### 1、声明变量（也叫定义变量）

##### 2、非变量赋值

##### 3、使用变量



### 三、变量使用的注意事项

##### 1、变量表示内存中的一个存储区域

##### 2、该区域有自己的名称（变量名）和类型

##### 3、golang变量使用的三种方式

​	3.1 第一种：指定变量类型，声明后若不赋值，使用默认值

​	3.2 第二种：根据值自行判定变量类型（类型推导）

​	3.3 第三种：省略var,注意：左侧的变量不应该是已经声明过的，否则会导致编译错误

​	3.4 第四种：多变量声明

​			在编程中，有时我们需要一次性声明多个变量，go也提供了语法

```
package main

import "fmt"

func main(){
	// 该案例演示golang如何一次性声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
	
	//一次性声明多个变量的方式2
	// var n1, name, n3 = 100, "tom", 888
	// fmt.Println("n1=",n1,"name=", name, "n3=",n3)
	
	//一次性声明多个变量方式3.同样可以使用类型推导
	n1, name, n3 := 100, "tom~", 888
	fmt.Println("n1=",n1,"name=", name, "n3=",n3)
}
```

​		如何一次性声明多个全局变量【在go中函数外部定义变量就是全局变量】

```
//定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"
var(
	n3 = 300
	n4 = 400
	name2 = "jerry"
)

	//输出全局变量
	// fmt.Println("n1=",n1,"name=", name, "n2=",n2)
	fmt.Println("n3=",n3,"name2=", name2, "n4=",n4)
```

##### 4、该区域的数据值可以在同一类型范围内不断变化

##### 5、变量在同一个作用域（在一个函数或者在代码块）内不能重名

##### 6、变量=变量名+值+数据类型，这一点要注意

##### 7、Golang的变量如果没有赋初值，编译器会使用默认值，比如int默认值为0,string默认值为空



