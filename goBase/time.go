package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	// 下面演示字符串类型的数字转unix时间戳
	i, err := strconv.ParseInt("1544768906", 10, 64) // 第一个参数一定是字符串类型。 字符串转int64
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm) // 2018-12-14 01:28:26 -0500 EST
}