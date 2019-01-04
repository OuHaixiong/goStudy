// Go 数组 。数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
package main;

import "fmt"
import "strings"

var romanNumeralDict = map[int]string{ // 函数外面声明数组时，要用完整模式。 且函数外声明的变量可以不使用
    1:"a", // 在这里，数组下标不是以0开头也不会报错
    2:"b",
}

func main() {
    isType();
    var array [5] float32; // 数组声明需要指定元素类型及元素个数
    array[0] = 3.14;
    array[2] = 0.222;
    array[4] = 4.44;
//    array[5] = 8.88; // 下标越界，直接报错： invalid array index 5 (out of bounds for 5-element array)
    for k, v := range array { // 未赋值的元素默认为0
        fmt.Println("数组array：", k, " => ", v);
    }
    
    var balance = []float32 {1000.0, 2.0, 3.4, 7, 50}; // 初始化数组
    //  初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
    //  如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小
    for k, v := range balance {
        fmt.Println(k, "=>", v);
    }

    var a = [3][4] int{ // 多维数组
        {0, 1, 2, 3},
        {4, 5},
        {8, 9, 10, 11},
    };
    println(a[2][3]);
    var i, j int;
    for i=0; i<3; i++ {
        for j=0; j<4; j++ {
            fmt.Printf("a[%d][%d] = %d \n", i, j, a[i][j]);
        }
    }
    
    var intNumber = [] int {1000, 2, 3, 17, 50}; // 可以不定义数组的长度
    // var avg float32;
    // avg = getAverage(intNumber);
    // fmt.Printf("平均值为：%f \n", avg);
    fmt.Printf("平均值为：%f \n", getAverage(intNumber));

    var testArr = []string{"a", "ouhaixiong", "欧欧海雄"}; // 声明一个切片数组
    testString := arrayToString(testArr, ",")
    fmt.Println(testString);
}

// 下面演示形参传递数组
func getAverage(arr []int) float32 { // 求平均值。 形式参数也可以不定义数组的长度
    var i, sum int;
    // var avg float32;
    var size int = len(arr); // len 求数组的长度
    
    for i=0; i<size; i++ {
        sum += arr[i];
    }
    // avg = float32(sum/size);
    // return avg; // 如果直接就返回出去了，就没有必要多加一个变量进行赋值了
    return float32(sum/size); 
}

func isType() {
    // var e interface{}
    // e[0] = "start"
    // e[1] = "end"


    // switch v := e.(type) {
    // case int : println("is integer"); break;
    // case string : println("is string"); break;
    // default : println("unknown");break;
    // }
    var x = []float32 {1000.0, 2.0, 3.4, 7, 50}; // 初始化数组
    fmt.Printf("x 的类型：%T \n", x);
}

/**
 * 数组转字符串（通过一个分隔符，连接数组的值）
 * @param []string arr 需要转字符串的数组
 * @param string separate 分隔符
 * @return string 返回通过分隔符连接在一起的字符串
 */
func arrayToString(arr []string, separate string) string {
    return strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", separate, -1); // [a b c]=>a b c=>a,b,c
}
