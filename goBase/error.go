// Go 错误处理
/*
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
error类型是一个接口类型，这是它的定义：
type error interface {
    Error() string
}
我们可以在编码中通过实现 error 接口类型来生成错误信息。
*/
package main

// import ("fmt","math") // 这样会报错： expected ';', found ','
import("fmt";"math";);
/* 下面的为提倡写法
import (
    "fmt"
    "math"
)
*/

func main() {
    r, e := sqrt(-1);
    if e != "" {
        fmt.Println(e);
        fmt.Println(r);
    }
    //r, e := sqrt(9); // 我靠为什么在这里就报错：no new variables on left side of :=
    //fmt.Println(r);
    fmt.Println(math.Sqrt(9));
    
    // 正常情况
    if result, errorMessage := divide(100, 10); errorMessage == "" { // 注意这里判断的写法
        fmt.Println("100/10 = ", result);
    }
    
    // 当除数为0时，返回错误信息
    if _, errorMessage := divide(100, 0); errorMessage != "" { // 特别注意的是用 := 在使用函数的返回值中重复定义是不会报错的（貌似for也一样），但在其他地方是会报错的
        fmt.Println("error message is :", errorMessage);
    }
    
    // tmp := 5;
    // tmp := 6; // 这样算重复定义，会报错:no new variables on left side of :=
    
    result, errorMessage := divide(10, 3);
    if ("" != errorMessage) {
        fmt.Println("error message is :", errorMessage);
    } else {
        fmt.Println(result);
    }
    
}


func sqrt(x float64) (float64, string) { // 平方根
    if 0 > x {
        err := "math: square root of negative number";
        return 0, err;
    } else {
        z := float64(1);
        tmp := float64(0);
        for math.Abs(tmp - z) > 0.0000000001 {
            tmp = z
            z = (z + x/z)/2
        }
        return z, "";
    }
}


type divideStruct struct { // 定义一个除法结构体
    dividend int // 被除数
    divider int // 除数
}

// 注意这里的结构体是按指针来传的
func (de *divideStruct) Error() string { // 实现error接口的Error方法
    stringFormat := `
        Cannot proceed, the divider is zero.
        dividend: %d
        divider: 0
    `; // 多行字符串的声明
    return fmt.Sprintf(stringFormat, de.dividend); // 格式化后返回字符串
}

/**
 * 除法运算
 * 这里的返回值为int，如果是出现小数点的话，直接舍弃掉了，相当于转整了，如：10/3=3
 * 返回值为两个
 */
func divide(dividend int, divider int) (result int, errorMessage string) { // 除法运算
    if divider == 0 {
        dData := divideStruct{ // 声明一个结构体变量，并赋初始值
            // "dividend": dividend, 这种写法在go中是错误的，属性名一定不能加引号
            dividend : dividend,
            divider : divider,
        };
        errorMessage = dData.Error(); // 调用结构体的方法
        return; // 我靠，这样写也是可以的
    } else {
        return dividend / divider, "";
    }
}
