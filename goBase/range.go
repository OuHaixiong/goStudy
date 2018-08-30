// Go 语言范围(Range)。Go 语言中 range 关键字用于for循环中迭代数组(array)、切片(slice)、链表(channel)或集合(map)的元素。在数组和切片中它返回元素的索引值，在集合中返回  key-value 对的 key 值。
package main

import "fmt"

func main() {
    numbers := []int{2, 3, 4};
    sum := 0;
    for _, num := range numbers { // 通过range循环数组。空白符"_"是对数组的索引进行省略（丢掉，不使用）
        sum += num;
    }
    fmt.Println("sum:", sum);
    
    kvs := map[string]string{"a":"apple", "b":"banana"}; // map 数组类型，go中的map就和js中的json对象是一样的
    for k, v := range kvs {
        fmt.Printf("%s => %s \n", k, v);
    }
    
    for i, c := range "ago" { // range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
        fmt.Println(i, c); // i为索引，c为ASCII值
    }
    
}
