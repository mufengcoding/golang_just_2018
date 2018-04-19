package main

import (
	"fmt"
	"os"
)

func main() {
	target := "用户名"
	x := "密码"
	fmt.Println("请输入", target, x)
	if len(os.Args) > 1 {
		target = os.Args[1]
		x = os.Args[2]
		if target == "1" && x == "2" {
			fmt.Printf("登录成功\n")
		}
	}

}
