package main			

import "fmt"


func getUserinfo() (string,int){
	return "liming",10
}




// var a = "st"    //全局变量
// b := "dd"    //错误写法，局部变量不能用于全局

func main(){


	// fmt.Println(a)
	// fmt.Println(b)


/*
	go语言中变量名由字母、数字、下划线组成，首个字符不能为数字。go语言中的关键字和保留字都不能用作变量名。
	go语言中的变量需要申明后才能使用，同一作用域内不支持重复申明。并且go语言的变量申明后必须使用。
	
	1.var 申明变量
	var 变量名称 类型
*/
	// var username string
	// fmt.Println(username)  //变量声明后没有初始化值为空

	// var a1="libai"
	// fmt.Println(a1)  

	// go语言变量定义以及初始化
	//第一种方法
	// var username string
	// username = "李白"
	// fmt.Println(username)
	//第二种方法
	// var username string = "杜甫"
	// fmt.Println(username)
	//第三种
	// var username = "杜牧"
	// fmt.Println(username)
	
	// var username = "康尔"
	// var age = 22
	// var sex = "男"
	// username = "kang"  //重新给username赋值
	// fmt.Println(username,age,sex)

/*  
	一次申明多个变量
	var 变量名称，变量名称 类型
	var (		
		变量名称 类型
		变量名称 类型

	)
*/

	// var a1,a2 string
	// 	a1 = "sdf"
	// 	a2 = "ddd"
	// 	fmt.Println(a1,a2)

	// var (
	// 	username string
	// 	age int

	// )
	// username = "life"
	// age = 12
	// fmt.Println(username,age)

/*	
	类型推导
	var (
		username = "liming"
		age = 15
		sex = "男"
	)
	fmt.Println(username,age,sex)
*/

/*
短变量申明法：在函数内部，可以使用更简洁的 := 方式申明并初始化变量

注意：短变量只能用于申明局部变量，不能用于全局变量的申明
*/	
	// username := "kangge"    //申明变量和赋值变量是一体的
	// fmt.Println(username)
	// fmt.Printf("类型:%T",username)

/*
	使用短变量申明法一次申明多个变量
	a,b,c := 12, 14, "fdsf"
	fmt.Printf("a的类型:%T b的类型:%T c的类型:%T",a,b,c)
*/

/*
	匿名变量：在使用多重赋值时 如果想要忽略某个值，可以使用匿名变量 
	匿名变量用 _ 表示
	func getUserinfo() (string,int){
	return "liming",10
}

*/
// var username,age = getUserinfo()
// fmt.Println(username,age)

	var username, _ = getUserinfo()  //忽略age的值
	fmt.Println(username)
	//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复申明,上边用了 _ 还可以继续使用
	var _, age = getUserinfo()
	fmt.Println(age)

//	var username = "liuba" //错误演示，说明变量不能重复申明




}