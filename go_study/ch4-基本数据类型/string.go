package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1.定义string类型:3种
	// var str1 string = "golang"
	// var str2 = "hello"
	// str3 := "world"
	// fmt.Printf("%v-%T\n",  str1, str1)
	// fmt.Printf("%v-%T\n",  str2, str2)
	// fmt.Printf("%v-%T\n",  str3, str3)

	// 2.字符串转义符
	// str1 := "this is \n golang"   // 换行      \n
	// fmt.Println(str1)
	// str2 := "C:\\Go\\bin"       // C:\Go\bin   \\表示输出\
	// fmt.Println(str2)
	// str3 := "C:go\"bin"          // C:go"bin   \"表示输出"
	// fmt.Println(str3)

	//3.多行字符串           `是反引号，位于tab键上面的键
	// str4 := `this is str
	// 	this is str
	// 	this is str
	// `
	// fmt.Println(str4)

	// 4.len求字符串长度
	// str5 := `this is str`
	// fmt.Println(len(str5))

	// 5.拼接字符串     +  或者  Sprintf
	// str6 := `this `
	// str7 := `is `
	// str8 := `str`
	// // fmt.Println(str6+str7+str8)
	// a := fmt.Sprintf("%v %v %v",str6,str7,str8)
	// fmt.Println(a)
	//多行拼接
	// str9 := "aaaaa" +
	// "bbbbb" +
	// "ccccc"
	// fmt.Println(str9)

	// 6.分割字符串,x需要引入strings包   	strings.Split()
	// str := "123-453-678"
	// arr := strings.Split(str,"-")
	// fmt.Println(arr)   //[123 456 678]  切片，简单理解切片就是数组，但是和数组还有有些区别

	// 7.join 可以把切片转换成字符串 	strings.Join()
	// str := "123-453-678"
	// arr := strings.Split(str,"-")
	// str1 := strings.Join(arr,"-")
	// fmt.Println(str1)

	// arr:= []string{"c","c++","go"}
	// str5 := strings.Join(arr,"-")
	// fmt.Printf("%v %T",str5,str5)
	/*

		8.golang的数据类型
		基本类型：str bool int float
		复合类型：数组 切片 结构体 函数 map 通道(channel) 接口等
	*/

	// 9.判断是否包含  	strings.Contains()
	// str1 := "123"
	// str2 := "2"
	// fm := strings.Contains(str1, str2)
	// fmt.Println(fm)     //true   说明str1中包含str2

	// 10.前缀、后缀判断  strings.HasPrefix()    strings.HasSuffix()
	// str1 := "123 45"
	// str2 := "1"
	// str3 := "4"
	// fm1 := strings.HasPrefix(str1,str2)
	// fm2 := strings.HasSuffix(str1,str3)
	// fmt.Println(fm1,fm2) 

	// 11.判断子字符串出现的位置
	str1 := "123 4533323"
	str3 := "3"
	fm1 := strings.Index(str1,str3)   //第一次出现的下标
	fm2 := strings.LastIndex(str1,str3)   //最后一次出现的下标
	fm3 := strings.Count(str1,str3)
	str2 := "8"
	fm4 := strings.Index(str1,str2)
	fmt.Println(fm1,fm2,fm3,fm4)  //查找不到返回-1
	




}
