package main

import "fmt"

	/*常量：程序运行中不变的量    const  
	定义了常量之后不能修改
	*/ 


const pi = 1.333

//批量声明常量
const (
	n1 = 12
	n2 = 5
)

// 批量声明常量时，如果某一行没有赋值，默认和上一行一致
const (
	c1 = iota
	c2
	c3
)

// iota  是go语言中的常量计数器，只能在常量表达式中使用
// iota 在const关键字出现的时候将被重置为0	
// const中每新增一行常量声明将使iota计数一次

// const (
// 	a = iota  // 0
// 	b         // 1
// 	c         // 2
// )


const(
	a = iota   // 0
	b = 12
	c = iota    // 2   const中每新增一行常量声明将使iota计数一次
	d           // 3
) 

//多个常量声明在一行
const(
	a1,a2 = iota+1,iota+2     //iota  0

	a3,a4 = iota+1,iota+2    // iota  1
)

func main() {
	// fmt.Println("n1",n1)
	// fmt.Println("n2",n2)
	// fmt.Printf("%v %v %v",c1,c2,c3)
	// fmt.Printf("%v %v %v",a,b,c)
	// fmt.Printf("%v %v %v %v",a,b,c,d)
	fmt.Printf("%v %v %v %v",a1,a2,a3,a4)








}
