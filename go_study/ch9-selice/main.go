package main

import "fmt"

/*
切片：是一个拥有相同数据类型元素的可变长度的序列，基于数组类型做的一层封装，支持自动扩容
是一个引用类型  内部结构包含地址、长度、容量



*/

func main() {
	//1.定义一个存放int类型元素的切片
	// var s1 []int
	// var s2 []string //string类型的
	// fmt.Println(s1, s2)   //[] []
	// fmt.Println(s1 == nil)  // nil代表空

	//2.初始化
	// s1 = []int{1, 3, 4}
	// s2 = []string{"aa", "bb", "cc"}
	// fmt.Println(s1, s2)   //[1 3 4] [aa bb cc]

	//3.切片的长度len() 和容量cap()
	// fmt.Printf("len(s1):%v cap(s1):%v ",len(s1),cap(s1))

	// //4.由数组得到切片
	// a := [...]int{1, 3, 4, 5, 6, 7}
	// b := a[0:3]
	// fmt.Println(b) //[1 3 4]   左闭右开，左包含右不包含
	// c := a[:4]
	// fmt.Println(c) //[1 3 4 5]
	// d := a[2:]
	// fmt.Println(d) //[4 5 6 7]
	// e := a[:]
	// fmt.Println(e) //[1 3 4 5 6 7]
	// // 切片指向了底层的数组  切片长度是它元素的个数    切片的容量是底层数组从切片的第一个元素到最后的元素数量
	// fmt.Printf("len(a):%v cap(a):%v\n", len(a), cap(a)) //len(a):6 cap(a):6
	// fmt.Printf("len(b):%v cap(b):%v\n", len(b), cap(b)) //len(b):3 cap(b):6
	// fmt.Printf("len(c):%v cap(c):%v\n", len(c), cap(c)) //len(c):4 cap(c):6
	// //切片只能向右扩容，所以会丢失左边的容量部分
	// fmt.Printf("len(d):%v cap(d):%v\n", len(d), cap(d)) //len(d):4 cap(d):4

	// //5.切片再切片
	// f := d[3:]
	// fmt.Printf("len(f):%v cap(f):%v\n", len(f), cap(f))   //len(f):1 cap(f):1
	// // 切片是引用类型，都指向了底层的一个数组
	// a[5] = 7000
	// fmt.Println(d)   //[4 5 6 7000]

	// b[0] =1000
	// fmt.Println(a)

	//make()函数创造切片
	//  s1 := make([]int,5,10)
	//  fmt.Printf("s1:%v len(s1)%v cap(s1)%v\n",s1,len(s1),cap(s1))  //s1:[0 0 0 0 0] len(s1)5 cap(s1)10

	//  s2 := make([]int,0,10)
	//  fmt.Printf("s2:%v len(s2)%v cap(s2)%v",s2,len(s2),cap(s2))   //s2:[] len(s2)0 cap(s2)10

	//切片的本质
	//切片就是一个框，框住了一块连续的内存
	//切片属于引用类型，真正的数据都是保存在底层数组里

	//切片不能直接比较，切片唯一合法的比较操作是和nil比较。一个nil值的切片没有底层数组，它的切片的长度和容量都是0.但我们不能说长度和容量都是0的切片一定是nil
	// s3 := []int{}
	// fmt.Println(s3)  //[]

	//切片的赋值    s3和s4指向了同一个底层数组
	// s3 := []int{1,3,4}
	// s4 := s3
	// s3[0] = 1000
	// fmt.Println(s3,s4)  //[1000 3 4] [1000 3 4]

	//切片的遍历
	//1.索引遍历
	s5 := []int{1,3,4}
	for i := 0;i< len(s5);i++{
		fmt.Println(s5[i])
	}
	//2.for range循环
	for _,v := range s5{
		fmt.Println(v)
	}




	
}
