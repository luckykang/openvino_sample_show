package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 数值类型之间的转换：

	// 1.整型之间的转换
	// var a int8 = 12
	// var b int16 = 34
	// fmt.Println(int16(a) + b)

	// 2.浮点型之间的转换
	// var a float32 = 12.12
	// var b float64 = 34.32
	// fmt.Println(float64(a) + b)

	// 3.整型和浮点型之间的转换
	// 建议：转换的时候把低位转换为高位，高位转换为低位的时候如果转换不成功就会溢出，和想要的结果不一样
	// var a int8 = 2
	// var b float32 = 12.4
	// fmt.Println(float32(a)+b)  //14.4
	// fmt.Println(a+int8(b))    //14 不准确

	// 4.其他类型转换为string
	// 方法1：使用fmt.Sprintf()

	// 使用中需要注意的转换的格式： int %d  float %f  bool %t  byte %c
	// var a int8 = 12
	// str1 := fmt.Sprintf("%d",a)
	// fmt.Printf("%v %T",str1,str1)  //12 string
	// var b bool = true
	// str2 := fmt.Sprintf("%t",b)
	// fmt.Printf("%v %T",str2,str2)

	// var c float32 = 12.1734
	// str2 := fmt.Sprintf("%.2f",c)
	// fmt.Printf("%v %T",str2,str2)  //12.17 string

	// var d byte = 'p'
	// str2 := fmt.Sprintf("%c",d)
	// fmt.Printf("%v %T",str2,str2)  //p string

	// 方法2：通过strconv把其他类型转换为string类型

	// var a int = 13
	// str1 := strconv.FormatInt(int64(a), 10)    //参数1：FormatInt必须是int64类型的数   参数2：传入int类型的进制
	// fmt.Printf("%v--%T", str1, str1)          //13--string

	/*
		参数1：传入的值
		参数2：格式化类型
		'f'(-ddd.dddd)、
		'b'(ddddp±ddd,指数为二进制)、
		'e'(-d.dddde±dd,十进制指数)、
		'E'(-d.ddddE±dd,十进制指数)、
		'g'(指数很大时用'e'格式，否则'f'格式)、
		'G'(指数很大时用'E'格式，否则'f'格式)、
		参数3：保留的小数点， -1表示不对小数点格式化
		参数4：格式化的类型  传入 64  32
	*/
	// var c float32 = 12.4698
	// str3 := strconv.FormatFloat(float64(c),'f',2,64)
	// fmt.Printf("%v %T", str3, str3)  //12.47 string

	// var b bool = true
	// str2 := strconv.FormatBool(b)
	// fmt.Printf("%v--%T", str2, str2)  //true--string

	// a := 'e'
	// str4 := strconv.FormatUint(uint64(a),10) //  10进制格式输出
	// fmt.Printf("%v %T",str4,str4)           //   101 string

	// 5.string转换为数值类型
	// (1)转换为int类型
	// a := "123"
	// str1,_ := strconv.ParseInt(a, 10, 64)   // 返回两个值，这里使用匿名变量忽略err值
	// /*
	// 参数1：string数据
	// 参数2：转换后的进制
	// 参数3：位数
	// */
	// fmt.Printf("%v %T", str1, str1)      //123 int64
	// (2)转换为float类型
	// b := "123.456676"
	// str2,_ := strconv.ParseFloat(b,10)
	// fmt.Printf("%v %T",str2,str2)        //123.456676 float64

	// 6.数值类型无法直接转换为bool类型，bool类型也无法直接转化为数值类型






}
