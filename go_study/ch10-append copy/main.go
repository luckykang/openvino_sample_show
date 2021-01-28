package main

import (
	"fmt"
	_ "sort"
)

func main() {
	//1.append  为切片增加元素
	// s1 := []string{"beijing", "shanghai", "tianjin"}
	// // 错误的写法，会导致索引越界   报错：runtime error: index out of range [3] with length 3
	// //s1[3] = "shenzhen"
	// // 调用append函数必须用原来的切片变量接收返回值
	// //append增加元素，原来的底层数组放不下的时候，go底层就会把底层数组换一个
	// // append扩容策略的原理见图
	// s1= append(s1,"shenzhen")
	// // fmt.Println(s1)
	// s2 := []string{"wuhan","xian","suzhou"}
	// s1 = append(s1, s2...)   //   ...表示拆开元素
	// fmt.Printf("s1=%v\n len(s1):%d\n cap(s1):%d\n",s1,len(s1),cap(s1))

	//2.copy
	// a1 := []int{1,3,4}
	// a2 := a1  		//赋值
	// var a3 = make([]int,3,3)    //长度为3，容量为3
	// copy(a3,a1)   //把a1的元素copy到a3
	// fmt.Println(a1,a2,a3)   //[1 3 4] [1 3 4] [1 3 4]
	// a1[1] = 100
	// fmt.Println(a1,a2,a3)   //[1 100 4] [1 100 4] [1 3 4]

	//3.go中没有删除切片元素的专用方法，我们可以使用切片本身的特性删除元素
	// a := []int{1,2,3,4,5,6,7,8,9}
	// //删除索引为2的元素
	// a = append(a[:2],a[3:]...)
	// fmt.Println(a)          //[1 2 4 5 6 7 8 9]

	// a := [...]int{1,3,4,5,6,7,8}  //数组
	// b := a[:]                 //切片
	// fmt.Printf("b=%v len(b)%d cap(b)%d\n",b,len(b),cap(b))   //b=[1 3 4 5 6 7 8] len(b)7 cap(b)7
	//操作：删除索引为3的元素
	//1.切片不保存具体的值
	//2.切片对应一个底层数组
	//3.底层数组都是占用一块连续的内存
	// fmt.Printf("%p\n",&b[3])    //改之前b的地址
	// b = append(b[:3],b[4:]...)   //修改了底层数组！！！把后面的元素整体前挪一位，最后一位复写一次
	// fmt.Printf("b=%v len(b)%d cap(b)%d\n",b,len(b),cap(b))  //b=[1 3 4 6 7 8] len(b)6 cap(b)7
	// fmt.Printf("%p\n",&b[3])    //改之后b的地址
	// fmt.Println(a)  //[1 3 4 6 7 8 8] 底层数组

	// 切片的练习
	// var a = make([]int,5,10)  //创建切片，长为5，容量为10
	// fmt.Printf("a=%v len(a)%d cap(a)%d\n",a,len(a),cap(a)) //a=[0 0 0 0 0] len(a)5 cap(a)10
	// for i := 0;i<10;i++{
	// 	a = append(a, i)
	// }
	// fmt.Printf("a=%v len(a)%d cap(a)%d",a,len(a),cap(a))  //a=[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9] len(a)15 cap(a)20

	//sort包
	// var a1 = [...]int{1,2,3,4,5,6,7,8}
	// sort.Ints(a1[:])     //把数组转化为切片再排序
	// fmt.Println(a1)

	// 指针   go语言中不存在指针操作，只需要记住两个符号：1. & ：取地址   2. * ：根据地址取值
	a := 2
	fmt.Println(&a)  //取地址  0xc000014090   
	

}
