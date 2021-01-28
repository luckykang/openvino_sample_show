package main

import (
	_ "debug/plan9obj"
	"fmt")


func main() {
	// go语言中常用的流程控制有if、for,而switch和goto主要是为了简化代码、降低重复代码而生的结构

	// 1.if判断
	// 多个判断条件
	// age := 19
	// if age > 12 {
	// 	fmt.Println("aaa")
	// }else if age > 18 {
	// 	fmt.Println("bbb")
	// }else{
	// 	fmt.Println("ccc")
	// }

	//特殊写法 此时的age变量只在if条件判断语句中生效
	// if age := 19; age>= 18{
	// fmt.Println("可以参加")
	// }else{
	// 	fmt.Println("不可以参加")
	// }

	// 2.go中只有一种循环结构-for循环
	// for i := 0;i <= 8;i++{
	// 	fmt.Println(i)
	// }
	// 变种1：
	// var i = 5
	// for ; i <= 10; i++ {
	// 	fmt.Println(i)
	// }
	// 变种2：
	// var i = 5
	// for ; i <= 10; {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for range循环
	// s := "你好，beijing"
	// for i,v := range s{
	// 	fmt.Printf("%d %c\n",i,v)
	// }

	// 流程控制之跳出for循环
	// for i := 0; i < 10; i++{
	// 	if i == 5{
	// 		break  //当i=5的时候跳出循环
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++{
	// 	if i == 5{
	// 		continue // 当i=5的时候继续下一次循环
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("over")

	// for i := 0; i < 10;{
	// 	if i == 5{
	// 		continue // 当i=5的时候继续下一次循环
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }
	// fmt.Println("over")

	//switch case 可以简化大量的判断(一个变量和具体的值做比较)
	// var n = 3
	// switch n {
	// case 1:
	// 	fmt.Println("11")
	// case 2:
	// 	fmt.Println("22")
	// case 3:
	// 	fmt.Println("33")
	// case 4:
	// 	fmt.Println("44")
	// default:
	// 	fmt.Println("无效的")
	// }

	// switch n := 5; n{
	// case 1,3,5,7,9:
	// 	fmt.Println("奇数")
	// case 2,4,6,8,10:
	// 	fmt.Println("偶数")
	// }

	// case可以加判断
	// var age = 55
	// switch {
	// case age < 18 :
	// 	fmt.Println("好好学习")
	// case age >= 18 && age <= 60 :
	// 	fmt.Println("好好工作")
	// case age > 60:
	// 	fmt.Println("享受生活")
	// default:
	// 	fmt.Println("生命美好 莫辜负")
	// }

	//fallthrough语法可以执行满足条件的case的下一个case,是为了兼容c语言中的case设计
	// 自己不要写
	// s := "a"
	// switch {
	// case s == "a":
	// 	fmt.Println("打印值为a")
	// 	fallthrough
	// case s == "b":
	// 	fmt.Println("打印值为b")
	// 	fallthrough
	// case s == "c":
	// 	fmt.Println("打印值为c")
	// }

	// 跳出多层for循环
	var flag = false
	for i := 0; i <= 10; i++ {
		for j := 'a'; j <= 'z'; j++ {
			if j == 'c'{
				flag = true
				break //跳出内层的for循环
			}
			fmt.Printf("%d %c\n",i,j)
		}
		if flag {
			break //跳出外循环
		}
	}


}
