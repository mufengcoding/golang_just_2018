// go 标准库 os
package main

import (
	"fmt"
	"os"
)

func main() {
	//预定义变量用来保存命令行参数
	fmt.Print(os.Args)
	fmt.Print(os.Hostname())
	fmt.Print(os.Getegid())

	env := os.Environ()
	for k, v := range env {
		fmt.Println(k, v)
	}
}
