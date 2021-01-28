package main

import (
	"fmt"
	_ "unsafe" //如果暂时不用包，可以使用 _
)

func main() {
	// byte和rune类型
	// 1.golang中定义字符   字符属于int类型  使用''
	// 注意和字符串不同，字符串用""
	// a := 'c'
	// fmt.Printf("%v %T", a,a)  // 输出c对应ASCII码的值  99 int32
	// 2.原样输出字符
	// fmt.Printf("%c %T", a,a)  // 输出c对应ASCII码的值  c int32

	//3.golang中字符有两种：
	// - uint类型，也叫byte类型，代表了ASCII码的一个字符
	// - rune类型，代表了一个UTF-8字符
	// str := "access"
	// fmt.Printf("值：%v 原样输出：%c 类型：%T",str[4],str[4],str[4])   // 115 s uint8

	// 4.一个汉字占用3个字节(utf-8)   一个字母占用一个字节(ascii)
	// 注意：unsafe.Sizeof()没法查看string类型数据所占用的存储空间，但是可以使用len来获取长度
	// str := "你好哈哈"
	// fmt.Println(unsafe.Sizeof(str))  //显示16，无法获取字符长度
	// fmt.Println(len(str))            //12  占用12个字节

	// 5.定义一个字符，字符的值为汉字。汉字使用utf-8编码，编码后的值为int类型
	// str1 := '明'
	// fmt.Printf("%v %c %T",str1,str1,str1)  //26126 明 int32

	// 6.通过循环输出字符
	// byte类型
	// a := "golang"
	// for i := 0;i < len(a);i++ {
	// 	fmt.Printf("%v(%c)",a[i],a[i])
	// }   // 103(g)111(o)108(l)97(a)110(n)103(g)

	// rune类型
	// b := "中国你好"
	// for _, v := range b {
	// 	fmt.Printf("%v(%c)",v,v)
	// }   //20013(中)22269(国)20320(你)22909(好)

	// 7.修改字符串：要先把字符串转换为byte或rune类型，完成后在转回string类型。无论是哪种转换，都会重新分配内存，并复制字节数组
	s1 := "aaple"
	byte_s1 := []byte(s1)
	byte_s1[1] = 'p'
	fmt.Println(string(byte_s1)) //apple

	s2 := "中国我好"
	rune_s2 := []rune(s2)
	rune_s2[2] = '你'
	fmt.Println(string(rune_s2)) //中国你好

}
