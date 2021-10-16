// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag"	// 标识变量
	"fmt"
	"strings"
)
// 标识变量在调用程序命令行参数时使用
var n = flag.Bool("n", false, "omit trailing newline")
// 设定标识变量n，类型为bool，标识名为n，初始值为false，提供非法标识或非法参数以及输入-h或-help时输出‘usage’内容
var sep = flag.String("s", " ", "separator")
// 设定标识变量sep，类型为string，标识名为s，初始值为" "，

func main() {
	fmt.Println(*n)	//
	flag.Parse()
	fmt.Println(*n)
	fmt.Println(flag.Args())
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(*n)
}

//!-
