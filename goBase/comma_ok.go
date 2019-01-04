// 下面演示断言
package main

import 
( // import的括号也可以这样写，但不提倡
	"fmt"
	"strconv"
)

type Element interface{}
type Lists []Element

type Person struct {
	name string
	age int
}

func (p Person) String() string { // 定义了String方法，实现了fmt.Stringer
    return "(name:" + p.name + " - age:" + strconv.Itoa(p.age) + " years)"
}

func main() {
	lists := make(Lists, 3)
	lists[0] = 1 // an int
	lists[1] = "Hello" // a string
	lists[2] = Person{name:"ouyanghaixiong", age:35} // a struct
	for index, element := range lists {
		if value, ok := element.(int); ok { // 进行断言: value, ok := X.(type) 断言成功ok返回true，否则返回false，value为变量X的值，type为go中的数据类型
			fmt.Printf("lists[%d] is an int and its value is %d \n", index, value)
		} else if value, ok := element.(string); ok { // if 中的断语句，这里的变量是随着if条件结束后会自动失效
            fmt.Printf("lists[%d] is a string and its value is %s \n", index, value)
		} else if value, ok := element.(Person); ok {
            fmt.Printf("lists[%d] is a Person and its value is %s \n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type \n", index)
		}
	}
	// 下面用switch实现断言
    for k, v := range lists {
		switch value := v.(type) {
		case int :
			fmt.Printf("lists[%d] is an int and its value is %d \n", k, value)
		case string :
			fmt.Printf("lists[%d] is a string and its value is %s \n", k, value)
		case Person :
			fmt.Printf("lists[%d] is a Person and its value is %s \n", k, value)
		default :
            fmt.Printf("list[%d] is of a different type \n", k)
		}
	}

}
