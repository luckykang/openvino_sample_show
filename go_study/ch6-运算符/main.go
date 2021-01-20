package main

import "fmt"

func main() {
	// 1.算术运算符
	// var a  = 9
	// var b  = 4
	// fmt.Println(a + b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)  //  2  除法注意：如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分

	// var a = 10.0
	// var b = 3.0
	// fmt.Println(a / b)  //3.3333333333333335   如果是浮点数相除，会保留小数部分

	// 2.取余注意： 余数=a-(a/b)*b

	// 3.注意： ++(自增)  --(自减) 在go中是单独的语句，不是运算符

	// 4.关系运算符  == != >= <= > <

	/*
		5.逻辑运算符：
		与 AND   && 两边为True，则为True
		或 OR    || 两边有True，则为True
		非NOT	 !  如果条件为True,则为False，否则为True

	*/
	/*
		6.赋值运算符
		=   赋值
		+=  相加后再赋值
		-=
		x=
		/=
		%= 求余后再赋值
	*/

	// 练习：交换变量 使用第三个变量
	// a := 30
	// b := 2
	// t := a
	// a = b
	// b = t
	// fmt.Println(a,b)

	// a := 30  不使用第三个变量
	// b := 2
	// a = a + b
	// b = a - b
	// a = a - b
	// fmt.Println(a, b)

	//练习：距离放假还有100天，算一下还有几周零几天
	// a := 100 / 7
	// b := 100 % 7
	// fmt.Printf("距离放假有%v周零%v天", a, b)

	/*  7.位运算符
		&  两位均为1才为1
		|  两位有1个为1就是1
		^  相异或  两位不一样则为1
		<< 左移n位就是乘以2的n次方 “a<<b”表示把a的二进制位向左移动b位
		>> 右移n位就是除以2的n次方

	*/

	a := 2 // 10
	b := 3 // 11
	c := 1  //01
	fmt.Println("a&b=", a&b)   // 10   值为2
	fmt.Println("a|b=", a|b)   // 11   值为3
	fmt.Println("a^b=", a^b)   // 01   值为1
	fmt.Println("a<<b=", a<<b) //10000 值为16
	fmt.Println("a>>b=",a>>c)  //01    值为1


}
