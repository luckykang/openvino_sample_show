// 如何命令行传参
package main

import (
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) > 1 {
	fmt.Println("hello world",os.Args[1])
	}

//go中的main函数不支持任何返回值，通过os.Exit来返回状态

}

// 终端执行 go run hello.go kang