// Go 类型转换
package main;

import "fmt";

func main() {
    var sum int = 17;
    var count int = 5;
    var mean float32;
    
    mean = float32(sum)/float32(count);
    fmt.Printf("mean的值为： %f \n", mean);
    
    var x float32 = 3.4;
    var y float32 = 3.5;
    var z float32 = 3.9;
    var u float32 = 3.0;
    var v int;
    v = int(x);
    fmt.Println(v);
    v = int(y);
    fmt.Println(v);
    v = int(z);
    fmt.Println(v);
    v = int(u); // 全部为3，浮点型通过int进行转整后，小数点后面的数直接舍弃
    fmt.Println(v);
    var str string = "我靠";
    fmt.Println(str);
    // v = int(str); // 直接报错了：cannot convert str (type string) to type int 提示字符串型不能转整型
    str = "123";
    // v = int(str); // 同样的道理，这里的字符串也不能转整
    fmt.Println("转字符串前v为：", v);
    str = string(v); // 转后为：  
    fmt.Println("转后字符串为：", str);
    
}

// 总结：go为强类型的语言，貌似在类型转换上面支持并不好