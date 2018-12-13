// Go Map集合。Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。 Map 是使用 hash 表来实现的。 
package main

import "fmt"

func main() {
    var countryCapitalMap map[string]string; // 声明一个集合
    countryCapitalMap = make(map[string]string); // 我靠，这样重复定义都不会报错。 且一定需要这样定义，不然报错：panic: assignment to entry in nil map
    // 即：如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
    countryCapitalMap["France"] = "Paris"; // 对map插入key-value键值对
    countryCapitalMap["Italy"] = "Rome"; // 各个国家对应的首都
    countryCapitalMap["Japan"] = "Tokyo";
    countryCapitalMap["India"] = "New Delhi";
    for key := range countryCapitalMap { // 对map循环，单个变量的话，返回的是key
        fmt.Println("Capital of ", key, " is ", countryCapitalMap[key]); // 特别注意了，这里返回的值是随机的，并不是赋值那样顺序的
    }
    
    captial, ok := countryCapitalMap["United States"]; // 判断元素是否在map集合中，是否存在这样的key，存在的话会返回其值和true
    if (ok) {
        fmt.Println("Capital of United States is ", captial);
    } else {
        fmt.Println("Capital of United States is not present");
    }
    
    deleteMap();
}

func deleteMap() { // 测试删除map中的值
    countryCapitalMap := map[string]string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo", "India":"New Delhi"}; // 如果初始化了就不需要make
    fmt.Println("原始的map：");
    for country, capital := range countryCapitalMap {
        fmt.Println("Capital of", country, "is", capital);
    }
    
    delete(countryCapitalMap, "France"); // delete 删除map中的一个值
    fmt.Println("Entry for France is deleted") 
    fmt.Println("删除元素后的map：");
    for country, capital := range countryCapitalMap {
        fmt.Println("Capital of", country, "is", capital);
    } 
}
