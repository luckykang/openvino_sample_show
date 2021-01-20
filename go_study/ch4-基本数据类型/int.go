/*
数据类型介绍：
go中的数据类型分为：基本数据类型和复合数据类型
基本数据类型有：整型、浮点型、布尔型、字符串
复合数据类型有：数组、切片、结构体、函数、map、通道(channel)、接口等。

整型分为两类：
有符号的整型按长度分为：int8(1个字节，-128~127)、int16(2个字节)、int32(4个字节)、int64(8个字节)
对应的无符号整型：uint8(1个字节，0~255)、uint16(2个字节)、uint32(4个字节)、uint64(8个字节)

*/
package main

import (
	"fmt"
	// "unsafe"
)

func main() {

	// //1.定义一个整型
	// var num  = 10
	// num = 22
	// fmt.Printf("值=%v  类型:%T",num,num)

	//int8范围(-128~127)
	// var num int8 = 127
	// fmt.Printf("值=%v 类型:%T",num,num)

	//unsafe.Sizeof 查看不同程度的整型 在内存里面的存储空间
	// var num int8 = 43
	// fmt.Printf("值=%v 类型：%T\n", num, num)
	// fmt.Println(unsafe.Sizeof(num)) //输出结果表示占用1个字节

	// var num uint16 = 256
	// fmt.Printf("值=%v 类型：%T\n", num, num)
	// fmt.Println(unsafe.Sizeof(num)) //输出结果表示占用2个字节

	//int的类型转换
	// var num1  uint16 = 34
	// var num2  uint32 = 2
	// fmt.Println(num1 + uint16(num2))   //num2转换成uint16

	//高位向低位转换要注意
	// var num1 int32 = 130
	// fmt.Println(int8(num1))  //int8为-127~128 结果为-126不准确

	// num := 32
	// fmt.Printf("%v %T\n",num,num)
	// fmt.Println(unsafe.Sizeof(num)) //结果为8位字节，说明int类型默认在64位系统是int64，在32位系统上是int32

	//数字字面量语法： 
	num := 32
	fmt.Printf("%v %T\n", num, num)  //%v表示原样输出,%T表示类型
	fmt.Printf("%d %T\n", num, num)  //%d表示十进制输出
	fmt.Printf("%b %T\n", num, num)  //%b表示二进制输出
	fmt.Printf("%o %T\n",num,num)    //%o表示八进制输出
	fmt.Printf("%x %T\n",num,num)    //%x比哦是十六进制输出

}
