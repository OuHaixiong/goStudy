// 练习使用go语言的关键字“go”进行并发执行
package main

import (
	"fmt"
	"time"
	"runtime"
)

func goroutine1() { // 循环内部每隔两秒钟打印一次总共执行次数
	var index = 0
	for {
		index++
		fmt.Println("goroutine 1 , 执行第 ", index, " 次")
		time.Sleep(time.Second * 2)  // 睡眠等待两秒
	}
}

func goroutine2() {
	index := 0
	for { // 循环内部每隔3秒钟打印一次总共执行次数
		index++
		fmt.Println("goroutine 2 , 执行第 ", index, " 次")
		// time.Sleep(time.Second * 2.5) // 不能这样写，因为这里的参数只能为整型（integer）
		time.Sleep(time.Second * 3)
	}
}

func main() {
	fmt.Println("main function is start")
	// runtime.GOMAXPROCS(8) //设置CPU核数
	var number = runtime.NumCPU()
	runtime.GOMAXPROCS(number)
	fmt.Println("cpu number is ", number)

	


	// 开启并发执行
	go goroutine1()
	go goroutine2()
    // 主函数中等待goroutine执行，如果主函数退出，goroutine将全部退出
	for {}
	// 程序将打印类似：
    // goroutine 1 , 执行第  29  次
    // goroutine 2 , 执行第  20  次
}
