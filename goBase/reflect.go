// 下面演示go语言的反射
package main

import (
	"fmt"
	"reflect" // 引入反射包
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)  // 对变量x进行反射。 得到实际的值，通过v我们获取存储在里面的值，还可以去改变值。返回reflect对象
	fmt.Println("type:", v.Type()) // type: float64
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) // kind is float64: true
	fmt.Println("value:", v.Float()) // value: 3.4
	// 下面对值进行修改
	p := reflect.ValueOf(&x)
	value := p.Elem()
	// value := v.Elem() // 不能这样写，只能按照上面的两行那样写
	value.SetFloat(8.88) // 重新赋值
	fmt.Println("value:", x) // value: 8.88

	
	
	


}