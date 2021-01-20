// go支持两种浮点型数：float32 和 float64

package main

import (
	"fmt"
	// "math/big"
	//"unsafe"
	// "github.com/shopspring/decimal"
)

func main() {

	//1.定义float32类型
	// var num float32 = 12.1234356789
	// fmt.Printf("值:%v--%f 类型:%T\n", num, num, num)  //float十进制用 %f 输出，保留6位小数
	// fmt.Println(unsafe.Sizeof(num))  //占用4个字节

	//2.定义float64类型
	// var num float64 = 56.2
	// fmt.Printf("值:%v--%f 类型:%T\n", num, num, num)  //float十进制用 %f 输出，保留6位小数
	// fmt.Println(unsafe.Sizeof(num))  //占用8个字节

	//3. %f 输出float类型   %.3f  输出float保留3位小数
	// var num float64 = 3.1234567890
	// fmt.Printf("值:%v--%.3f--%.7f", num, num, num)

	//4.在64位系统中go语言默认的浮点数是float64
	// num := 3.32423654
	// fmt.Println(unsafe.Sizeof(num))
	// fmt.Printf("%f--%T",num,num)

	//5.使用科学计数法表示浮点数
	// num2 := 3.13e2  //"3.13e2"表示3.13乘以10的2次方
	// fmt.Printf("%v--%f--%T\n",num2,num2,num2)

	// num3 := 3.13e-2  //"3.13e-2"表示3.13除以10的2次方
	// fmt.Printf("%v--%f--%T",num3,num3,num3)

	// 6.golang中float精度丢失问题
	// num4 := 1129.6
	// fmt.Println(num4 * 100) // 输出：112959.99999999999
	// m1 := 8.2
	// m2 := 4.8
	// fmt.Println(m1 - m2)   // 输出：3.3999999999999995
	//如何解决这个问题：可以借助第三方包：https://github.com/shopspring/decimal
	//如何安装： go get github.com/shopspring/decimal
	// c1 := 8.4
	// c2 := 4.2
	// 加法 Add
	// c3 := decimal.NewFromFloat(c1).Add(decimal.NewFromFloat(c2))
	// fmt.Println(c3)
	// //减法 Sub
	// c4 := decimal.NewFromFloat(c1).Sub(decimal.NewFromFloat(c2))
	// fmt.Println(c4)
	// //乘法 Mul
	// c5 := decimal.NewFromFloat(c1).Mul(decimal.NewFromFloat(c2))
	// fmt.Println(c5)
	// //除法 Div
	// c6 := decimal.NewFromFloat(c1).Div(decimal.NewFromFloat(c2))
	// fmt.Println(c6)

	// 7.int类型转换float类型
	// a := 10
	// b := float64(a)
	// fmt.Printf("a的类型：%T,b的类型:%T",a,b)

	// // 8.float 32转换float 64
	// var c float32 = 23.3
	// d := float64(c)
	// fmt.Printf("c的类型：%T,d的类型：%T",c,d)

	// 9.float类型转换int类型 -- 会丢失精度，不建议这样做
	// var e float32 = 21.65   
	// f := int(e)
	// fmt.Printf("%v  %T",f,f)

	/* 10.bool类型：true和false
		- bool类型的默认值是false
		- go语言中不允许整型强制转换成bool
		- bool无法参与数值运算，也无法与其他类型转换
	*/
	// var b bool
	// fmt.Printf("%v,%T",b,b)

	// a := 12
	// b := bool(a)
	// fmt.Printf("%v %T",b,b)

	var c bool
	if c {
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}

	// a := 1
	// var b bool
	// fmt.Println(a + b)

	// 11.string类型默认值为空
	// var s1 string
	// fmt.Printf("%v--%T",s1,s1)

	// 12.int类型默认值为0
	// var i1 int
	// fmt.Printf("%v--%T",i1,i1)

	//13.float类型默认值为0
	// var f1 float64
	// fmt.Printf("%v--%T",f1,f1)







}
