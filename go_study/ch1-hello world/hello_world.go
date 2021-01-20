// 包，表明代码所在的模块
package main // package必须为main函数

//fmt用于打印数据
import (
    "fmt"

) // 引入代码依赖

// 功能实现  必须为main函数
func main()  {
    fmt.Print("ni","hao","ya")
    fmt.Println("hello","world","k")
    fmt.Println(".......")
    fmt.Print("hello golang")
    fmt.Println("hello golang")

// Print和Println区别：
// 1.一次输入多个值的时候，Println中间有空格，Print没有
// 2.Println会自动换行，print不会

    //var a = 0  // go语言中变量定义了必须使用
    //var b = "NN"
    //fmt.Println(a)
    //fmt.Println("%v",b)

    //也可以用类型推导方式定义变量
    a := 10
    b := 20
    c := 30
    fmt.Printf("%v %v\n %v",a,b,c)
    //fmt.Printf("a=%v a的类型是%T",a,a)

}





