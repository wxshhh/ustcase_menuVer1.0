package main

import "fmt"

func main() {
	for {
		fmt.Println("请输入命令：")
		var cmd string
		fmt.Scanln(&cmd)
		switch cmd {
		case "quit":
			fmt.Println("退出成功！")
			return
		case "help":
			fmt.Println("这里可以提供帮助！")
		}
	}
}
