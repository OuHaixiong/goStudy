// Go 条件语句
package main;

import "fmt";

func main() {
    var grade string = "B"; // 定义局部变量
    var marks int = 90;
    
    switch marks {
        case 90 : grade = "A"; // 不需要break
        case 80 : grade = "B";
        case 50, 60, 70 : grade = "C"; // 多个可能符合条件的值，使用逗号分割它们
        default : grade = "D";
    }
    
    switch { // switch 的另一种表现形式
        case grade == "A" : fmt.Println("优秀！");
        case grade == "B", grade == "C" : fmt.Println("良好");
        case grade == "D" : fmt.Println("不及格");
        default : fmt.Println("差");
        
    }
    
    fmt.Println("你的等级是：", grade);
    
    typeSwitch();
}

func typeSwitch() { // switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
    var x interface {}
    switch i := x.(type) {
        case nil : fmt.Printf("x 的类型：%T", i);
        case int : fmt.Printf("x 是 int 型");
        case float64 : fmt.Printf("x 是 float64 型");
        case func(int) float64 : fmt.Printf("x 是 func(int) 型"); // 这判断貌似有问题
        case bool, string : fmt.Printf("x 是 bool 或 string 型");
        default : fmt.Printf("未知类型");
    } 
    fmt.Print("\n");
}

// TODO  有个select语句，也是用来最选择的，有机会再学习
